package services

import (
	"fmt"

	"github.com/ma-guo/admin-core/xorm/models"

	"github.com/ma-guo/niuhe"
)

type _RoleSvc struct {
	*_Svc
}

func (svc *_Svc) Role() *_RoleSvc {
	return &_RoleSvc{svc}
}

func (svc *_RoleSvc) GetById(id int64) (*models.SysRole, error) {
	if id <= 0 {
		return nil, fmt.Errorf("部门信息不存在")
	}
	return svc.dao().Role().GetById(id)
}
func (svc *_RoleSvc) GetByIds(ids ...int64) ([]*models.SysRole, error) {
	return svc.dao().Role().GetByIds(ids...)
}

func (svc *_RoleSvc) Update(id int64, row *models.SysRole) error {
	_, err := svc.dao().Role().Update(id, row)
	return err
}

func (svc *_RoleSvc) Add(row *models.SysRole) (*models.SysRole, error) {
	id := svc.Has(row)
	if id > 0 {
		row.Id = id
		_, err := svc.dao().Role().Update(row.Id, row)
		if err != nil {
			niuhe.LogInfo("%v", err)
			return nil, err
		}
		return row, nil
	}
	_, err := svc.dao().Role().Insert(row)
	if err != nil {
		niuhe.LogInfo("%v", err)
		return nil, err
	}
	return row, nil
}
func (svc *_RoleSvc) GetByName(roleName string) (*models.SysRole, error) {
	return svc.dao().Role().GetByName(roleName)
}

func (svc *_RoleSvc) Has(row *models.SysRole) int64 {
	role, _ := svc.dao().Role().GetByNameCode(row.Name, row.Code)
	if role != nil && role.Id > 0 {
		return role.Id
	}
	return 0
}

func (svc *_RoleSvc) GetAll() ([]*models.SysRole, error) {
	return svc.dao().Role().GetAll()
}

// 根据 userId 获取角色列表
func (svc *_RoleSvc) GetByUserId(uid int64) ([]*models.SysRole, error) {
	roles, err := svc.UserRole().GetAll(uid)
	if err != nil {
		return nil, err
	}
	ids := make([]int64, 0)
	for _, role := range roles {
		ids = append(ids, role.RoleId)
	}
	return svc.GetByIds(ids...)
}

func (svc *_RoleSvc) GetPage(keywords string, page, size int) ([]*models.SysRole, int64, error) {
	return svc.dao().Role().GetPage(keywords, page, size)
}
