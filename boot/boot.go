package boot

// Generated by niuhe.idl

import (
	apiViews "github.com/ma-guo/admin-core/app/v1/views"
	"github.com/ma-guo/admin-core/config"

	"github.com/ziipin-server/niuhe"
)

type boot struct{}

type AdminBoot struct {
	boot
	protocol niuhe.IApiProtocol
}

func (AdminBoot) LoadConfig(path string) error {
	return config.LoadConfig(path)
}

func (AdminBoot) InitConfig(conf config.AdminConfig) error {
	config.InitConfig(conf)
	return nil
}

func (AdminBoot) BeforeBoot(svr *niuhe.Server) {}

func (admin AdminBoot) RegisterModules(svr *niuhe.Server) {
	if admin.protocol != nil {
		apiViews.SetProtocol(admin.protocol)
	}
	svr.RegisterModule(apiViews.GetModule())
}

func (AdminBoot) Serve(svr *niuhe.Server) {
	svr.Serve(config.Config.ServerAddr)
}

// func main() {
// 	if len(os.Args) < 2 {
// 		niuhe.LogInfo("usage: %s <config-path>", os.Args[0])
// 		return
// 	}
// 	path := os.Args[1]
// 	boot := AdminBoot{}
// 	if err := boot.LoadConfig(path); err != nil {
// 		panic(err)
// 	}
// 	svr := niuhe.NewServer()
// 	boot.BeforeBoot(svr)
// 	boot.RegisterModules(svr)
// 	boot.Serve(svr)
// }