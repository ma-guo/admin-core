package views

// Generated by niuhe.idl

import (
	"github.com/ma-guo/admin-core/app/v1/protos"

	"github.com/ziipin-server/niuhe"
)

type _Gen_Dept struct{}

// 获取部门列表
func (v *_Gen_Dept) List_GET(c *niuhe.Context, req *protos.V1DeptListReq, rsp *protos.V1DeptListRsp) error {
	return niuhe.NewCommError(-1, "Not Implemented")
}

// 添加部门
func (v *_Gen_Dept) Add_POST(c *niuhe.Context, req *protos.V1DeptAddReq, rsp *protos.NoneRsp) error {
	return niuhe.NewCommError(-1, "Not Implemented")
}

// 获取部门下拉列表
func (v *_Gen_Dept) Options_GET(c *niuhe.Context, req *protos.NoneReq, rsp *protos.V1DeptOptionsRsp) error {
	return niuhe.NewCommError(-1, "Not Implemented")
}

// 修改部门
func (v *_Gen_Dept) Update_POST(c *niuhe.Context, req *protos.V1DeptUpdateReq, rsp *protos.NoneRsp) error {
	return niuhe.NewCommError(-1, "Not Implemented")
}

// 获取部门表单数据
func (v *_Gen_Dept) Form_GET(c *niuhe.Context, req *protos.V1DeptFormReq, rsp *protos.V1DeptFormRsp) error {
	return niuhe.NewCommError(-1, "Not Implemented")
}

// 删除部门
func (v *_Gen_Dept) Delete_POST(c *niuhe.Context, req *protos.V1DeptDeleteReq, rsp *protos.NoneRsp) error {
	return niuhe.NewCommError(-1, "Not Implemented")
}
