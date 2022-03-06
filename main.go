package main

import (
	"fmt"
	framehandling "lpb/frame_handling"
	"lpb/multilogger"
	"lpb/optest"
	"lpb/storage"
	"lpb/tracking"
	"time"

	"lpb/tui"

	"github.com/go-gl/mathgl/mgl64"
)

var Ping chan int

// TODO: Set machine into error mode, provide own module:
func ErrorHandler(chan string) {
	// Receive error message from tui or framehandling

	// Send error resolved after acknowledgemnt to tui:
}

// The pinger prints a ping and waits for a pong
func pinger(pinger <-chan mgl64.Vec3) {
	for {
		<-pinger
		multilogger.AddTuiLog("got message")
		time.Sleep(time.Second)
	}
}

// The ponger prints a pong and waits for a ping
func ponger(pinger chan<- int) {
	for {
		fmt.Println("pong")
		time.Sleep(time.Second)
		pinger <- 1
	}
}

func main() {

	// TODO: Load user specific data from yaml-config
	optest.SetConfig("~/lpb", "config")

	// TODO: move DB Handling in Goroutine
	err := storage.Init("profiles.db")
	if err != nil {
		panic("DB not initialised")
	}
	multilogger.Init()

	frames := make(chan framehandling.DataFrame)
	framehandling.Init(frames)

	// Create a channel for passing data from tui to tracking entity
	data := make(chan mgl64.Vec3, 3)
	tracking.Initialize(3, data)
	go tui.StartJobQueue(data)

	//p := tui.Launch()
	//go tracking.StartReceiveQueue(p, data)
	//Ping := make(chan mgl64.Vec3)

	go pinger(data)
	//go ponger(ping)
	for {
		// Block the main thread until an interrupt
		time.Sleep(time.Second)
	}
}
