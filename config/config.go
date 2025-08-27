package config

import (
	"os"

	"github.com/ma-guo/admin-core/xorm/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ma-guo/niuhe"
	"gopkg.in/yaml.v2"
	"xorm.io/xorm"
	"xorm.io/xorm/log"
)

type AdminConfig struct {
	ServerAddr string // 端口或者 socket 文件
	Debug      bool   // 是否 debug 模式
	LogLevel   string // 日志级别
	DB         struct {
		ShowSQL bool   // 是否显示 sql 日志
		Sync    bool   // 是否同步数据库
		Debug   bool   // 是否开启 debug 模式
		Main    string // 数据库链接信息 dataSourceName
	}
	// Filedir    string // 文件存储目录
	Host       string // http + 域名
	Fileprefix string // OSS文件前缀
	Templates  string // 模板目录
	Secretkey  string // bearer secret key
	// TODO 在这下面添加其他配置内容
}

var Config AdminConfig
var MainDB *xorm.Engine

func LoadConfig(path string) error {
	configBytes, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	if err := yaml.Unmarshal(configBytes, &Config); err != nil {
		return err
	}
	niuhe.SetLogLevel(niuhe.MustParseLogLevelName(Config.LogLevel))

	if err := initDatabases(); err != nil {
		return err
	}
	err = initDatabaseData()
	if err != nil {
		niuhe.LogInfo("init database data error: %v", err)
	}
	return nil
}

// 初始化配置
func InitConfig(config AdminConfig) error {
	Config = config
	niuhe.SetLogLevel(niuhe.MustParseLogLevelName(Config.LogLevel))
	if err := initDatabases(); err != nil {
		niuhe.LogError("init databases error: %v", err)
		return err
	}
	return nil
}

func initDatabases() error {
	var err error
	// 初始化主库
	if MainDB, err = xorm.NewEngine("mysql", Config.DB.Main); err != nil {
		return err
	}
	if Config.DB.ShowSQL {
		niuhe.LogInfo("xorm: main db show sql")
		MainDB.ShowSQL(true)
	}
	if Config.DB.Debug {
		niuhe.LogInfo("xorm: main db log debug")
		MainDB.Logger().SetLevel(log.LOG_DEBUG)
	}
	if Config.DB.Sync {
		err := MainDB.Sync2(models.GetSyncModels()...)
		niuhe.LogInfo("xorm: main db sync, %v", err)
	}

	return nil
}

func init() {

}
