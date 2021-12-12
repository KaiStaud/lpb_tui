package main

import (
	"lpb/optest"
	"lpb/storage"
	tui "lpb/textinterface"
)

func main() {
	optest.SetConfig("~/lpb", "config")
	storage.Init()
	tui.Launch()
}
