package services

import (
	"fmt"

	"github.com/ma-guo/admin-core/xorm/models"
)

type _DictSvc struct {
	*_Svc
}

func (svc *_Svc) Dict() *_DictSvc {
	return &_DictSvc{svc}
}

func (svc *_DictSvc) GetById(id int64) (*models.SysDict, error) {
	if id <= 0 {
		return nil, fmt.Errorf("字典不存在")
	}
	return svc.dao().Dict().GetById(id)
}

func (svc *_DictSvc) Update(id int64, row *models.SysDict) error {
	_, err := svc.dao().Dict().Update(id, row)
	return err
}

func (svc *_DictSvc) Add(row *models.SysDict) (*models.SysDict, error) {
	item, has, err := svc.Has(row)
	if err != nil {
		return nil, err
	}
	if has {
		row.Id = item.Id
		_, err := svc.dao().Dict().Update(row.Id, row)
		if err != nil {
			return nil, err
		}
		return row, nil
	}
	_, err = svc.dao().Dict().Insert(row)
	if err != nil {
		return nil, err
	}
	return row, nil
}

// 插入一条数据
func (svc *_DictSvc) Insert(row *models.SysDict) error {
	_, err := svc.dao().Dict().Insert(row)
	return err
}

func (svc *_DictSvc) Delete(did int64) error {
	_, err := svc.dao().Dict().Delete(&models.SysDict{Id: did})
	return err
}
func (svc *_DictSvc) Has(row *models.SysDict) (*models.SysDict, bool, error) {
	return svc.dao().Dict().GetByName(row.TypeCode, row.Name)
}

func (svc *_DictSvc) GetAll() ([]*models.SysDict, error) {
	return svc.dao().Dict().GetAll()
}

func (svc *_DictSvc) GetByCode(typeCode string) ([]*models.SysDict, error) {
	return svc.dao().Dict().GetByCode(typeCode)
}

func (svc *_DictSvc) GetByIds(ids ...int64) (map[int64]*models.SysDict, error) {
	data := make(map[int64]*models.SysDict, 0)
	rows, err := svc.dao().Dict().GetByIds(ids...)
	if err != nil {
		return nil, err
	}
	for _, row := range rows {
		data[row.Id] = row
	}
	return data, nil
}

func (svc *_DictSvc) GetPage(keywords, typeCode string, page, size int) ([]*models.SysDict, int64, error) {
	return svc.dao().Dict().GetPage(keywords, typeCode, page, size)
}

// 返回 name 为 key 的 map [name]value
func (svc *_DictSvc) GetByCodeToMap(typeCode string) (map[string]string, error) {
	rows, err := svc.dao().Dict().GetByCode(typeCode)
	if err != nil {
		return nil, err
	}
	data := make(map[string]string)
	for _, row := range rows {
		data[row.Name] = row.Value
	}
	return data, nil
}
