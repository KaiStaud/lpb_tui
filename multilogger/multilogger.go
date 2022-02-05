package multilogger

import (
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
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
var errLog *log.Logger

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

// Spawn goroutines
func Init() {

	writeSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	var logger = zap.New(core)
	slogger := logger.Sugar()
	slogger.Info("Info() uses sprint")
	slogger.Infof("Infof() uses %s", "sprintf")
	slogger.Infow("Infow() allows tags", "name", "Legolas", "type", 1)

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
