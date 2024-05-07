package services

import (
	"fmt"

	"github.com/ma-guo/admin-core/xorm/models"
)

type _DeptSvc struct {
	*_Svc
}

func (svc *_Svc) Dept() *_DeptSvc {
	return &_DeptSvc{svc}
}

func (svc *_DeptSvc) GetById(id int64) (*models.SysDept, error) {
	if id <= 0 {
		return nil, fmt.Errorf("部门信息不存在")
	}
	return svc.dao().Dept().GetById(id)
}

func (svc *_DeptSvc) Update(id int64, row *models.SysDept) error {
	_, err := svc.dao().Dept().Update(id, row)
	return err
}

func (svc *_DeptSvc) Add(row *models.SysDept) (*models.SysDept, error) {
	id := svc.Has(row)
	if id > 0 {
		row.Id = id
		_, err := svc.dao().Dept().Update(row.Id, row)
		if err != nil {
			return nil, err
		}
		return row, nil
	}
	_, err := svc.dao().Dept().Insert(row)
	if err != nil {
		return nil, err
	}
	return row, nil
}

func (svc *_DeptSvc) Delete(did int64) error {
	_, err := svc.dao().Dept().Delete(&models.SysDept{Id: did})
	return err
}
func (svc *_DeptSvc) Has(row *models.SysDept) int64 {
	team, _ := svc.dao().Dept().GetByName(row.Name)
	if team != nil && team.Id > 0 {
		return team.Id
	}
	return 0
}

func (svc *_DeptSvc) GetAll() ([]*models.SysDept, error) {
	return svc.dao().Dept().GetAll()
}

func (svc *_DeptSvc) Search(name string, status int) ([]*models.SysDept, error) {
	return svc.dao().Dept().Search(name, status)
}

func (svc *_DeptSvc) GetByIds(ids ...int64) (map[int64]*models.SysDept, error) {
	data := make(map[int64]*models.SysDept, 0)
	rows, err := svc.dao().Dept().GetByIds(ids...)
	if err != nil {
		return nil, err
	}
	for _, row := range rows {
		data[row.Id] = row
	}
	return data, nil
}
