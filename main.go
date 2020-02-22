package main

import (
	"github.com/jianggushi/proxypool/cmd"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetLevel(logrus.InfoLevel)
	// logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.TextFormatter{
		// DisableColors: true,
		FullTimestamp: true,
	})
	cmd.Execute()
}
