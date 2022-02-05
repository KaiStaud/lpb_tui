package multilogger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

/*
Provides multiple loginstances.
Each logger is parameterized with an absolute path to its log-file.
Logs are written via channels, execution is done seperately in a goroutine.
*/

var Slogger *zap.SugaredLogger
var tuilogs <-chan string

func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./test.log",
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

// Initialize logging  and setup gorotine:
// Expects channel on which logs are received
func Init(logs <-chan string) {

	writeSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	var logger = zap.New(core)
	Slogger := logger.Sugar()
	Slogger.Info("Info() uses sprint")
	Slogger.Infof("Infof() uses %s", "sprintf")
	Slogger.Infow("Infow() allows tags", "name", "Legolas", "type", 1)

	go TuiLogger(logs)
}

// Goroutine for TUI
func TuiLogger(logs <-chan string) {
	msg := <-logs
	Slogger.Info(msg)

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
