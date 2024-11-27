package services

import (
	"fmt"
	"strings"

	"github.com/ma-guo/admin-core/app/common/consts"
	"github.com/ma-guo/admin-core/xorm/models"
)

type _MenuSvc struct {
	*_Svc
}

func (svc *_Svc) Menu() *_MenuSvc {
	return &_MenuSvc{svc}
}

func (svc *_MenuSvc) GetById(id int64) (*models.SysMenu, error) {
	if id <= 0 {
		return nil, fmt.Errorf("菜单信息不存在")
	}
	row, has, err := svc.dao().Menu().GetBy(&models.SysMenu{Id: id})
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, fmt.Errorf("菜单信息不存在")
	}
	return row, nil
}

func (svc *_MenuSvc) GetByIds(ids ...int64) ([]*models.SysMenu, error) {
	return svc.dao().Menu().GetByIds(ids...)
}
func (svc *_MenuSvc) GetByName(name string) (*models.SysMenu, bool, error) {
	if name == "" {
		return nil, false, fmt.Errorf("菜单信息不存在")
	}
	return svc.dao().Menu().GetBy(&models.SysMenu{Name: name})
}

func (svc *_MenuSvc) Update(id int64, row *models.SysMenu) error {
	_, err := svc.dao().Menu().Update(id, row)
	return err
}

func (svc *_MenuSvc) Delete(mid int64) error {
	_, err := svc.dao().Menu().Delete(&models.SysMenu{Id: mid})
	return err
}

func (svc *_MenuSvc) Insert(row *models.SysMenu) (*models.SysMenu, error) {
	if row.Type == consts.MenuTypeGroup.CATALOG.Value {
		if row.Component == "" {
			row.Component = "Layout"
		}
	}
	_, err := svc.dao().Menu().Insert(row)
	if err != nil {
		return nil, err
	}
	return row, nil
}

func (svc *_MenuSvc) Has(row *models.SysMenu) bool {
	_, has, _ := svc.dao().Menu().GetBy(&models.SysMenu{Id: row.Id, Name: row.Name})
	return has
}

func (svc *_MenuSvc) GetAll(keyword string) ([]*models.SysMenu, error) {
	return svc.dao().Menu().GetAll(strings.Trim(keyword, " "))
}

func (svc *_MenuSvc) GetPage(keyword string, menuType, page, size int) ([]*models.SysMenu, int64, error) {
	return svc.dao().Menu().GetPage(keyword, menuType, page, size)
}
