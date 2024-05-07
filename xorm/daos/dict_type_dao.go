package daos

import (
	"fmt"

	"github.com/ma-guo/admin-core/xorm/models"

	"github.com/ziipin-server/niuhe"
)

type _DictTypeDao struct {
	*Dao
}

func (dao *Dao) DictType() *_DictTypeDao {
	return &_DictTypeDao{Dao: dao}
}

func (dao *_DictTypeDao) GetById(id int64) (*models.SysDictType, error) {
	item := &models.SysDictType{
		Id: id,
	}
	has, err := dao.db().Get(item)
	if err != nil {
		return nil, err
	} else if !has || id == 0 {
		return nil, fmt.Errorf("do not have this dict_type id: %v", id)
	}
	return item, nil
}

func (dao *_DictTypeDao) GetByIds(ids ...int64) ([]*models.SysDictType, error) {
	rows := make([]*models.SysDictType, 0)
	err := dao.db().In("`id`", ids).Find(&rows)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (dao *_DictTypeDao) Insert(row *models.SysDictType) (bool, error) {
	affected, err := dao.db().Insert(row)
	return affected > 0, err
}

func (dao *_DictTypeDao) Update(id int64, row *models.SysDictType) (bool, error) {
	affected, err := dao.db().Where("id=?", id).Update(row)
	return affected > 0, err
}

func (dao *_DictTypeDao) Has(row *models.SysDictType) (int64, error) {
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

func (dao *_DictTypeDao) Delete(row *models.SysDictType) (bool, error) {
	affected, err := dao.db().Delete(row)
	return affected > 0, err
}

func (dao *_DictTypeDao) GetByCode(code string) (*models.SysDictType, bool, error) {
	item := &models.SysDictType{
		Code: code,
	}
	has, err := dao.db().Get(item)
	if err != nil {
		return nil, false, err
	}
	return item, has, nil
}

func (dao *_DictTypeDao) GetAll() ([]*models.SysDictType, error) {
	rows := make([]*models.SysDictType, 0)
	err := dao.db().Asc("`sort`").Find(&rows)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (dao *_DictTypeDao) GetPage(keywords string, page, size int) ([]*models.SysDictType, int64, error) {
	rows := make([]*models.SysDictType, 0)
	session := dao.db()
	if keywords != "" {
		session.Where("`name` LIKE ? OR `code` LIKE ?", "%"+keywords+"%", "%"+keywords+"%")
	}

	total, err := session.Limit(size, page*size-size).FindAndCount(&rows)
	if err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}
