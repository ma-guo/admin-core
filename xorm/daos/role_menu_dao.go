package daos

import (
	"fmt"

	"github.com/ma-guo/admin-core/xorm/models"
)

type _RoleMenuDao struct {
	*Dao
}

func (dao *Dao) RoleMenu() *_RoleMenuDao {
	return &_RoleMenuDao{Dao: dao}
}

func (dao *_RoleMenuDao) GetBy(rid, mid int64) (*models.SysRoleMenu, error) {
	item := &models.SysRoleMenu{
		RoleId: rid,
		MenuId: mid,
	}
	has, err := dao.db().Get(item)
	if err != nil {
		return nil, err
	} else if !has || rid == 0 || mid == 0 {
		return nil, fmt.Errorf("do not have this role_menu(role:%v, menu:%v)", rid, mid)
	}
	return item, nil
}

func (dao *_RoleMenuDao) Insert(row *models.SysRoleMenu) (bool, error) {
	affected, err := dao.db().Insert(row)
	return affected > 0, err
}

func (dao *_RoleMenuDao) Delete(row *models.SysRoleMenu) (bool, error) {
	affected, err := dao.db().Delete(row)
	return affected > 0, err
}

func (dao *_RoleMenuDao) Has(row *models.SysRoleMenu) (bool, error) {
	_, err := dao.GetBy(row.RoleId, row.MenuId)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (dao *_RoleMenuDao) GetRoleMenus(rid ...int64) ([]*models.SysRoleMenu, error) {
	rows := make([]*models.SysRoleMenu, 0)
	err := dao.db().In("`role_id`", rid).Find(&rows)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (dao *_RoleMenuDao) GetAll() ([]*models.SysRoleMenu, error) {
	rows := make([]*models.SysRoleMenu, 0)
	err := dao.db().Find(&rows)
	if err != nil {
		return nil, err
	}
	return rows, nil
}
