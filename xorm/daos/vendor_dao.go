package daos

import (
	"fmt"

	"github.com/ma-guo/admin-core/xorm/models"

	"github.com/ma-guo/niuhe"
)

type _VendorDao struct {
	*Dao
}

func (dao *Dao) Vendor() *_VendorDao {
	return &_VendorDao{Dao: dao}
}

func (dao *_VendorDao) GetById(id int64) (*models.SysVendor, error) {
	item := &models.SysVendor{
		Id: id,
	}
	has, err := dao.db().Get(item)
	if err != nil {
		return nil, err
	} else if !has || id == 0 {
		return nil, fmt.Errorf("do not have this vendor id: %v", id)
	}
	return item, nil
}

func (dao *_VendorDao) GetByKey(vendor, key string) (*models.SysVendor, bool, error) {
	item := &models.SysVendor{
		Vendor: vendor,
		Key:    key,
	}
	has, err := dao.db().Get(item)
	if err != nil {
		return nil, false, err
	}
	return item, has, nil
}

func (dao *_VendorDao) GetByIds(ids ...int64) ([]*models.SysVendor, error) {
	rows := make([]*models.SysVendor, 0)
	err := dao.db().In("`id`", ids).Find(&rows)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (dao *_VendorDao) Insert(row *models.SysVendor) (bool, error) {
	affected, err := dao.db().Insert(row)
	return affected > 0, err
}

func (dao *_VendorDao) Update(id int64, row *models.SysVendor) (bool, error) {
	affected, err := dao.db().Where("id=?", id).Update(row)
	return affected > 0, err
}

func (dao *_VendorDao) Has(row *models.SysVendor) (int64, error) {
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

func (dao *_VendorDao) GetAll() ([]*models.SysVendor, error) {
	rows := make([]*models.SysVendor, 0)
	err := dao.db().Asc("`sort`").Find(&rows)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (dao *_VendorDao) GetPage(vendor string, page, size int) ([]*models.SysVendor, int64, error) {
	rows := make([]*models.SysVendor, 0)
	session := dao.db().Where("`vendor`= ?", vendor)
	total, err := session.Limit(size, page*size-size).FindAndCount(&rows)
	if err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}
