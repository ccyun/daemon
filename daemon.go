package daemon

import (
	"log"
	"os/exec"
	"path/filepath"
)

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
	path, err := filepath.Abs(appConf.String("scriptFile"))
	if err != nil {
		return err
	}
	cmd := exec.Command(path)
	cmd.Start()
	log.Println(cmd.Process.Pid)
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
