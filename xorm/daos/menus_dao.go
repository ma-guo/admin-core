package daos

import (
	"fmt"

	"github.com/ma-guo/admin-core/xorm/models"

	"github.com/ziipin-server/niuhe"
)

type _MenuDao struct {
	*Dao
}

func (dao *Dao) Menu() *_MenuDao {
	return &_MenuDao{Dao: dao}
}

func (dao *_MenuDao) GetBy(item *models.SysMenu) (*models.SysMenu, bool, error) {
	has, err := dao.db().Get(item)
	if err != nil {
		return nil, false, err
	}
	return item, has, nil
}
func (dao *_MenuDao) GetByIds(ids ...int64) ([]*models.SysMenu, error) {
	rows := make([]*models.SysMenu, 0)
	err := dao.db().In("`id`", ids).Asc("`sort`").Find(&rows)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (dao *_MenuDao) Insert(row *models.SysMenu) (bool, error) {
	affected, err := dao.db().Insert(row)
	return affected > 0, err
}

func (dao *_MenuDao) Update(id int64, row *models.SysMenu) (bool, error) {
	affected, err := dao.db().Where("`id`=?", id).AllCols().Update(row)
	return affected > 0, err
}

func (dao *_MenuDao) Delete(row *models.SysMenu) (bool, error) {
	affected, err := dao.db().Delete(row)
	return affected > 0, err
}

func (dao *_MenuDao) Has(row *models.SysMenu) (bool, error) {
	niuhe.LogInfo("Has row: %v", row)
	if row.Id > 0 {
		_, has, err := dao.GetBy(&models.SysMenu{Id: row.Id})
		if err != nil {
			return false, err
		}
		return has, nil
	}
	return false, fmt.Errorf("user id is invalid")
}

func (dao *_MenuDao) GetAll(keyword string) ([]*models.SysMenu, error) {
	rows := make([]*models.SysMenu, 0)
	session := dao.db()
	if keyword != "" {
		session.Where("`name` LIKE ?", "%"+keyword+"%")
	}
	err := session.Asc("`sort`").Find(&rows)
	// err := dao.db().Asc("`sort`").Find(&rows)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (dao *_MenuDao) GetPage(keyword string, deptId int64, page, size int) ([]*models.SysMenu, int64, error) {
	rows := make([]*models.SysMenu, 0)
	session := dao.db().Where("`user_name` LIKE ? OR `nickname` LIKE ?", "%"+keyword+"%", "%"+keyword+"%")
	if deptId > 0 {
		session.And("`dept_id`=?", deptId)
	}
	total, err := session.Asc("`sort`").Limit(size, page*size-size).FindAndCount(&rows)
	if err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}
