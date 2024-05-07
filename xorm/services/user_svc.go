package services

import (
	"fmt"
	"strings"

	"github.com/ma-guo/admin-core/app/v1/protos"
	"github.com/ma-guo/admin-core/xorm/models"
)

type _UserSvc struct {
	*_Svc
}

func (svc *_Svc) User() *_UserSvc {
	return &_UserSvc{svc}
}

func (svc *_UserSvc) GetById(id int64) (*models.SysUser, error) {
	if id <= 0 {
		return nil, fmt.Errorf("用户信息不存在")
	}
	row, has, err := svc.dao().User().GetBy(&models.SysUser{Id: id})
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, fmt.Errorf("用户信息不存在")
	}
	return row, nil
}

func (svc *_UserSvc) GetByUsername(username string) (*models.SysUser, bool, error) {
	if username == "" {
		return nil, false, fmt.Errorf("用户信息不存在")
	}
	return svc.dao().User().GetBy(&models.SysUser{Username: username})
}

func (svc *_UserSvc) GetByName(username, nickname string) (*models.SysUser, bool, error) {
	if username == "" || nickname == "" {
		return nil, false, fmt.Errorf("用户名或昵称不能为空")
	}
	return svc.dao().User().GetBy(&models.SysUser{Username: username, Nickname: nickname})
}

func (svc *_UserSvc) Update(id int64, row *models.SysUser) error {
	_, err := svc.dao().User().Update(id, row)
	return err
}

func (svc *_UserSvc) Insert(row *models.SysUser) (*models.SysUser, error) {
	_, err := svc.dao().User().Insert(row)
	if err != nil {
		return nil, err
	}
	return row, nil
}

func (svc *_UserSvc) Has(row *models.SysUser) bool {
	_, has, _ := svc.dao().User().GetBy(&models.SysUser{Username: row.Username, Nickname: row.Nickname})
	return has
}

func (svc *_UserSvc) GetAll() ([]*models.SysUser, error) {
	return svc.dao().User().GetAll()
}

func (svc *_UserSvc) GetPage(param *protos.V1UsersPageReq, query string) ([]*models.SysUser, int64, error) {
	if param.PageNum < 0 {
		param.PageNum = 1
	}
	if param.PageSize < 10 {
		param.PageSize = 10
	}
	param.Keywords = strings.Trim(param.Keywords, " ")
	param.StartTime = strings.Trim(param.StartTime, " ")
	param.EndTime = strings.Trim(param.EndTime, " ")
	hasStatus := false
	if strings.Contains(query, "status") {
		hasStatus = true
	}
	deptIds := make([]int64, 0)
	if param.DeptId > 0 {
		deptIds = append(deptIds, param.DeptId)
		children, err := svc.dao().Dept().GetChildren(param.DeptId)
		if err != nil {
			return nil, 0, err
		}
		for _, dept := range children {
			deptIds = append(deptIds, dept.Id)
		}
	}

	return svc.dao().User().GetPage(param, hasStatus, deptIds...)
}
