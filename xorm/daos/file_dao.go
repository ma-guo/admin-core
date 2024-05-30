package daos

import (
	"fmt"

	"github.com/ma-guo/admin-core/xorm/models"

	"github.com/ma-guo/niuhe"
)

type _FileDao struct {
	*Dao
}

func (dao *Dao) File() *_FileDao {
	return &_FileDao{Dao: dao}
}

func (dao *_FileDao) GetById(id int64) (*models.SysFile, error) {
	item := &models.SysFile{
		Id: id,
	}
	has, err := dao.db().Get(item)
	if err != nil {
		return nil, err
	} else if !has || id == 0 {
		return nil, fmt.Errorf("do not have this file id: %v", id)
	}
	return item, nil
}

func (dao *_FileDao) GetByPath(path string) (*models.SysFile, bool, error) {
	item := &models.SysFile{
		Path: path,
	}
	has, err := dao.db().Get(item)
	if err != nil {
		return nil, false, err
	}
	return item, has, nil
}

func (dao *_FileDao) GetByKey(vendor, key string) (*models.SysFile, bool, error) {
	item := &models.SysFile{
		Vendor: vendor,
		Key:    key,
	}
	has, err := dao.db().Get(item)
	if err != nil {
		return nil, false, err
	}
	return item, has, nil
}

func (dao *_FileDao) GetByIds(ids ...int64) ([]*models.SysFile, error) {
	rows := make([]*models.SysFile, 0)
	err := dao.db().In("`id`", ids).Find(&rows)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (dao *_FileDao) Insert(row *models.SysFile) (bool, error) {
	affected, err := dao.db().Insert(row)
	return affected > 0, err
}

func (dao *_FileDao) Update(id int64, row *models.SysFile) (bool, error) {
	affected, err := dao.db().Where("id=?", id).AllCols().AllCols().Update(row)
	return affected > 0, err
}

func (dao *_FileDao) Has(row *models.SysFile) (int64, error) {
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

func (dao *_FileDao) Delete(row *models.SysFile) (bool, error) {
	affected, err := dao.db().Delete(row)
	return affected > 0, err
}

func (dao *_FileDao) GetAll() ([]*models.SysFile, error) {
	rows := make([]*models.SysFile, 0)
	err := dao.db().Desc("`id`").Find(&rows)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (dao *_FileDao) GetByVendor(vendor string) ([]*models.SysFile, error) {
	rows := make([]*models.SysFile, 0)
	err := dao.db().Desc("`id`").Find(&rows, &models.SysFile{
		Vendor: vendor,
	})
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (dao *_FileDao) GetPage(name, vendor string, page, size int) ([]*models.SysFile, int64, error) {
	rows := make([]*models.SysFile, 0)
	session := dao.db().Asc("`id`")
	if name != "" {
		session.Where("`name` LIKE ?", "%"+name+"%")
	}
	if vendor != "" {
		session.And("`vendor`=?", vendor)
	}

	total, err := session.Limit(size, page*size-size).FindAndCount(&rows)
	if err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}
