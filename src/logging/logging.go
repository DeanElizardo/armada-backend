package logging

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func init() {
	Log = logrus.New()
	Log.SetLevel(logrus.TraceLevel)
	Log.SetReportCaller(true)

	logfile := "/home/del/Code/armada/armada-backend/syslog.txt"
	f, err := os.OpenFile(logfile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Failed to create logfile", logfile)
		panic(err)
	}

	Log.SetOutput(f)
}
