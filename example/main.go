package main

import (
	"log"
	"os"

	"github.com/ccyun/daemon"
)

func main() {
	var err error
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "start":
			err = daemon.Start()
		case "restart":
			err = daemon.Restart()
		case "stop":
			err = daemon.Stop()
		default:
			err = daemon.Start()
		}
	} else {
		err = daemon.Start()
	}
	log.Println(err)
}
