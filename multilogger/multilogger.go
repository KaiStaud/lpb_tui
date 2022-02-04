package multilogger

import (
	"go.uber.org/zap"
)

/*
Provides multiple loginstances.
Each logger is parameterized with an absolute path to its log-file.
Logs are written via channels, execution is done seperately in a goroutine.
*/
//Channels are exported -> Capital letters
var TuiChannel string
var CanChannel string
var DbChannel string

// Spawn goroutines
func Init() {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()
	sugar.Debugf("TEST")
	/* 	TuiChannel := make(chan string)
	   	CanChannel := make(chan string)
	   	DbChannel := make(chan string)
	*/
	go TuiLogger()
	go CanLogger()
	go DBLogger()
	go DataLogger()
}

// Goroutine for TUI
func TuiLogger() {

}

// Goroutine for CAN
func CanLogger() {

}

// Goroutine for DB
func DBLogger() {

}

// Goroutine for Datahandling
func DataLogger() {

}
