package multilogger

import (
	"fmt"
	"lpb/pipes"
	"strings"
	"time"

	"github.com/syrinsecurity/gologger"
)

const (
	Info     = iota
	Debug    = iota
	Warning  = iota
	Error    = iota
	Critical = iota
	Panic    = iota
)

/*
Provides multiple loginstances.
Each logger is parameterized with an absolute path to its log-file.
Logs are written via channels, execution is done seperately in a goroutine.
*/

// Initialize logging  and setup gorotine:
// Expects channel on which logs are received
func Init() {

	// Buffer messages on channel
	pipes.DebugMessages = make(chan string, 10)

	go TuiLogger(pipes.DebugMessages)
}

// Goroutine for TUI
func TuiLogger(logs <-chan string) {

	logger, err := gologger.New("./file.log", 200)
	if err != nil {
		panic(err)
	}

	for {
		msg := <-logs
		fields := strings.Split(msg, ":")
		logger.WriteString(fmt.Sprintf("<%d-%02d-%02d %02d:%02d:%02d> [%s] %s",
			time.Now().Year(), time.Now().Month(), time.Now().Day(),
			time.Now().Hour(), time.Now().Minute(), time.Now().Second(),
			fields[0], fields[1]))
	}

}

// Goroutine for CAN
func CanLogger(frames <-chan string) {

}

// Goroutine for DB
func DBLogger(logs <-chan string) {

}

// Goroutine for Datahandling
func DataLogger(logs <-chan string) {

}
