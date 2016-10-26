package main

import (
	"log"
	"time"
)

func main() {
	for true {
		log.Println(time.Now())
		time.Sleep(1 * time.Second)
	}
}
