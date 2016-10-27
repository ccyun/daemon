package daemon

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
)

var (
	//PID 进程PID
	PID        int
	scriptFile string
	pidFile    string
)

func init() {
	//初始化配置
	if err := InitConf(); err != nil {
		panic(err)
	}
	if err := InitLog(); err != nil {
		panic(err)
	}
	if File, err := filepath.Abs(appConf.String("scriptFile")); err != nil {
		panic(err)
	} else {
		scriptFile = File
	}
	if File, err := filepath.Abs(appConf.String("pidFile")); err != nil {
		panic(err)
	} else {
		pidFile = File
	}
}

//Start 启动进程
func Start() error {
	getPID()
	log.Println(PID)
	if PID != 0 {
		log.Println("Run is Successed!")
		return nil
	}
	cmd := exec.Command(scriptFile)
	if err := cmd.Start(); err != nil {
		return err
	}
	PID = cmd.Process.Pid
	if err := setPID(); err != nil {
		return err
	}
	log.Println("Run Successed!")
	return nil
}

//Restart 重新启动进程
func Restart() error {
	Stop()
	if err := Start(); err != nil {
		return err
	}
	return nil
}

//Stop 停止进程
func Stop() error {
	defer os.Remove(pidFile)
	if err := getPID(); err != nil {
		return err
	}
	cmd, err := os.FindProcess(PID)
	if err != nil {
		return err
	}
	if err := cmd.Kill(); err != nil {
		return err
	}
	return nil
}

//setPid PID写入文件
func setPID() error {
	dstFile, err := os.Create(pidFile)
	defer dstFile.Close()
	if err != nil {
		return err
	}
	dstFile.WriteString(strconv.Itoa(PID))
	return nil
}

//getPID 读取PID
func getPID() error {
	dstFile, err := os.Open(pidFile)
	defer dstFile.Close()
	if err != nil {
		return err
	}
	b, err := ioutil.ReadAll(dstFile)
	if err != nil {
		return err
	}
	pid, err := strconv.Atoi(string(b))
	if err != nil {
		return err
	}
	PID = pid
	return nil
}
