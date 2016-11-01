package model

import "log"

//Task 任务表结构
type Task struct {
	ID           uint   `orm:"column(id)"`
	SiteID       uint   `orm:"column(site_id)"`
	CustomerCode string `orm:"column(customer_code)"`
	TaskType     string `orm:"column(task_type)"`
	Action       string `orm:"column(action)"`
	Status       uint8  `orm:"column(status)"`
	TryCount     uint8  `orm:"column(try_count)"`
	SetTimer     uint64 `orm:"column(set_timer)"`
	ModifiedAt   uint64 `orm:"column(modified_at)"`
}

func (t *Task) GetOne(id uint) {
	task := new(Task)
	//var data []*Task
	var data *Task
	O.QueryTable(task).OrderBy("status", "id").Limit(0, 1).One(&data)
	log.Println(data)
}
