package main

import (
	"os"

	"github.com/ccyun/daemon"
)

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "start":
			daemon.Start()
		case "restart":
			daemon.Restart()
		case "stop":
			daemon.Stop()
		default:
			daemon.Start()
		}
	} else {
		daemon.Start()
	}
}
