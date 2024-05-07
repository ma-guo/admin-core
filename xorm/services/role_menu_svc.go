package services

import (
	"fmt"

	"github.com/ma-guo/admin-core/xorm/models"

	"github.com/ziipin-server/niuhe"
)

type _RooleMenuSvc struct {
	*_Svc
}

func (svc *_Svc) RoleMenu() *_RooleMenuSvc {
	return &_RooleMenuSvc{svc}
}

func (svc *_RooleMenuSvc) GetById(rid, mid int64) (*models.SysRoleMenu, error) {
	if mid <= 0 || rid <= 0 {
		return nil, fmt.Errorf("角色信息不存在")
	}
	return svc.dao().RoleMenu().GetBy(rid, mid)
}

func (svc *_RooleMenuSvc) Delete(rid, mid int64) error {
	_, err := svc.dao().RoleMenu().Delete(&models.SysRoleMenu{RoleId: rid, MenuId: mid})
	return err
}

func (svc *_RooleMenuSvc) Add(row *models.SysRoleMenu) (*models.SysRoleMenu, error) {
	has := svc.Has(row)
	if has {
		return row, nil
	}
	_, err := svc.dao().RoleMenu().Insert(row)
	if err != nil {
		return nil, err
	}
	return row, nil
}

func (svc *_RooleMenuSvc) Has(row *models.SysRoleMenu) bool {
	has, _ := svc.dao().RoleMenu().Has(row)
	return has
}

func (svc *_RooleMenuSvc) Update(rid int64, menuIds ...int64) error {
	rows, err := svc.dao().RoleMenu().GetRoleMenus(rid)
	if err != nil {
		niuhe.LogInfo("%v", err)
		return err
	}
	data := make(map[int64]*models.SysRoleMenu, 0)
	for _, row := range rows {
		data[row.MenuId] = row
	}
	// 添加没有的
	for _, mid := range menuIds {
		if _, ok := data[mid]; ok {
			continue
		}
		niuhe.LogInfo("add role_menu: %v, %v", rid, mid)
		_, err := svc.Add(&models.SysRoleMenu{RoleId: rid, MenuId: mid})
		if err != nil {
			niuhe.LogInfo("%v", err)
			return err
		}
	}
	// 删除多余的
	for _, row := range rows {
		has := false
		for _, mid := range menuIds {

			if row.MenuId == mid {
				has = true
				break
			}
		}
		if !has {
			niuhe.LogInfo("delete role_menu: %v, %v", rid, row.MenuId)
			err := svc.Delete(rid, row.MenuId)
			if err != nil {
				niuhe.LogInfo("%v", err)
				return err
			}
		}
	}
	return nil
}

// 获取角色对应的菜单id
func (svc *_RooleMenuSvc) GetRoleMenus(rid ...int64) ([]*models.SysRoleMenu, error) {
	return svc.dao().RoleMenu().GetRoleMenus(rid...)
}

func (svc *_RooleMenuSvc) GetAll() ([]*models.SysRoleMenu, error) {
	return svc.dao().RoleMenu().GetAll()
}
