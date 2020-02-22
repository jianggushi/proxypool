package cmd

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/jianggushi/proxypool/pkg/schedule"
	"github.com/spf13/cobra"
)

var crawlCmd = &cobra.Command{
	Use:   "crawl",
	Short: "run crawl",
	Run:   runCrawl,
}

func runCrawl(cmd *cobra.Command, args []string) {
	signalChan := make(chan os.Signal, 1)
	sig := []os.Signal{syscall.SIGINT, syscall.SIGTERM}
	signal.Notify(signalChan, sig...)

	go schedule.ScheduleCrawl()
	go schedule.DaemonVerifyCrawl()
	go schedule.ScheduleVerifyDB()

	<-signalChan
}

func init() {
	rootCmd.AddCommand(crawlCmd)
}
