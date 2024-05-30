package daos

import (
	"fmt"

	"github.com/ma-guo/admin-core/xorm/models"

	"github.com/ma-guo/niuhe"
)

type _DeptDao struct {
	*Dao
}

func (dao *Dao) Dept() *_DeptDao {
	return &_DeptDao{Dao: dao}
}

func (dao *_DeptDao) GetById(id int64) (*models.SysDept, error) {
	item := &models.SysDept{
		Id: id,
	}
	has, err := dao.db().Get(item)
	if err != nil {
		return nil, err
	} else if !has || id == 0 {
		return nil, fmt.Errorf("do not have this dept id: %v", id)
	}
	return item, nil
}

func (dao *_DeptDao) GetByIds(ids ...int64) ([]*models.SysDept, error) {
	rows := make([]*models.SysDept, 0)
	err := dao.db().In("`id`", ids).Find(&rows)
	if err != nil {
		return nil, err
	}
	return rows, nil
}
func (dao *_DeptDao) GetChildren(id int64) ([]*models.SysDept, error) {
	rows := make([]*models.SysDept, 0)
	err := dao.db().Find(&rows, &models.SysDept{ParentId: id})
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (dao *_DeptDao) Insert(row *models.SysDept) (bool, error) {
	affected, err := dao.db().Insert(row)
	return affected > 0, err
}

func (dao *_DeptDao) Update(id int64, row *models.SysDept) (bool, error) {
	affected, err := dao.db().Where("id=?", id).AllCols().Update(row)
	return affected > 0, err
}

func (dao *_DeptDao) Has(row *models.SysDept) (int64, error) {
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

func (dao *_DeptDao) Delete(row *models.SysDept) (bool, error) {
	affected, err := dao.db().Delete(row)
	return affected > 0, err
}

func (dao *_DeptDao) GetByName(name string) (*models.SysDept, error) {
	item := &models.SysDept{
		Name: name,
	}
	has, err := dao.db().Get(item)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, fmt.Errorf("do not have this name: %s", name)
	}
	return item, nil
}

func (dao *_DeptDao) GetAll() ([]*models.SysDept, error) {
	rows := make([]*models.SysDept, 0)
	err := dao.db().Asc("`sort`").Find(&rows)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (dao *_DeptDao) Search(name string, status int) ([]*models.SysDept, error) {
	rows := make([]*models.SysDept, 0)
	session := dao.db()
	if name != "" {
		session = session.Where("`name` like ?", "%"+name+"%")
	}
	if status >= 0 {
		session = session.And("`status`=?", status)
	}
	err := session.Asc("`sort`").Find(&rows)
	if err != nil {
		return nil, err
	}
	return rows, nil
}
