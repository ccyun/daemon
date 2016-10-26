package daemon

import (
	"path/filepath"

	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/logs"
)

var (
	appConf  config.Configer
	appCache cache.Cache
)

//InitConf 初始化配置
func InitConf() error {
	configPath, err := filepath.Abs("conf.ini")
	if err != nil {
		return err
	}
	ac, err := config.NewConfig("ini", configPath)
	if err != nil {
		return err
	}
	appConf = ac
	return nil
}

//InitCache 初始化缓存
func InitCache() error {
	bm, err := cache.NewCache("memory", `{"interval":60}`)
	if err != nil {
		return err
	}
	appCache = bm
	return nil
}

//InitLog 初始化log
func InitLog() error {
	err := logs.SetLogger("console")
	if err != nil {
		return err
	}
	return nil
}
