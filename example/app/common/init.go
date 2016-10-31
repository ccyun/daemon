package common

import (
	"errors"

	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	//mysql driver
	_ "github.com/go-sql-driver/mysql"
)

//Conf 配置
var Conf config.Configer

func init() {
	func(f ...func() error) {
		for _, function := range f {
			if err := function(); err != nil {
				panic(err)
			}
		}
	}(InitConfig, InitLog, InitMySQL)
}

//InitConfig 初始化配置
func InitConfig() error {
	conf, err := config.NewConfig("ini", "conf.ini")
	if err != nil {
		return err
	}
	Conf = conf
	return nil
}

//InitLog 初始化log
func InitLog() error {
	logs.SetLogger("file", `{"filename":"`+Conf.String("log_file")+`"}`)
	return nil
}

//InitMySQL 初始化数据库
func InitMySQL() error {
	dsn := Conf.String("mysql_dsn")
	pool, _ := Conf.Int("mysql_pool")
	if dsn == "" || pool <= 0 {
		return errors.New("InitMySQL error, Configuration error.[mysql_dsn,mysql_pool]")
	}
	if err := orm.RegisterDriver("mysql", orm.DRMySQL); err != nil {
		return err
	}
	//最大数据库连接//最大空闲连接
	if err := orm.RegisterDataBase("default", "mysql", dsn, pool, pool); err != nil {
		return err
	}
	return nil
}
