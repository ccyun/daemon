package main

import (
	"log"
	"time"
)

func main() {

	app := new(app)
	app.init = func() {

	}
	app.doFunc = func() {
		for i := 0; i < 60; i++ {
			log.Println(i)
			time.Sleep(1 * time.Second)
		}
	}
	app.run()
}
