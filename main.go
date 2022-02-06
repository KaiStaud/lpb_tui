package main

import (
	"fmt"
	"lpb/multilogger"
	"lpb/optest"
	"lpb/pipes"
	"lpb/storage"
	"lpb/tui"
	"time"
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
	storage.Init()
	multilogger.Init()
	tui.Launch()

	// Wait for all channels to be closed

	messages := make(chan string)
	go ping(messages)
	go pong(messages)
	time.Sleep(2 * time.Second)

	//messages <- "ping"
	pipes.Init()
}
