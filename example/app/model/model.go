package model

import "github.com/astaxie/beego/orm"

var (
	DBType   string
	DBPrefix string
	DB       orm.QueryBuilder
	O        orm.Ormer
)

//RegisterModels 注册Model
func RegisterModels() {
	orm.Debug = true
	O = orm.NewOrm()
	DB, _ = orm.NewQueryBuilder(DBPrefix)

	orm.RegisterModelWithPrefix(DBPrefix, new(Task))
}
