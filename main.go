package main

import (
	"lpb/storage"
	tui "lpb/textinterface"
)

func main() {
	/*
		config, err := cmd.LoadConfig()

		if err != nil {
			log.Fatal("error while looading config")
		}
			fmt.Println("Struct:", config)
	*/
	storage.Init()
	tui.Launch()
}
