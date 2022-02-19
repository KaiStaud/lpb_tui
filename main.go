package main

import (
	"fmt"
	framehandling "lpb/frame_handling"
	"lpb/multilogger"
	"lpb/optest"
	"lpb/pipes"
	"lpb/storage"
)

func ping(pings chan<- string) {
	pings <- "msg"
}

func pong(pings <-chan string) {
	for {
		select {
		case <-quit:
			return
		default:
			msg := <-pings
			fmt.Print(msg)
		}

}

func main() {

	optest.SetConfig("~/lpb", "config")
	// TODO: move DB Handling in Goroutine
	db, _ := storage.Init()
	multilogger.Init()

	frames := make(chan framehandling.DataFrame)
	framehandling.Init(db, frames)

	tui.Launch()

	// Wait for all channels to be closed

	// messages := make(chan string)
	// go ping(messages)
	// go pong(messages)
	// time.Sleep(2 * time.Second)

	//messages <- "ping"
	pipes.Init()
}
