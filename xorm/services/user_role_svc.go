package services

import (
	"fmt"

	"github.com/ma-guo/admin-core/xorm/models"

	"github.com/ma-guo/niuhe"
)

type _UserRoleSvc struct {
	*_Svc
}

func (svc *_Svc) UserRole() *_UserRoleSvc {
	return &_UserRoleSvc{svc}
}

func (svc *_UserRoleSvc) GetById(uid, rid int64) (*models.SysUserRole, error) {
	if uid <= 0 || rid <= 0 {
		return nil, fmt.Errorf("角色信息不存在")
	}
	return svc.dao().UserRole().GetBy(uid, rid)
}

func (svc *_UserRoleSvc) Delete(uid, rid int64) error {
	_, err := svc.dao().UserRole().Delete(&models.SysUserRole{UserId: uid, RoleId: rid})
	return err
}

func (svc *_UserRoleSvc) Add(row *models.SysUserRole) (*models.SysUserRole, error) {
	has := svc.Has(row)
	if has {
		return row, nil
	}
	_, err := svc.dao().UserRole().Insert(row)
	if err != nil {
		return nil, err
	}
	return row, nil
}

func (svc *_UserRoleSvc) Has(row *models.SysUserRole) bool {
	has, _ := svc.dao().UserRole().Has(row)
	return has
}

func (svc *_UserRoleSvc) GetAll(userId int64) ([]*models.SysUserRole, error) {
	return svc.dao().UserRole().GetAll(userId)
}

func (svc *_UserRoleSvc) Update(userId int64, roleIds ...int64) error {
	rows, err := svc.dao().UserRole().GetAll(userId)
	if err != nil {
		niuhe.LogInfo("%v", err)
		return err
	}
	data := make(map[int64]*models.SysUserRole, 0)
	for _, row := range rows {
		data[row.RoleId] = row
	}
	// 添加没有的
	for _, rid := range roleIds {
		if _, ok := data[rid]; ok {
			continue
		}
		_, err := svc.Add(&models.SysUserRole{UserId: userId, RoleId: rid})
		if err != nil {
			niuhe.LogInfo("%v", err)
			return err
		}
	}
	// 删除多余的
	for _, row := range rows {
		has := false
		for _, id := range roleIds {

			if row.RoleId == id {
				has = true
				break
			}
		}
		if !has {
			err := svc.Delete(userId, row.RoleId)
			if err != nil {
				niuhe.LogInfo("%v", err)
				return err
			}
		}
	}
	return nil
}
