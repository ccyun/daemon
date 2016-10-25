package daemon

import "log"

func init() {
	//初始化配置
	if err := InitConf(); err != nil {
		panic(err)
	}
	if err := InitLog(); err != nil {
		panic(err)
	}
}

//Start 启动进程
func Start() error {
	log.Println(pid)
	pid = run()
	log.Println(pid)
	return nil
}

//Restart 重新启动进程
func Restart() error {
	return nil
}

//Stop 停止进程
func Stop() error {
	return nil
}
