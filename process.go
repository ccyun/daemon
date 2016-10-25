package daemon

import (
	"log"
	"os"
)

var pid int

func run() int {
	//cmd := exec.Command("cmd", "dir")
	cmd, _ := os.FindProcess(8312)
	//log.Println(os.FindProcess(8312))
	// s, _ := (cmd.Output())
	// log.Println(string(s))
	//cmd.Process
	//cmd.Process.Kill()
	log.Println(cmd.Kill())

	//log.Println(cmd)
	return 0
}
