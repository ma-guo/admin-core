package daos

import (
	"fmt"

	"github.com/ma-guo/admin-core/app/v1/protos"
	"github.com/ma-guo/admin-core/xorm/models"
)

type _UserDao struct {
	*Dao
}

func (dao *Dao) User() *_UserDao {
	return &_UserDao{Dao: dao}
}

func (dao *_UserDao) GetBy(item *models.SysUser) (*models.SysUser, bool, error) {
	has, err := dao.db().Get(item)
	if err != nil {
		return nil, false, err
	}
	return item, has, nil
}

func (dao *_UserDao) Insert(row *models.SysUser) (bool, error) {
	affected, err := dao.db().Insert(row)
	return affected > 0, err
}

func (dao *_UserDao) Update(id int64, row *models.SysUser) (bool, error) {
	affected, err := dao.db().Where("id=?", id).AllCols().Update(row)
	return affected > 0, err
}

func (dao *_UserDao) Has(row *models.SysUser) (bool, error) {
	if row.Id > 0 {
		_, has, err := dao.GetBy(&models.SysUser{Id: row.Id})
		if err != nil {
			return false, err
		}
		return has, nil
	}
	return false, fmt.Errorf("user id is invalid")
}

func (dao *_UserDao) GetAll() ([]*models.SysUser, error) {
	rows := make([]*models.SysUser, 0)
	err := dao.db().Find(&rows)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (dao *_UserDao) GetPage(param *protos.V1UsersPageReq, hasStatus bool, deptIds ...int64) ([]*models.SysUser, int64, error) {
	rows := make([]*models.SysUser, 0)
	keyword := "%" + param.Keywords + "%"
	session := dao.db().Where("`deleted`=?", 0)
	if hasStatus {
		session.And("`status`=?", param.Status)
	}
	if param.Keywords != "" {
		session.And("`username` LIKE ? OR `nickname` LIKE ? OR `mobile` LIKE ?", keyword, keyword, keyword)
	}
	if len(deptIds) > 0 {
		session.In("`dept_id`", deptIds)
	}
	if param.StartTime != "" {
		session.And("`create_time` >= ?", param.StartTime)
	}
	if param.EndTime != "" {
		session.And("`create_time` <= ?", param.EndTime)
	}

	total, err := session.Limit(param.PageSize, param.PageNum*param.PageSize-param.PageSize).FindAndCount(&rows)
	if err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}
