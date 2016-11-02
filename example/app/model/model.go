package model

import "github.com/astaxie/beego/orm"

var (
	//Debug 调试模式
	Debug bool
	//DBType 数据库类型
	DBType string
	//DBPrefix 表前缀
	DBPrefix string
	//DB 高级查询
	DB orm.QueryBuilder
)

//RegisterModels 注册Model
func RegisterModels() {
	orm.Debug = Debug
	orm.RegisterModelWithPrefix(DBPrefix, new(Queue))
	DB, _ = orm.NewQueryBuilder(DBType)
}

//Model 基础模型
type Model struct {
}
