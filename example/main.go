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
	if err != nil {
		log.Println("\033[31m " + err.Error() + "\033[0m")
	} else {
		log.Println("\033[32m Successfully!\033[0m")
	}
}
