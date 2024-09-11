# `niuhe` 教程
在稀土掘金开了[niuhe 插件](https://juejin.cn/column/7376620206338785314) 系列教程。如有问题, 亦可咨询 `1057981162` (QQ/微信)

# 线上 demo
[admindemo](http://admindemo.zuxing.net/)
- 用户名 admin
- 密码 123456

# 引入本项目

下面为一个例子
```go
package main

import (
	"os"

	"github.com/ma-guo/admin-core/boot"
	"github.com/ma-guo/niuhe"
)

func main() {
	if len(os.Args) < 2 {
		niuhe.LogInfo("usage: %s <config-path>", os.Args[0])
		return
	}
	path := os.Args[1]
	boot := boot.AdminBoot{}
    // path 传 conf/admincore.yaml 路径
	if err := boot.LoadConfig(path); err != nil {
		panic(err)
	}
	svr := niuhe.NewServer()
	boot.BeforeBoot(svr)
	boot.RegisterModules(svr)
	boot.Serve(svr)
}
```

前端搭配项目为 [vue3-element-admin](https://github.com/ma-guo/vue3-element-admin), 本项目相关 api idl 项目为 [admin-core-niuhe](https://github.com/ma-guo/admin-core-niuhe)