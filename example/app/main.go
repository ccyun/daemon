package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"
)

var i = 1

func main() {
	go func() {
		for {
			if i == 0 {
				break
			}

			log.Println("11111")
		}
	}()
	c := make(chan os.Signal)
	signal.Notify(c)
	//监听指定信号
	//signal.Notify(c, syscall.SIGHUP, syscall.SIGUSR2)

	//阻塞直至有信号传入
	s := <-c

	fmt.Println("get signal:", s)
	logFile, _ := filepath.Abs("./log")
	dstFile, _ := os.Create(logFile)
	dstFile.WriteString("接收到的信号是：" + s.String())
	i = 0
}
