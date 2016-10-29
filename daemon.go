package daemon

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"syscall"
	"time"
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
	if PID != -1 {
		return errors.New("Script is running!")
	}
	cmd := exec.Command(scriptFile)
	if err := cmd.Start(); err != nil {
		return err
	}
	PID = cmd.Process.Pid
	if err := setPID(); err != nil {
		return err
	}
	return nil
}

//Restart 重新启动进程
func Restart() error {
	if err := Stop(); err != nil {
		return err
	}
	return Start()
}

//Stop 停止进程
func Stop() error {
	var (
		err error
		cmd *os.Process
	)
	if err = getPID(); err != nil {
		return err
	}
	defer os.Remove(pidFile)

	if cmd, err = os.FindProcess(PID); err != nil {
		return err
	}
	if err = cmd.Signal(syscall.SIGQUIT); err != nil {
		return err
	}
	if err = exitedProcess(); err != nil {
		return err
	}
	return nil
}

//setPid PID写入文件
func setPID() error {
	dstFile, err := os.Create(pidFile)
	if err != nil {
		return err
	}
	defer dstFile.Close()
	dstFile.WriteString(strconv.Itoa(PID))
	return nil
}

//getPID 读取PID
func getPID() error {
	PID = -1
	dstFile, err := os.Open(pidFile)
	if err != nil {
		return err
	}
	defer dstFile.Close()
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

//exitedProcess 判断进程是否退出
func exitedProcess() error {
	ProDir := "/proc/" + strconv.Itoa(PID) + "/"
	s := "\033[33m Please waiting "

	for {
		s += "."
		log.Println(s + "\033[0m")
		time.Sleep(5 * time.Second)
		if _, err := os.Stat(ProDir); err != nil {
			break
		}
	}
	return nil
}
