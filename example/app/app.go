package main

import (
	"os"
	"os/signal"
	"runtime"
	"strconv"
	"syscall"
	"time"
)

type app struct {
	thread int
	done   chan bool
	close  bool
	doFunc func()
}

func (app *app) run() {
	if len(os.Args) > 1 {
		app.thread, _ = strconv.Atoi(os.Args[1])
	}
	if app.thread < 1 { //使用CPU多核处理
		app.thread = runtime.NumCPU()
	}
	runtime.GOMAXPROCS(app.thread)
	app.done = make(chan bool, app.thread)
	go app.listenSignal()
	app.work()
}

//doWork 执行程序
func (app *app) work() {
	for i := 0; i < app.thread; i++ {
		go func(i int) {
			for {
				if app.close == true {
					app.done <- true
					break
				}
				app.doFunc()
				time.Sleep(2 * time.Second)
			}
		}(i)
	}
	for i := 0; i < app.thread; i++ {
		<-app.done
	}
}

//listenSignal 监听信号
func (app *app) listenSignal() {
	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGQUIT, syscall.SIGKILL, syscall.SIGINT, syscall.SIGTERM)
	<-sig
	app.close = true
}
