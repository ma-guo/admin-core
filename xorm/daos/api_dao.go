package daos

import (
	"fmt"

	"github.com/ma-guo/admin-core/xorm/models"

	"github.com/ma-guo/niuhe"
)

type _ApiDao struct {
	*Dao
}

func (dao *Dao) Api() *_ApiDao {
	return &_ApiDao{Dao: dao}
}

func (dao *_ApiDao) GetById(id int64) (*models.SysApi, error) {
	item := &models.SysApi{
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

func (dao *_ApiDao) GetByMethod(method, path string) (*models.SysApi, bool, error) {
	item := &models.SysApi{
		Method: method,
		Path:   path,
	}
	has, err := dao.db().Get(item)
	if err != nil {
		return nil, false, err
	}
	return item, has, nil
}

func (dao *_ApiDao) GetByIds(ids ...int64) ([]*models.SysApi, error) {
	rows := make([]*models.SysApi, 0)
	err := dao.db().In("`id`", ids).Find(&rows)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (dao *_ApiDao) Insert(row *models.SysApi) (bool, error) {
	if row.MenuIds == nil {
		row.MenuIds = []int64{}
	}
	affected, err := dao.db().Insert(row)
	return affected > 0, err
}

func (dao *_ApiDao) Update(id int64, row *models.SysApi) (bool, error) {
	if row.MenuIds == nil {
		row.MenuIds = []int64{}
	}
	affected, err := dao.db().Where("id=?", id).AllCols().Update(row)
	return affected > 0, err
}

func (dao *_ApiDao) Has(row *models.SysApi) (int64, error) {
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

func (dao *_ApiDao) GetAll() ([]*models.SysApi, error) {
	rows := make([]*models.SysApi, 0)
	err := dao.db().Asc("`id`").Find(&rows)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (dao *_ApiDao) GetPage(keyword string, page, size int) ([]*models.SysApi, int64, error) {
	rows := make([]*models.SysApi, 0)
	session := dao.db().Where("`name` LIKE ?", "%"+keyword+"%")
	total, err := session.Limit(size, page*size-size).FindAndCount(&rows)
	if err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}
