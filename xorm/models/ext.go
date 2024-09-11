package models

import (
	"strings"

	"github.com/ma-guo/admin-core/app/common/consts"
	"github.com/ma-guo/admin-core/app/v1/protos"
)

func (row *SysRole) Enable() bool {
	return row.Status == consts.CommStatusEnum.Enable.Value
}

func (row *SysUser) IntStatus() int {
	if row.Status {
		return 1
	}
	return 0
}

func (row *SysUser) SetStatus(status int) {
	row.Status = status == 1
}

func (row *SysUser) GenderLabel() string {
	if label, has := consts.GenderEnum.GetChoices()[row.Gender]; has {
		return label
	}
	return ""
}

// 获取按钮类型对应的枚举值 NOLL, BUTTON, MENU, CATEGORY, EXTLINK
func (row *SysMenu) GetTypeValue() string {
	if label, has := consts.MenuTypeGroup.GetChoices()[row.Type]; has {
		return label
	}
	return consts.MenuTypeEnum.NULL.Value
}

func (rw *SysMenu) IsRoute() bool {
	return rw.Type == consts.MenuTypeGroup.MENU.Value || rw.Type == consts.MenuTypeGroup.EXTLINK.Value || rw.Type == consts.MenuTypeGroup.CATALOG.Value
}

func (row *SysMenu) SetType(_type string) {
	choices := consts.MenuTypeGroup.GetChoices()
	for key, value := range choices {
		if value == _type {
			if key == consts.MenuTypeGroup.CATALOG.Value {
				// 目录类型 path 需要检查是否以 / 开头
				if strings.Index(row.Path, "/") != 0 {
					row.Path = "/" + row.Path
				}
				// 目录类型 component 为 Layout
				row.Component = "Layout"
			}
			row.Type = key
			return
		}
	}
}

func (row *SysMenu) ToProtos() *protos.MenuItem {
	item := &protos.MenuItem{
		Children:  make([]*protos.MenuItem, 0),
		Component: row.Component,
		Icon:      row.Icon,
		Id:        row.Id,
		Name:      row.Name,
		ParentId:  row.ParentId,
		Path:      row.Path,
		Perm:      row.Perm,
		Redirect:  row.Redirect,
		Sort:      row.Sort,
		Type:      row.GetTypeValue(),
		Visible:   row.Visible,
	}
	return item
}

func (row *SysDept) ToProtos() *protos.V1DeptItem {
	item := &protos.V1DeptItem{
		Children:   make([]*protos.V1DeptItem, 0),
		Id:         row.Id,
		Name:       row.Name,
		ParentId:   row.ParentId,
		Sort:       row.Sort,
		Status:     row.Status,
		CreateTime: row.CreateTime.Format(consts.FullTimeLayout),
		UpdateTime: row.UpdateTime.Format(consts.FullTimeLayout),
	}
	return item
}

func (row *SysVendor) ToProps() *protos.V1VendorItem {
	return &protos.V1VendorItem{
		Id:         row.Id,
		Vendor:     row.Vendor,
		Name:       row.Name,
		Value:      row.Value,
		Remark:     row.Remark,
		Key:        row.Key,
		UpdateTime: row.UpdateTime.Format(consts.FullTimeLayout),
	}
}
