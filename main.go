package main

import (
	"fmt"
	"lpb/multilogger"
	"lpb/optest"
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

	tui_logs := make(chan string)

	optest.SetConfig("~/lpb", "config")
	storage.Init()
	multilogger.Init(tui_logs)
	tui.Launch(tui_logs)

	// Wait for all channels to be closed

	messages := make(chan string)
	go ping(messages)
	go pong(messages)
	time.Sleep(2 * time.Second)

	//messages <- "ping"
}
