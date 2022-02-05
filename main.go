package main

import (
	"lpb/multilogger"
	"lpb/optest"
	"lpb/storage"
)

func main() {

	optest.SetConfig("~/lpb", "config")
	storage.Init()
	multilogger.Init()
	//tui.Launch()
}
