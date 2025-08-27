package services

import (
	"fmt"

	"github.com/ma-guo/admin-core/xorm/models"

	"github.com/ma-guo/niuhe"
)

type _RoleMenuSvc struct {
	*_Svc
}

func (svc *_Svc) RoleMenu() *_RoleMenuSvc {
	return &_RoleMenuSvc{svc}
}

func (svc *_RoleMenuSvc) GetById(rid, mid int64) (*models.SysRoleMenu, error) {
	if mid <= 0 || rid <= 0 {
		return nil, fmt.Errorf("角色信息不存在")
	}
	return svc.dao().RoleMenu().GetBy(rid, mid)
}

func (svc *_RoleMenuSvc) Delete(rid, mid int64) error {
	_, err := svc.dao().RoleMenu().Delete(&models.SysRoleMenu{RoleId: rid, MenuId: mid})
	return err
}

func (svc *_RoleMenuSvc) Add(row *models.SysRoleMenu) (*models.SysRoleMenu, error) {
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

func (svc *_RoleMenuSvc) Has(row *models.SysRoleMenu) bool {
	has, _ := svc.dao().RoleMenu().Has(row)
	return has
}

func (svc *_RoleMenuSvc) Update(rid int64, menuIds ...int64) error {
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
func (svc *_RoleMenuSvc) GetRoleMenus(rid ...int64) ([]*models.SysRoleMenu, error) {
	return svc.dao().RoleMenu().GetRoleMenus(rid...)
}

func (svc *_RoleMenuSvc) GetAll() ([]*models.SysRoleMenu, error) {
	return svc.dao().RoleMenu().GetAll()
}

// / 根据用户 id 获取对应的按钮
func (svc *_RoleMenuSvc) GetMenusByUid(uid int64) (map[int64]int64, error) {
	prefix := "menusByUid"
	cache, has := svc.dao().GetCache(prefix, uid)
	if has {
		return cache.(map[int64]int64), nil
	}
	roles, err := svc.Role().GetByUserId(uid)
	if err != nil {
		niuhe.LogInfo("%v", err)
		return nil, err
	}
	roleIds := make([]int64, 0)
	for _, role := range roles {
		roleIds = append(roleIds, role.Id)
	}
	roleMenus, err := svc.RoleMenu().GetRoleMenus(roleIds...)
	if err != nil {
		niuhe.LogInfo("%v", err)
		return nil, err
	}
	tmp := make(map[int64]int64, 0)
	for _, row := range roleMenus {
		tmp[row.MenuId] = row.MenuId
	}
	svc.dao().SetCache(tmp, prefix, uid)

	return tmp, nil
}
