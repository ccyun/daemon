package common

//Task 任务
type Task struct {
	ID int64
}

//Tasker 任务接口
type Tasker interface {
	getOne() Task
}

//适配器
var adapters = make(map[string]Tasker)

//Register 初测适配器
func Register(name string, adapter Tasker) {
	if adapter == nil {
		panic("task: Register adapter is nil")
	}
	if _, ok := adapters[name]; ok {
		panic("task: Register called twice for adapter " + name)
	}
	adapters[name] = adapter
}

//Run 运行
func Run() {
	//dao.task.getOne()
}
