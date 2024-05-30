package daos

import (
	"fmt"

	"github.com/ma-guo/admin-core/xorm/models"

	"github.com/ma-guo/niuhe"
)

type _DictDao struct {
	*Dao
}

func (dao *Dao) Dict() *_DictDao {
	return &_DictDao{Dao: dao}
}

func (dao *_DictDao) GetById(id int64) (*models.SysDict, error) {
	item := &models.SysDict{
		Id: id,
	}
	has, err := dao.db().Get(item)
	if err != nil {
		return nil, err
	} else if !has || id == 0 {
		return nil, fmt.Errorf("do not have this dict id: %v", id)
	}
	return item, nil
}

func (dao *_DictDao) GetByName(typeCode, name string) (*models.SysDict, bool, error) {
	item := &models.SysDict{
		TypeCode: typeCode,
		Name:     name,
	}
	has, err := dao.db().Get(item)
	if err != nil {
		return nil, false, err
	}
	return item, has, nil
}

func (dao *_DictDao) GetByIds(ids ...int64) ([]*models.SysDict, error) {
	rows := make([]*models.SysDict, 0)
	err := dao.db().In("`id`", ids).Find(&rows)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (dao *_DictDao) Insert(row *models.SysDict) (bool, error) {
	affected, err := dao.db().Insert(row)
	return affected > 0, err
}

func (dao *_DictDao) Update(id int64, row *models.SysDict) (bool, error) {
	affected, err := dao.db().Where("id=?", id).AllCols().Update(row)
	return affected > 0, err
}

func (dao *_DictDao) Has(row *models.SysDict) (int64, error) {
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

func (dao *_DictDao) Delete(row *models.SysDict) (bool, error) {
	affected, err := dao.db().Delete(row)
	return affected > 0, err
}

func (dao *_DictDao) GetAll() ([]*models.SysDict, error) {
	rows := make([]*models.SysDict, 0)
	err := dao.db().Asc("`sort`").Find(&rows)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (dao *_DictDao) GetByCode(typeCode string) ([]*models.SysDict, error) {
	rows := make([]*models.SysDict, 0)
	err := dao.db().Asc("`sort`").Find(&rows, &models.SysDict{
		TypeCode: typeCode,
	})
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (dao *_DictDao) GetPage(keywords, typeCode string, page, size int) ([]*models.SysDict, int64, error) {
	rows := make([]*models.SysDict, 0)
	session := dao.db()
	if keywords != "" {
		session.Where("`name` LIKE ?", "%"+keywords+"%")
	}
	if typeCode != "" {
		session.Where("`type_code` LIKE ?", "%"+typeCode+"%")
	}

	total, err := session.Limit(size, page*size-size).FindAndCount(&rows)
	if err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}
