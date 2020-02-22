package cmd

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/jianggushi/proxypool/pkg/schedule"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var scheduleCmd = &cobra.Command{
	Use:   "schedule",
	Short: "run schedule",
	Run:   runSchedule,
}

func runSchedule(cmd *cobra.Command, args []string) {
	signalChan := make(chan os.Signal, 1)
	sig := []os.Signal{syscall.SIGINT, syscall.SIGTERM}
	signal.Notify(signalChan, sig...)

	go schedule.ScheduleCrawl()
	go schedule.DaemonVerifyCrawl()
	go schedule.ScheduleVerifyDB()

	<-signalChan
}

func init() {
	logrus.SetLevel(logrus.InfoLevel)
	// logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.TextFormatter{
		// DisableColors: true,
		FullTimestamp: true,
	})
	// file, err := os.OpenFile("log/crawl.log", os.O_CREATE|os.O_WRONLY, 0666)
	// if err != nil {
	// 	logrus.Fatalf("create crawl.log: %v", err)
	// }
	// logrus.SetOutput(file)
	rootCmd.AddCommand(scheduleCmd)
}
