package services

import (
	"fmt"
	"time"

	"github.com/ma-guo/admin-core/xorm/models"
	"github.com/ma-guo/niuhe"
)

type _ApiSvc struct {
	*_Svc
}

func (svc *_Svc) Api() *_ApiSvc {
	return &_ApiSvc{svc}
}

func (svc *_ApiSvc) GetById(id int64) (*models.SysApi, error) {
	if id <= 0 {
		return nil, fmt.Errorf("api 不存在")
	}
	return svc.dao().Api().GetById(id)
}

func (svc *_ApiSvc) Update(id int64, row *models.SysApi) error {
	_, err := svc.dao().Api().Update(id, row)
	return err
}

func (svc *_ApiSvc) Add(row *models.SysApi) (*models.SysApi, error) {
	item, has, err := svc.Has(row)
	if err != nil {
		return nil, err
	}
	if has {
		row.Id = item.Id
		_, err := svc.dao().Api().Update(row.Id, row)
		if err != nil {
			return nil, err
		}
		return row, nil
	}
	_, err = svc.dao().Api().Insert(row)
	if err != nil {
		return nil, err
	}
	return row, nil
}

// 插入一条数据
func (svc *_ApiSvc) Insert(row *models.SysApi) error {
	_, err := svc.dao().Api().Insert(row)
	return err
}

func (svc *_ApiSvc) Has(row *models.SysApi) (*models.SysApi, bool, error) {
	return svc.dao().Api().GetByMethod(row.Method, row.Path)
}

func (svc *_ApiSvc) GetAll() ([]*models.SysApi, error) {
	return svc.dao().Api().GetAll()
}

func (svc *_ApiSvc) GetByMethod(methd string, path string) (*models.SysApi, bool, error) {
	return svc.dao().Api().GetByMethod(methd, path)
}

func (svc *_ApiSvc) GetByIds(ids ...int64) (map[int64]*models.SysApi, error) {
	data := make(map[int64]*models.SysApi, 0)
	rows, err := svc.dao().Api().GetByIds(ids...)
	if err != nil {
		return nil, err
	}
	for _, row := range rows {
		data[row.Id] = row
	}
	return data, nil
}

func (svc *_ApiSvc) GetPage(keyword string, page, size int) ([]*models.SysApi, int64, error) {
	return svc.dao().Api().GetPage(keyword, page, size)
}

// 返回 name 为 key 的 map [key]value
func (svc *_ApiSvc) GetByVendorToMap(keyword string) (map[string]string, error) {
	rows, _, err := svc.dao().Api().GetPage(keyword, 1, 100)
	if err != nil {
		return nil, err
	}
	data := make(map[string]string)
	for _, row := range rows {
		data[row.GetKey()] = row.Path
	}
	return data, nil
}

func (svc *_ApiSvc) GetMenus(method, path string) ([]*models.SysMenu, error) {
	prefix := "apiGetMenus"
	cache, has := svc.dao().GetCache(prefix, method, path)
	if has {
		return cache.([]*models.SysMenu), nil
	}
	api, has, err := svc.GetByMethod(method, path)
	if err != nil {
		niuhe.LogInfo("%v", err)
		return nil, err
	}
	if !has {
		// 没有配置的路径不检查权限, 返回空列表
		return []*models.SysMenu{}, nil
	}
	if len(api.MenuIds) == 0 {
		// 没有配置的路径不检查权限, 返回空列表
		return []*models.SysMenu{}, nil
	}
	menus, err := svc.Menu().GetByIds(api.MenuIds...)
	if err != nil {
		niuhe.LogInfo("%v", err)
		return nil, err
	}
	svc.dao().SetCache(menus, 1*time.Minute, prefix, method, path)
	return menus, nil
}
