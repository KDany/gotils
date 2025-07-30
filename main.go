package main

import (
	"gotils/logger"
)

func main() {
	// Example usage of Logger interface
	logger.SetConfig(logger.Config{
		Announcer:   "MyApp",
		TimeFormat:  "2006-01-02 15:04:05",
		LogPrefix:   "[INFO]",
		WarnPrefix:  "[WARN]",
		ErrorPrefix: "[ERROR]",
		ExitOnError: true,
	})
	logger.Log.Info("This is an info message")
	logger.Log.Error("This is an error message")
}
