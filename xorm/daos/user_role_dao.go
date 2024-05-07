package daos

import (
	"fmt"

	"github.com/ma-guo/admin-core/xorm/models"
)

type _UserRoleDao struct {
	*Dao
}

func (dao *Dao) UserRole() *_UserRoleDao {
	return &_UserRoleDao{Dao: dao}
}

func (dao *_UserRoleDao) GetBy(roleId, userId int64) (*models.SysUserRole, error) {
	item := &models.SysUserRole{
		RoleId: roleId,
		UserId: userId,
	}
	has, err := dao.db().Get(item)
	if err != nil {
		return nil, err
	} else if !has || roleId == 0 || userId == 0 {
		return nil, fmt.Errorf("do not have this user_role(user:%v, role:%v)", userId, roleId)
	}
	return item, nil
}

func (dao *_UserRoleDao) Insert(row *models.SysUserRole) (bool, error) {
	affected, err := dao.db().Insert(row)
	return affected > 0, err
}

func (dao *_UserRoleDao) Delete(row *models.SysUserRole) (bool, error) {
	affected, err := dao.db().Delete(row)
	return affected > 0, err
}

func (dao *_UserRoleDao) Has(row *models.SysUserRole) (bool, error) {
	_, err := dao.GetBy(row.UserId, row.RoleId)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (dao *_UserRoleDao) GetAll(userId int64) ([]*models.SysUserRole, error) {
	rows := make([]*models.SysUserRole, 0)
	err := dao.db().Where("`user_id`=?", userId).Find(&rows)
	if err != nil {
		return nil, err
	}
	return rows, nil
}
