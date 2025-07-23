package log

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Logger = logrus.New()

func Init() {
	// set formatter: text or JSON
	Logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		ForceColors:   true,
	})

	// set log level: info, debug, warn, error
	Logger.SetLevel(logrus.InfoLevel)

	// log to stdout (default); you can also set output to file:
	Logger.SetOutput(os.Stdout)
}
