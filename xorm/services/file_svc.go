package services

import (
	"fmt"

	"github.com/ma-guo/admin-core/xorm/models"
)

type _FileSvc struct {
	*_Svc
}

func (svc *_Svc) File() *_FileSvc {
	return &_FileSvc{svc}
}

func (svc *_FileSvc) GetById(id int64) (*models.SysFile, error) {
	if id <= 0 {
		return nil, fmt.Errorf("文件不存在")
	}
	return svc.dao().File().GetById(id)
}

func (svc *_FileSvc) GetByPath(path string) (*models.SysFile, bool, error) {
	return svc.dao().File().GetByPath(path)
}
func (svc *_FileSvc) Update(id int64, row *models.SysFile) error {
	_, err := svc.dao().File().Update(id, row)
	return err
}

func (svc *_FileSvc) Add(row *models.SysFile) (*models.SysFile, error) {
	item, has, err := svc.Has(row)
	if err != nil {
		return nil, err
	}
	if has {
		row.Id = item.Id
		_, err := svc.dao().File().Update(row.Id, row)
		if err != nil {
			return nil, err
		}
		return row, nil
	}
	_, err = svc.dao().File().Insert(row)
	if err != nil {
		return nil, err
	}
	return row, nil
}

func (svc *_FileSvc) Delete(did int64) error {
	_, err := svc.dao().File().Delete(&models.SysFile{Id: did})
	return err
}
func (svc *_FileSvc) Has(row *models.SysFile) (*models.SysFile, bool, error) {
	return svc.dao().File().GetByKey(row.Vendor, row.Key)
}

func (svc *_FileSvc) GetAll() ([]*models.SysFile, error) {
	return svc.dao().File().GetAll()
}

func (svc *_FileSvc) GetByVendor(vendor string) ([]*models.SysFile, error) {
	return svc.dao().File().GetByVendor(vendor)
}

func (svc *_FileSvc) GetByIds(ids ...int64) (map[int64]*models.SysFile, error) {
	data := make(map[int64]*models.SysFile, 0)
	rows, err := svc.dao().File().GetByIds(ids...)
	if err != nil {
		return nil, err
	}
	for _, row := range rows {
		data[row.Id] = row
	}
	return data, nil
}

func (svc *_FileSvc) GetPage(name, vendor string, page, size int) ([]*models.SysFile, int64, error) {
	return svc.dao().File().GetPage(name, vendor, page, size)
}
