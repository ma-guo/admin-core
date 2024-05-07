package services

import (
	"fmt"

	"github.com/ma-guo/admin-core/xorm/models"
)

type _DictTypeSvc struct {
	*_Svc
}

func (svc *_Svc) DictType() *_DictTypeSvc {
	return &_DictTypeSvc{svc}
}

func (svc *_DictTypeSvc) GetById(id int64) (*models.SysDictType, error) {
	if id <= 0 {
		return nil, fmt.Errorf("部门信息不存在")
	}
	return svc.dao().DictType().GetById(id)
}

func (svc *_DictTypeSvc) Update(id int64, row *models.SysDictType) error {
	_, err := svc.dao().DictType().Update(id, row)
	return err
}

func (svc *_DictTypeSvc) Add(row *models.SysDictType) (*models.SysDictType, error) {
	tmp, has, err := svc.dao().DictType().GetByCode(row.Code)
	if err != nil {
		return nil, err
	}
	if has {
		return tmp, nil
	}
	_, err = svc.dao().DictType().Insert(row)
	if err != nil {
		return nil, err
	}
	return row, nil
}

func (svc *_DictTypeSvc) Delete(did int64) error {
	_, err := svc.dao().DictType().Delete(&models.SysDictType{Id: did})
	return err
}
func (svc *_DictTypeSvc) GetByCode(code string) (*models.SysDictType, bool, error) {
	return svc.dao().DictType().GetByCode(code)
}

func (svc *_DictTypeSvc) GetAll() ([]*models.SysDictType, error) {
	return svc.dao().DictType().GetAll()
}

func (svc *_DictTypeSvc) GetByIds(ids ...int64) (map[int64]*models.SysDictType, error) {
	data := make(map[int64]*models.SysDictType, 0)
	rows, err := svc.dao().DictType().GetByIds(ids...)
	if err != nil {
		return nil, err
	}
	for _, row := range rows {
		data[row.Id] = row
	}
	return data, nil
}

func (svc *_DictTypeSvc) GetPage(keywords string, page, size int) ([]*models.SysDictType, int64, error) {
	return svc.dao().DictType().GetPage(keywords, page, size)
}
