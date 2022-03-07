package main

import (
	framehandling "lpb/frame_handling"
	"lpb/multilogger"
	"lpb/optest"
	"lpb/storage"
	"lpb/tracking"
	"time"

	"lpb/tui"
)

// TODO: Set machine into error mode, provide own module:
func ErrorHandler(chan string) {
	// Receive error message from tui or framehandling

	// Send error resolved after acknowledgemnt to tui:
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
	idle := make(chan bool)
	movement := make(chan bool)
	id_loopback := make(chan int)

	framehandling.Init(frames, id_loopback)

	// Create a channel for passing data from tui to tracking entity
	data := make(chan int, 3)
	tracking.Initialize(3, data)
	tui.StartJobQueue(data)

	p := tui.Launch(idle, movement)
	go tracking.StartTransmitReceiveQueue(p, data, id_loopback)
	framehandling.StartSimSwitch(p, idle, movement)

	//TODO: Use Waitgroup!
	for {
		time.Sleep(time.Second)
	}
}
