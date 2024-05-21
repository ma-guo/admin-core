package daos

import (
	"fmt"

	"github.com/ma-guo/admin-core/xorm/models"

	"github.com/ma-guo/niuhe"
)

type _RoleDao struct {
	*Dao
}

func (dao *Dao) Role() *_RoleDao {
	return &_RoleDao{Dao: dao}
}

func (dao *_RoleDao) GetById(rid int64) (*models.SysRole, error) {
	item := &models.SysRole{
		Id: rid,
	}
	has, err := dao.db().Get(item)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, fmt.Errorf("do not have this role %v", rid)
	}
	return item, nil
}
func (dao *_RoleDao) GetByName(name string) (*models.SysRole, error) {
	item := &models.SysRole{
		Name: name,
	}
	has, err := dao.db().Get(item)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, fmt.Errorf("do not have this role %v", name)
	}
	return item, nil
}
func (dao *_RoleDao) GetByNameCode(name, code string) (*models.SysRole, error) {
	item := &models.SysRole{
		Name: name,
		Code: code,
	}
	has, err := dao.db().Get(item)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, fmt.Errorf("do not have this role %v, %v", name, code)
	}
	return item, nil
}

func (dao *_RoleDao) GetByIds(ids ...int64) ([]*models.SysRole, error) {
	rows := make([]*models.SysRole, 0)
	err := dao.db().In("`id`", ids).Find(&rows)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (dao *_RoleDao) Insert(row *models.SysRole) (bool, error) {
	affected, err := dao.db().Insert(row)
	return affected > 0, err
}

func (dao *_RoleDao) Update(id int64, row *models.SysRole) (bool, error) {
	affected, err := dao.db().Where("`id`=?", row.Id).AllCols().Update(row)
	return affected > 0, err
}

func (dao *_RoleDao) Has(row *models.SysRole) (int64, error) {
	niuhe.LogInfo("Has row: %v", row)
	if row.Id > 0 {
		tmp, err := dao.GetById(row.Id)
		if err != nil {
			return tmp.Id, err
		}
		return tmp.Id, nil
	}
	return row.Id, nil
}

func (dao *_RoleDao) GetAll() ([]*models.SysRole, error) {
	rows := make([]*models.SysRole, 0)
	err := dao.db().Find(&rows)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (dao *_RoleDao) GetPage(keywords string, page, size int) ([]*models.SysRole, int64, error) {
	rows := make([]*models.SysRole, 0)
	session := dao.db().Where("`deleted`=?", 0)
	if keywords != "" {
		session.And("`name` LIKE ?", "%"+keywords+"%")
	}
	total, err := session.Limit(size, page*size-size).FindAndCount(&rows)
	if err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}
