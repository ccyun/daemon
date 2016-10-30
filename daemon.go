package daemon

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"syscall"
	"time"

	"github.com/astaxie/beego/config"
)

type app struct {
	PID         int
	ScriptFile  string
	ScriptArgs  []string
	PidFile     string
	StopTimeOut int
}

var (
	appConf config.Configer
	//App 启动进程
	App app
)

func init() {
	var (
		configPath string
		err        error
		checkErr   func(error)
	)
	checkErr = func(err error) {
		if err != nil {
			panic(err)
		}
	}
	rootPath := filepath.Dir(os.Args[0]) + "/"
	configPath, err = filepath.Abs(rootPath + "./conf.ini")
	checkErr(err)
	appConf, err = config.NewConfig("ini", configPath)
	checkErr(err)
	App.ScriptFile = appConf.String("script_file")
	App.PidFile = appConf.String("pid_file")
	App.StopTimeOut, err = appConf.Int("stop_timeOut")
	checkErr(err)
	App.ScriptArgs = appConf.Strings("script_args")
}

//Run 运行指令
func Run() {
	var err error
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "start":
			err = App.start()
		case "restart":
			err = App.restart()
		case "stop":
			err = App.stop()
		default:
			err = errors.New("Invalid instruction.")
		}
		if err != nil {
			log.Println("\033[31m " + err.Error() + "\033[0m")
		} else {
			log.Println("\033[32m Successfully!\033[0m")
		}
	} else {
		fmt.Println(`This is a tool to manage the background script.
The commands are:
-------------------------------------------
start         Start the script process                                                                                                 
restart       Restart the script process                                                                                                      
stop          Stop the script process`)
	}
}

//Start 启动进程
func (App *app) start() error {
	App.getPID()
	if App.PID != -1 {
		return errors.New("Script is running!")
	}
	cmd := exec.Command(App.ScriptFile, App.ScriptArgs[0:]...)
	if err := cmd.Start(); err != nil {
		return err
	}
	App.PID = cmd.Process.Pid
	if err := App.setPID(); err != nil {
		return err
	}
	return nil
}

//Restart 重新启动进程
func (App *app) restart() error {
	if err := App.stop(); err != nil {
		return err
	}
	return App.start()
}

//Stop 停止进程
func (App *app) stop() error {
	var (
		err error
		cmd *os.Process
	)
	if err := App.getPID(); err != nil {
		return err
	}
	defer os.Remove(App.PidFile)

	if cmd, err = os.FindProcess(App.PID); err != nil {
		return err
	}
	if err = cmd.Signal(syscall.SIGQUIT); err != nil {
		return err
	}
	if err = App.exitedProcess(); err != nil {
		return err
	}
	return nil
}

//setPid PID写入文件
func (App *app) setPID() error {
	dstFile, err := os.Create(App.PidFile)
	if err != nil {
		return err
	}
	defer dstFile.Close()
	dstFile.WriteString(strconv.Itoa(App.PID))
	return nil
}

//getPID 读取PID
func (App *app) getPID() error {
	App.PID = -1
	dstFile, err := os.Open(App.PidFile)
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
	App.PID = pid
	return nil
}

//exitedProcess 判断进程是否退出
func (App *app) exitedProcess() error {
	ProDir := "/proc/" + strconv.Itoa(App.PID) + "/"
	log.Println("\033[33m Please waiting ...\033[0m")
	for i := 0; i < App.StopTimeOut; i++ {
		if i%5 == 0 {
			log.Println("\033[33m ...\033[0m")
		}
		time.Sleep(1 * time.Second)
		if _, err := os.Stat(ProDir); err != nil {
			App.PID = -1
			return nil
		}
	}
	return errors.New("Stop the script process timeOut")
}
