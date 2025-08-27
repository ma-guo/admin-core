package services

import (
	"github.com/ma-guo/admin-core/xorm/daos"
)

type _Svc struct {
	_dao   *daos.Dao
	prefix string
}

func NewSvc() *_Svc {
	return &_Svc{}
}

func (svc *_Svc) dao() *daos.Dao {
	if svc._dao == nil {
		svc._dao = daos.NewDao()
		svc.prefix = "svc"
	}
	return svc._dao
}

func (svc *_Svc) Close() {
	if svc._dao != nil {
		svc._dao.Close()
		svc._dao = nil
	}
}
