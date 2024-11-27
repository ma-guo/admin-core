package views

// Generated by niuhe.idl

import (
	"net/http"
	"time"

	"github.com/ma-guo/admin-core/app/common/consts"
	"github.com/ma-guo/admin-core/app/v1/protos"
	"github.com/ma-guo/admin-core/xorm/models"
	"github.com/ma-guo/admin-core/xorm/services"

	"github.com/ma-guo/niuhe"
)

type Api struct {
	_Gen_Api
}

// API列表
func (v *Api) Page_GET(c *niuhe.Context, req *protos.V1ApiPageReq, rsp *protos.V1ApiPageRsp) error {
	svc := services.NewSvc()
	defer svc.Close()
	rows, total, err := svc.Api().GetPage(req.Keyword, req.PageNum, req.PageSize)
	if err != nil {
		niuhe.LogInfo("%v", err)
		return err
	}
	rsp.Button = consts.MenuTypeGroup.BUTTON.Value
	rsp.Total = total
	rsp.Items = make([]*protos.V1ApiItem, 0)
	menuids := make([]int64, 0)
	for _, row := range rows {
		menuids = append(menuids, row.MenuIds...)
	}
	menus, err := svc.Menu().GetByIds(menuids...)
	if err != nil {
		niuhe.LogInfo("%v", err)
		return err
	}
	menusMap := make(map[int64]*models.SysMenu)
	for _, menu := range menus {
		menusMap[menu.Id] = menu
	}
	for _, row := range rows {
		item := &protos.V1ApiItem{
			Id:         row.Id,
			Method:     row.Method,
			Name:       row.Name,
			Path:       row.Path,
			Menus:      []string{},
			Remark:     row.Remark,
			UpdateTime: row.UpdateTime.Format(time.DateTime),
		}
		for _, menuId := range row.MenuIds {
			if menu, has := menusMap[menuId]; has {
				item.Menus = append(item.Menus, menu.Name)
			}
		}
		rsp.Items = append(rsp.Items, item)
	}
	return nil
}

// 获取API详情值
func (v *Api) Form_GET(c *niuhe.Context, req *protos.V1ApiFormReq, rsp *protos.V1ApiFormRsp) error {
	svc := services.NewSvc()
	defer svc.Close()
	row, err := svc.Api().GetById(req.Id)
	if err != nil {
		niuhe.LogInfo("%v", err)
		return err
	}
	rsp.Id = row.Id
	rsp.Method = row.Method
	rsp.Name = row.Name
	rsp.Path = row.Path
	rsp.Remark = row.Remark

	rsp.Methods = []string{
		http.MethodGet,
		http.MethodPost,
		http.MethodPut,
		http.MethodDelete,
		http.MethodHead,
		http.MethodPatch,
		http.MethodConnect,
		http.MethodOptions,
	}
	rsp.Menus = make([]*protos.V1MenuTiny, 0)
	if len(row.MenuIds) > 0 {
		menus, err := svc.Menu().GetByIds(row.MenuIds...)
		if err != nil {
			niuhe.LogInfo("%v", err)
			return err
		}
		for _, menu := range menus {
			rsp.Menus = append(rsp.Menus, &protos.V1MenuTiny{
				Id:   menu.Id,
				Name: menu.Name,
				Perm: menu.Perm,
			})
		}
	}
	return nil
}

// 更新配置项
func (v *Api) Update_POST(c *niuhe.Context, req *protos.V1ApiUpdateReq, rsp *protos.NoneRsp) error {
	svc := services.NewSvc()
	defer svc.Close()
	row := &models.SysApi{
		Id:      req.Id,
		Method:  req.Method,
		Name:    req.Name,
		Path:    req.Path,
		Remark:  req.Remark,
		MenuIds: req.Menus,
	}
	err := svc.Api().Update(row.Id, row)
	if err != nil {
		niuhe.LogInfo("%v", err)
		return err
	}
	return nil
}
func init() {
	GetModule().Register(&Api{})
}
