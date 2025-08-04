package views

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/ma-guo/admin-core/config"
	"github.com/ma-guo/admin-core/utils/bearer"
	"github.com/ma-guo/admin-core/xorm/models"
	"github.com/ma-guo/admin-core/xorm/services"

	"github.com/ma-guo/admin-core/app/common/consts"

	"github.com/ma-guo/niuhe"
	"github.com/ma-guo/zpform"
	cache "github.com/patrickmn/go-cache"
)

// 做简单的攻击检测

type isCustomRoot interface {
	ThisIsACustomRoot()
}

type V1ApiProtocol struct {
	store     *cache.Cache
	skipUrl   map[string]bool
	proxy     niuhe.IApiProtocol
	routes    []*niuhe.RouteItem // 路由列表
	routeInit bool               // 路由是否已经初始化
}

// 检查权限
func (proto V1ApiProtocol) checkPers(c *niuhe.Context, jwt *bearer.Bearer) error {
	path := c.Request.URL.Path
	method := c.Request.Method
	key := "checkPers"
	// niuhe.LogInfo("check %v", path)
	if _, has := proto.getCache(key, jwt.Uid, method, path); has {
		// niuhe.LogInfo("has menu cache %v", path)
		return nil
	}
	svc := services.NewSvc()
	defer svc.Close()
	menus, err := svc.Api().GetMenus(method, path)
	if err != nil {
		niuhe.LogInfo("%v", err)
		return err
	}
	if len(menus) == 0 {
		proto.cacheMinute(menus, key, jwt.Uid, method, path)
		// niuhe.LogInfo("no menu cache %v", path)
		return nil
	}

	roleMenus, err := svc.RoleMenu().GetMenusByUid(jwt.Uid)
	if err != nil {
		niuhe.LogInfo("%v", err)
		return err
	}
	for _, menu := range menus {
		if menu.Type != consts.MenuTypeGroup.BUTTON.Value {
			continue
		}

		// 存在一个权限就可以了
		if _, has := roleMenus[menu.Id]; has {
			// niuhe.LogInfo("has menu pers %v, %v", menu.Id, path)
			proto.cacheMinute(menus, key, jwt.Uid, method, path)
			return nil
		}
	}
	return fmt.Errorf("无API访问权限")
}

func (proto V1ApiProtocol) checkAuth(c *niuhe.Context) error {
	path := c.Request.URL.Path
	// 跳过 url 不检查权限
	isSkipUrl := false
	if _, has := proto.skipUrl[path]; has {
		isSkipUrl = true
	}
	skipOrError := func(code int, message string) error {
		if isSkipUrl {
			jtw := bearer.NewBearer(config.Config.Secretkey, 1, "tmpname")
			c.Set(consts.Authorization, jtw)
			return nil
		}
		return niuhe.NewCommError(code, message)
	}
	token := c.GetHeader(consts.Authorization)
	if len(token) < 10 {
		return skipOrError(consts.AuthError, "token error")
	}
	token = token[len(consts.Bearer)+1:]
	if old, has := proto.getCache(consts.Authorization, token); has {
		jwt := old.(*bearer.Bearer)
		c.Set(consts.Authorization, jwt)
		err := proto.checkPers(c, jwt)
		if err != nil {
			niuhe.LogInfo("%v %v", jwt.Username, err)
			return skipOrError(consts.PersError, "无API访问权限")
		}
		return nil
	}
	jwt := bearer.NewBearer(config.Config.Secretkey, 0, "")
	err := jwt.Parse(token)
	if err != nil {
		niuhe.LogInfo("%v", err)
		return skipOrError(consts.AuthError, err.Error())
	}
	err = proto.checkPers(c, jwt)
	if err != nil {
		niuhe.LogInfo("%v %v", jwt.Username, err)
		return skipOrError(consts.PersError, "无API访问权限")
	}
	c.Set(consts.Authorization, jwt)
	proto.setCache(jwt, 5*time.Minute, consts.Authorization, token)
	return nil
}

