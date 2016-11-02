package model

import (
	"time"

	"github.com/astaxie/beego/orm"
)

//Queue 任务表结构
type Queue struct {
	Model
	ID           uint64 `orm:"column(id)"`
	SiteID       uint64 `orm:"column(site_id)"`
	CustomerCode string `orm:"column(customer_code)"`
	TaskType     string `orm:"column(task_type)"`
	Action       string `orm:"column(action)"`
	Status       uint8  `orm:"column(status)"`
	TryCount     uint8  `orm:"column(try_count)"`
	SetTimer     uint64 `orm:"column(set_timer)"`
	ModifiedAt   uint64 `orm:"column(modified_at),index"`
}

//TableName 表名
func (q *Queue) TableName() string {
	return "task"
}

//GetOneTask 读取单条数据
func (q *Queue) GetOneTask() (Queue, error) {
	taskInfo := Queue{}
	o := orm.NewOrm()
	cond := orm.NewCondition()
	nowTime := uint64(time.Now().UnixNano() / 1e6)
	condition := cond.And("SetTimer__lt", nowTime).AndCond(cond.And("Status", 0).OrCond(cond.And("Status", 3).And("TryCount__lt", 3)))
	err := o.QueryTable(q).SetCond(condition).OrderBy("Status", "ID").One(&taskInfo)
	if err != nil {
		return Queue{}, err
	}
	num, err := o.QueryTable(q).Filter("ID", taskInfo.ID).Filter("Status", 0).Filter("ModifiedAt", taskInfo.ModifiedAt).Update(orm.Params{"Status": 1, "ModifiedAt": nowTime, "TryCount": orm.ColValue(orm.ColAdd, 1)})
	if num == 0 {
		err = orm.ErrNoRows
	}
	if err != nil {
		return Queue{}, err
	}
	return taskInfo, nil
}

//Update 修改数据
func (q *Queue) Update(ID uint64) error {
	o := orm.NewOrm()
	num, err := o.Update(&Queue{ID: ID, Status: 3, ModifiedAt: uint64(time.Now().UnixNano() / 1e6)})
	if num == 0 {
		err = orm.ErrNoRows
	}
	return err
}

//Delete 删除数据
func (q *Queue) Delete(ID uint64) error {
	o := orm.NewOrm()
	if _, err := o.Delete(&Queue{ID: ID}); err != nil {
		return err
	}
	return nil
}
