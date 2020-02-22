package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"

	"github.com/jianggushi/proxypool/pkg/schedule"
)

func main() {
	logrus.SetLevel(logrus.InfoLevel)
	// logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.TextFormatter{
		// DisableColors: true,
		FullTimestamp: true,
	})
	signalChan := make(chan os.Signal, 1)
	sig := []os.Signal{syscall.SIGINT, syscall.SIGTERM}
	signal.Notify(signalChan, sig...)

	go schedule.ScheduleCrawl()
	go schedule.DaemonVerifyCrawl()
	go schedule.ScheduleVerifyDB()

	<-signalChan
}