// 添加路由, prefix 一般情况下设置为空即可
func (proto *V1ApiProtocol) AddRoute(prefix string, routes []*niuhe.RouteItem) {
	if proto.routes == nil {
		proto.routes = make([]*niuhe.RouteItem, 0)
	}
	// 添加前缀
	for _, route := range routes {
		if prefix != "" && strings.Index(route.Path, prefix) != 0 {
			route.Path = prefix + route.Path
		}
	}
	proto.routes = append(proto.routes, routes...)
}

// / 初始化路由
func (proto *V1ApiProtocol) InitRoute() error {
	if proto.routeInit {
		return nil
	}
	proto.routeInit = true
	svc := services.NewSvc()
	defer svc.Close()
	apis, err := svc.Api().GetAll()
	if err != nil {
		niuhe.LogInfo("%v", err)
		return err
	}
	for _, route := range proto.routes {

		hasApi := false
		for _, api := range apis {
			if route.Method == api.Method && route.Path == api.Path {
				// 名字相同的则不处理
				hasApi = true
				if api.Name != route.Name {
					api.Name = route.Name
					// 更新 name 字段
					err = svc.Api().Update(api.Id, api)
					if err != nil {
						niuhe.LogInfo("update error %v", err)
						return err
					}
				}
				break
			}
		}
		if !hasApi {
			err = svc.Api().Insert(&models.SysApi{
				Method:  route.Method,
				Path:    route.Path,
				Name:    route.Name,
				Remark:  "",
				MenuIds: []int64{},
			})
			if err != nil {
				niuhe.LogInfo("insert error: %v", err)
				return err
			}
		}
	}
	return nil
}

func (proto V1ApiProtocol) Read(c *niuhe.Context, reqValue reflect.Value) error {
	if proto.proxy != nil {
		return proto.proxy.Read(c, reqValue)
	}
	err := proto.checkAuth(c)
	if err != nil {
		return err
	}
	if err = zpform.ReadReflectedStructForm(c.Request, reqValue); err != nil {
		return niuhe.NewCommError(-1, err.Error())
	}
	return nil
}

func (proto V1ApiProtocol) Write(c *niuhe.Context, rsp reflect.Value, err error) error {
	if proto.proxy != nil {
		return proto.proxy.Write(c, rsp, err)
	}
	if c.IsIgnoreResult() {
		return nil
	}
	rspInst := rsp.Interface()
	if _, ok := rspInst.(isCustomRoot); ok {
		c.JSON(http.StatusOK, rspInst)
	} else {
		var response map[string]interface{}
		if err != nil {
			if commErr, ok := err.(niuhe.ICommError); ok {
				c.CheckCode(commErr.GetCode())
				response = map[string]interface{}{
					"result":  commErr.GetCode(),
					"message": commErr.GetMessage(),
				}
				if commErr.GetCode() == 0 {
					response["data"] = rsp.Interface()
				}
			} else {
				c.CheckCode(-1)
				response = map[string]interface{}{
					"result":  -1,
					"message": err.Error(),
				}
			}
		} else {
			response = map[string]interface{}{
				"result": 0,
				"data":   rspInst,
			}
		}
		c.JSON(http.StatusOK, response)
	}
	return nil
}

func (proto *V1ApiProtocol) getCache(prefix string, args ...interface{}) (interface{}, bool) {
	key := prefix
	for _, arg := range args {
		key += fmt.Sprintf(":%v", arg)
	}
	return proto.store.Get(key)
}

func (proto *V1ApiProtocol) setCache(val interface{}, duration time.Duration, prefix string, args ...interface{}) {
	key := prefix
	for _, arg := range args {
		key += fmt.Sprintf(":%v", arg)
	}
	proto.store.Set(key, val, duration)
}

func (proto *V1ApiProtocol) cacheMinute(val interface{}, prefix string, args ...interface{}) {
	proto.setCache(val, 1*time.Minute, prefix, args...)
}
