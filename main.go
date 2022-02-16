package main

import (
	"fmt"
	framehandling "lpb/frame_handling"
	"lpb/multilogger"
	"lpb/optest"
	"lpb/pipes"
	"lpb/storage"
	"lpb/tui"
)

func ping(pings chan<- string) {
	pings <- "msg"
}

func pong(pings <-chan string) {
	msg := <-pings
	fmt.Print(msg)
}

func main() {

	optest.SetConfig("~/lpb", "config")
	// TODO: move DB Handling in Goroutine
	db, _ := storage.Init()
	multilogger.Init()
	framehandling.Init(db)

	// a0 := storage.Arm{1, 1, 1, 0}
	// a1 := storage.Arm{2, 2, 2, 0}

	// storage.AddArm(a0, db)
	// storage.AddArm(a1, db)
	// storage.AddArm(a1, db)

	tui.Launch()

	// Wait for all channels to be closed

	// messages := make(chan string)
	// go ping(messages)
	// go pong(messages)
	// time.Sleep(2 * time.Second)

	//messages <- "ping"
	pipes.Init()
}
