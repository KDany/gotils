package gotils

import (
	"github.com/KDany/gotils/log"
)

const Test = "This is a test constant for gotils package"

func main() {
	// Example usage of Logger interface
	log.Logger.SetConfig(log.Config{
		Announcer:   "MyApp",
		TimeFormat:  "2006-01-02 15:04:05",
		LogPrefix:   "[INFO]",
		WarnPrefix:  "[WARN]",
		ErrorPrefix: "[ERROR]",
		ExitOnError: true,
	})
	log.Info("This is an info message")
	log.Error("This is an error message")
}
