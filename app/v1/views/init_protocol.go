package views

import (
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/ma-guo/admin-core/config"
	"github.com/ma-guo/admin-core/utils/bearer"

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
	store   *cache.Cache
	skipUrl map[string]bool
	proxy   niuhe.IApiProtocol
}

func (proto V1ApiProtocol) checkAuth(c *niuhe.Context) error {
	path := c.Request.URL.Path
	if _, has := proto.skipUrl[path]; has {
		jtw := bearer.NewBearer(config.Config.Secretkey, 1, "tmpname")
		c.Set(consts.Authorization, jtw)
		return nil
	}
	token := c.GetHeader(consts.Authorization)
	if len(token) < 10 {
		return niuhe.NewCommError(consts.AuthError, "token error")
	}
	token = token[len(consts.Bearer)+1:]
	if old, has := proto.getCache(consts.Authorization, token); has {
		jwt := old.(*bearer.Bearer)
		c.Set(consts.Authorization, jwt)
		return nil
	}
	jtw := bearer.NewBearer(config.Config.Secretkey, 0, "")
	err := jtw.Parse(token)
	if err != nil {
		niuhe.LogInfo("%v", err)
		return niuhe.NewCommError(consts.AuthError, err.Error())
	}
	c.Set(consts.Authorization, jtw)
	proto.setCache(jtw, 5*time.Minute, consts.Authorization, token)
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
	rspInst := rsp.Interface()
	if _, ok := rspInst.(isCustomRoot); ok {
		c.JSON(http.StatusOK, rspInst)
	} else {
		var response map[string]interface{}
		if err != nil {
			if commErr, ok := err.(niuhe.ICommError); ok {
				if commErr.GetCode() == consts.CodeNoCommRsp {
					// 已经处理了返回，不需要再处理
					return nil
				}
				response = map[string]interface{}{
					"result":  commErr.GetCode(),
					"message": commErr.GetMessage(),
				}
				if commErr.GetCode() == 0 {
					response["data"] = rsp.Interface()
				}
			} else {
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
