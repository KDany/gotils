package log

import (
	"fmt"
	"os"
	"time"
)

type Config struct {
	Announcer   string
	TimeFormat  string
	LogPrefix   string
	WarnPrefix  string
	ErrorPrefix string
	ExitOnError bool
}

type loggerImpl struct{}

var conf Config

func (loggerImpl) SetConfig(c Config) {
	conf = c
}

var Logger loggerImpl

func logMessage(logType string, message ...any) {
	if conf != (Config{}) {
		var prefix string

		switch logType {
		case "info":
			prefix = conf.LogPrefix
		case "warn":
			prefix = conf.WarnPrefix
		case "error":
			prefix = conf.ErrorPrefix
		default:
			prefix = ""
		}

		formattedLog := " " + prefix + " "

		fmt.Print(
			time.Now().Format(conf.TimeFormat),
			formattedLog,
			conf.Announcer,
			": ",
		)
	}
	fmt.Println(message...)
}

func Info(message ...any) {
	logMessage("info", message...)
}

func Warn(message ...any) {
	logMessage("warn", message...)
}

func Error(message ...any) {
	logMessage("error", message...)
	if conf.ExitOnError {
		os.Exit(1)
	}
}
