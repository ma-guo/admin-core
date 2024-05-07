package services

import (
	"fmt"

	"github.com/ma-guo/admin-core/xorm/models"
)

type _VendorSvc struct {
	*_Svc
}

func (svc *_Svc) Vendor() *_VendorSvc {
	return &_VendorSvc{svc}
}

func (svc *_VendorSvc) GetById(id int64) (*models.SysVendor, error) {
	if id <= 0 {
		return nil, fmt.Errorf("vendor不存在")
	}
	return svc.dao().Vendor().GetById(id)
}

func (svc *_VendorSvc) Update(id int64, row *models.SysVendor) error {
	_, err := svc.dao().Vendor().Update(id, row)
	return err
}

func (svc *_VendorSvc) Add(row *models.SysVendor) (*models.SysVendor, error) {
	item, has, err := svc.Has(row)
	if err != nil {
		return nil, err
	}
	if has {
		row.Id = item.Id
		_, err := svc.dao().Vendor().Update(row.Id, row)
		if err != nil {
			return nil, err
		}
		return row, nil
	}
	_, err = svc.dao().Vendor().Insert(row)
	if err != nil {
		return nil, err
	}
	return row, nil
}

// 插入一条数据
func (svc *_VendorSvc) Insert(row *models.SysVendor) error {
	_, err := svc.dao().Vendor().Insert(row)
	return err
}

func (svc *_VendorSvc) Has(row *models.SysVendor) (*models.SysVendor, bool, error) {
	return svc.dao().Vendor().GetByKey(row.Vendor, row.Key)
}

func (svc *_VendorSvc) GetAll() ([]*models.SysVendor, error) {
	return svc.dao().Vendor().GetAll()
}

func (svc *_VendorSvc) GetByKey(vendor, key string) (*models.SysVendor, bool, error) {
	return svc.dao().Vendor().GetByKey(vendor, key)
}

func (svc *_VendorSvc) GetByIds(ids ...int64) (map[int64]*models.SysVendor, error) {
	data := make(map[int64]*models.SysVendor, 0)
	rows, err := svc.dao().Vendor().GetByIds(ids...)
	if err != nil {
		return nil, err
	}
	for _, row := range rows {
		data[row.Id] = row
	}
	return data, nil
}

func (svc *_VendorSvc) GetPage(vendor string, page, size int) ([]*models.SysVendor, int64, error) {
	return svc.dao().Vendor().GetPage(vendor, page, size)
}

// 返回 name 为 key 的 map [key]value
func (svc *_VendorSvc) GetByVendorToMap(vendor string) (map[string]string, error) {
	rows, _, err := svc.dao().Vendor().GetPage(vendor, 1, 100)
	if err != nil {
		return nil, err
	}
	data := make(map[string]string)
	for _, row := range rows {
		data[row.Key] = row.Value
	}
	return data, nil
}
