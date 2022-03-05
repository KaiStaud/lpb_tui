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

var (
	logger, _ = gologger.New("./file.log", 200)
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

	for {
		msg := <-logs
		fields := strings.Split(msg, ":")
		logger.WriteString(fmt.Sprintf("<%d-%02d-%02d %02d:%02d:%02d> [%s] %s",
			time.Now().Year(), time.Now().Month(), time.Now().Day(),
			time.Now().Hour(), time.Now().Minute(), time.Now().Second(),
			fields[0], fields[1]))
	}

}

// Write new log entry
func AddTuiLog(data string) {

	fields := strings.Split(data, ":")
	// No loglevel or empty string: tag w/ unknown
	if len(fields) <= 1 {
		fields[1] = "Unknown"
	}
	logger.WriteString(fmt.Sprintf("<%d-%02d-%02d %02d:%02d:%02d> [%s] %s",
		time.Now().Year(), time.Now().Month(), time.Now().Day(),
		time.Now().Hour(), time.Now().Minute(), time.Now().Second(),
		fields[0], strings.ToUpper(fields[1])))

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
