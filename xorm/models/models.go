package models

import "time"

// deleted 删除后前端不可见, 仅可数据库操作恢复
// 属性定义参考 https://xorm.io/zh/docs/chapter-02/4.columns/
type (

	// 部门表
	SysDept struct {
		Id         int64     `xorm:"NOTNULL PK AUTOINCR INT(11) COMMENT('主键')"`              // 排序 ID
		Name       string    `xorm:"VARCHAR(64) NOTNULL DEFAULT('') COMMENT('部门名称')"`        // 部门名称
		ParentId   int64     `xorm:"INT NOT NULL DEFAULT(0) INDEX COMMENT('父节点id')"`         // 父节点id
		TreePath   string    `xorm:"VARCHAR(255) NOTNULL DEFAULT('') COMMENT('父节点id路径')"`    // 父节点id路径
		Sort       int       `xorm:"INT NULL DEFAULT(0) COMMENT('显示顺序')"`                    // 显示顺序
		Status     int       `xorm:"TINYINT NOTNULL DEFAULT(1) COMMENT('状态(1:正常;0:禁用)')"`    // 状态(1:正常;0:禁用)
		Deleted    int       `xorm:"TINYINT NULL DEFAULT(0) COMMENT('逻辑删除标识(1:已删除;0:未删除)')"` // 逻辑删除标识(1:已删除;0:未删除)
		CreateTime time.Time `xorm:"created COMMENT('创建时间)"`                                 // 创建时间
		UpdateTime time.Time `xorm:"updated COMMENT('更新时间')"`                                // 更新时间
		CreateBy   int64     `xorm:"INT NULL DEFAULT(0) COMMENT('父节点id')"`                   // 创建人ID
		UpdateBy   int64     `xorm:"INT NULL DEFAULT(0) COMMENT('父节点id')"`                   // 修改人ID
	}

	// 字典数据表
	SysDict struct {
		Id         int64     `xorm:"NOTNULL PK AUTOINCR INT(11) COMMENT('主键')"`           // 排序 ID
		TypeCode   string    `xorm:"VARCHAR(64) NULL DEFAULT('') COMMENT('字典类型编码')"`      // 字典类型编码
		Name       string    `xorm:"VARCHAR(50) NULL DEFAULT('') COMMENT('字典项名称')"`       // 字典项名称
		Value      string    `xorm:"VARCHAR(50) NULL DEFAULT('') COMMENT('字典项值')"`        // 字典项值
		Sort       int       `xorm:"INT NULL DEFAULT(0) COMMENT('显示顺序')"`                 // 显示顺序
		Status     int       `xorm:"TINYINT NOTNULL DEFAULT(1) COMMENT('状态(1:正常;0:禁用)')"` // 状态(1:正常;0:禁用)
		Defaulted  bool      `xorm:"TINYINT NULL DEFAULT(0) COMMENT('是否默认(1:是;0:否)')"`    // 是否默认(1:是;0:否)
		Remark     string    `xorm:"VARCHAR(255) NULL DEFAULT('') COMMENT('备注')"`         // 备注
		CreateTime time.Time `xorm:"created COMMENT('创建时间)"`                              // 创建时间
		UpdateTime time.Time `xorm:"updated COMMENT('更新时间')"`                             // 更新时间
	}

	// 字典类型表
	SysDictType struct {
		Id         int64     `xorm:"NOTNULL PK AUTOINCR INT(11) COMMENT('主键')"`           // 排序 ID
		Name       string    `xorm:"VARCHAR(50) NULL DEFAULT('') COMMENT('类型名称')"`        // 类型名称
		Code       string    `xorm:"VARCHAR(50) NULL DEFAULT('') INDEX COMMENT('类型编码')"`  // 类型编码
		Status     int       `xorm:"TINYINT NOTNULL DEFAULT(1) COMMENT('状态(1:正常;0:禁用)')"` // 状态(1:正常;0:禁用)
		Remark     string    `xorm:"VARCHAR(255) NULL DEFAULT('') COMMENT('备注')"`         // 备注
		CreateTime time.Time `xorm:"created COMMENT('创建时间)"`                              // 创建时间
		UpdateTime time.Time `xorm:"updated COMMENT('更新时间')"`                             // 更新时间
	}

	// 菜单管理
	SysMenu struct {
		Id         int64     `xorm:"NOTNULL PK AUTOINCR INT(11) COMMENT('主键')"`                         // 排序 ID
		ParentId   int64     `xorm:"INT NOT NULL DEFAULT(0) COMMENT('父菜单ID')"`                          // 父菜单ID
		TreePath   string    `xorm:"VARCHAR(255) NULL COMMENT('父节点ID路径')"`                              // 父节点ID路径
		Name       string    `xorm:"VARCHAR(64) NOTNULL COMMENT('菜单名称'')"`                              // 菜单名称'
		Type       int       `xorm:"INT NULL DEFAULT(1) COMMENT('菜单类型(1:菜单 2:目录 3:外链 4:按钮)')"`          // 菜单类型(1:菜单 2:目录 3:外链 4:按钮)
		Path       string    `xorm:"VARCHAR(128) NULL DEFAULT('') COMMENT('路由路径(浏览器地址栏路径)')"`           // 路由路径(浏览器地址栏路径)
		Component  string    `xorm:"VARCHAR(128) NULL DEFAULT('') COMMENT('组件路径(vue页面完整路径，省略.vue后缀)')"` // 组件路径(vue页面完整路径，省略.vue后缀)
		Perm       string    `xorm:"VARCHAR(128) NULL DEFAULT('') COMMENT('权限标识')"`                     // 权限标识
		Visible    int       `xorm:"TINYINT NOTNULL DEFAULT(1) COMMENT('显示状态(1-显示;0-隐藏)')"`             // 显示状态(1-显示;0-隐藏)
		Sort       int       `xorm:"INT NULL DEFAULT(0) COMMENT('排序')"`                                 // 排序
		Icon       string    `xorm:"VARCHAR(64) NULL DEFAULT('') COMMENT('菜单图标')"`                      // 菜单图标
		Redirect   string    `xorm:"VARCHAR(64) NULL DEFAULT('') COMMENT('跳转路径')"`                      // 跳转路径
		CreateTime time.Time `xorm:"created COMMENT('创建时间)"`                                            // 创建时间
		UpdateTime time.Time `xorm:"updated COMMENT('更新时间')"`                                           // 更新时间
		AlwaysShow int       `xorm:"TINYINT DEFAULT(0) COMMENT('【目录】只有一个子路由是否始终显示(1:是 0:否)')"`          // 【目录】只有一个子路由是否始终显示(1:是 0:否)
		KeepAlive  int       `xorm:"TINYINT DEFAULT(0) COMMENT('菜单】是否开启页面缓存(1:是 0:否)')"`                // 菜单】是否开启页面缓存(1:是 0:否)
	}

	// 角色表
	SysRole struct {
		Id         int64     `xorm:"NOTNULL PK AUTOINCR INT(11) COMMENT('主键')"`                                 // 排序 ID
		Name       string    `xorm:"VARCHAR(50) NULL DEFAULT('') INDEX COMMENT('角色名称')"`                        // 角色名称
		Code       string    `xorm:"VARCHAR(50) NULL DEFAULT('') COMMENT('角色编码')"`                              // 角色编码
		Sort       int       `xorm:"INT NULL DEFAULT(0) COMMENT('显示顺序')"`                                       // 显示顺序
		Status     int       `xorm:"TINYINT NOTNULL DEFAULT(1) COMMENT('角色状态(1-正常；0-停用)')"`                     // 角色状态(1-正常；0-停用)
		DataScope  int       `xorm:"TINYINT NULL DEFAULT(0) COMMENT('数据权限(0-所有数据；1-部门及子部门数据；2-本部门数据；3-本人数据)')"` // 数据权限(0-所有数据；1-部门及子部门数据；2-本部门数据；3-本人数据)
		Deleted    bool      `xorm:"TINYINT NULL DEFAULT(0) COMMENT('逻辑删除标识(1:已删除;0:未删除)')"`                    // 逻辑删除标识(1:已删除;0:未删除)
		CreateTime time.Time `xorm:"created COMMENT('创建时间)"`                                                    // 创建时间
		UpdateTime time.Time `xorm:"updated COMMENT('更新时间')"`                                                   // 更新时间
	}

	// 角色和菜单关联表
	SysRoleMenu struct {
		RoleId int64 `xorm:"NOTNULL INT(11) COMMENT('角色ID')"` // 角色ID
		MenuId int64 `xorm:"NOTNULL INT(11) COMMENT('菜单ID')"` // 菜单ID
	}
	// 用户信息表
	SysUser struct {
		Id         int64     `xorm:"NOTNULL PK AUTOINCR INT(11) COMMENT('主键')"`              // 排序 ID
		Username   string    `xorm:"VARCHAR(64) NULL DEFAULT('') INDEX COMMENT('用户名')"`      // 用户名
		Nickname   string    `xorm:"VARCHAR(64) NULL DEFAULT('') INDEX COMMENT('昵称')"`       // 昵称
		Gender     int       `xorm:"TINYINT NULL DEFAULT(1) COMMENT('性别(1:男;2:女)')"`         // 性别(1:男;2:女)
		Password   string    `xorm:"VARCHAR(100) NULL DEFAULT('') COMMENT('密码')"`            // 密码
		DeptId     int64     `xorm:"NULL DEFAULT(0) INDEX COMMENT('部门ID')"`                  // 部门ID
		Avatar     string    `xorm:"VARCHAR(255) NULL DEFAULT('') COMMENT('用户头像')"`          // 用户头像
		Mobile     string    `xorm:"VARCHAR(20) NULL DEFAULT('') INDEX COMMENT('联系方式')"`     // 联系方式
		Status     bool      `xorm:"TINYINT NOTNULL DEFAULT(1) COMMENT('角色状态(1-正常；0-停用)')"`  // 角色状态(1-正常；0-停用)
		Email      string    `xorm:"VARCHAR(128) NULL DEFAULT('') COMMENT('用户邮箱')"`          // 用户邮箱
		Deleted    bool      `xorm:"TINYINT NULL DEFAULT(0) COMMENT('逻辑删除标识(1:已删除;0:未删除)')"` // 逻辑删除标识(1:已删除;0:未删除)
		CreateTime time.Time `xorm:"created COMMENT('创建时间)"`                                 // 创建时间
		UpdateTime time.Time `xorm:"updated COMMENT('更新时间')"`
	}

	// 用户和角色关联表
	SysUserRole struct {
		UserId int64 `xorm:"NOTNULL INT(11) INDEX COMMENT('用户ID')"` // 用户ID
		RoleId int64 `xorm:"NOTNULL INT(11) COMMENT('角色ID')"`       // 角色ID
	}

	// 文件列表
	SysFile struct {
		Id         int64     `xorm:"NOTNULL PK AUTOINCR INT(11) COMMENT('主键')"`           // 排序 ID
		UserId     int64     `xorm:"NOTNULL INT(11) INDEX COMMENT('上传者ID')"`              // 上传者ID
		Vendor     string    `xorm:"VARCHAR(32) NULL DEFAULT('') INDEX COMMENT('提供商')"`   // 提供商
		Name       string    `xorm:"VARCHAR(255) NULL DEFAULT('') INDEX COMMENT('文件名')"`  // 文件名
		Key        string    `xorm:"VARCHAR(255) NULL DEFAULT('') INDEX COMMENT('文件路径')"` // oss key, 删除时使用
		Path       string    `xorm:"VARCHAR(255) NULL DEFAULT('') INDEX COMMENT('文件路径')"` // 文件路径
		CreateTime time.Time `xorm:"created COMMENT('创建时间)"`                              // 创建时间
		UpdateTime time.Time `xorm:"updated COMMENT('更新时间')"`
		DeletedAt  time.Time `xorm:"deleted COMMENT('删除时间')"` // 删除时间
	}
	// OSS 服务配置信息
	SysVendor struct {
		Id         int64     `xorm:"NOTNULL PK AUTOINCR INT(11) COMMENT('主键')"`            // 排序 ID
		Vendor     string    `xorm:"VARCHAR(64) NULL DEFAULT('') INDEX COMMENT('字典类型编码')"` // 提供商名字
		Name       string    `xorm:"VARCHAR(50) NULL DEFAULT('') COMMENT('字典项名称')"`        // 配置名称
		Key        string    `xorm:"VARCHAR(50) NULL DEFAULT('') INDEX COMMENT('配置key')"`  // 配置key
		Value      string    `xorm:"VARCHAR(255) NULL DEFAULT('') COMMENT('字典项值')"`        // 配置值
		Remark     string    `xorm:"VARCHAR(255) NULL DEFAULT('') COMMENT('备注')"`          // 备注
		CreateTime time.Time `xorm:"created COMMENT('创建时间)"`                               // 创建时间
		UpdateTime time.Time `xorm:"updated COMMENT('更新时间')"`                              // 更新时间
	}
)

func GetSyncModels() []interface{} {
	return []interface{}{
		new(SysDept),
		new(SysDict),
		new(SysDictType),
		new(SysMenu),
		new(SysRole),
		new(SysRoleMenu),
		new(SysUser),
		new(SysUserRole),
		new(SysUser),
		new(SysFile),
		new(SysVendor),
	}
}
