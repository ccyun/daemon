package main

import (
	"os"
	"os/signal"
	"syscall"
)

type app struct {
	done   chan bool
	close  bool
	doFunc func()
	init   func()
}

func (app *app) run() {
	app.done = make(chan bool, 1)
	go app.listenSignal()
	app.init()
	app.doWork()
	<-app.done
}

//doWork 执行程序
func (app *app) doWork() {
	for {
		app.doFunc()
		if app.close == true {
			break
		}
	}
	app.done <- true
}

//listenSignal 监听信号
func (app *app) listenSignal() {
	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGQUIT, syscall.SIGKILL, syscall.SIGINT, syscall.SIGTERM)
	<-sig
	app.close = true
}
