package cmd

import (
	"fmt"
	"sync"

	"github.com/jianggushi/proxypool/conf"
	"github.com/jianggushi/proxypool/pkg/filter"
	"github.com/jianggushi/proxypool/pkg/model"
	"github.com/jianggushi/proxypool/pkg/spider/common"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	spiderName string
)

var crawlCmd = &cobra.Command{
	Use:   "crawl",
	Short: "run crawl",
	Run:   runCrawl,
}

func runCrawl(cmd *cobra.Command, args []string) {
	// find spider config
	spiderConf := conf.Spider{}
	for _, sp := range conf.Conf.Spiders {
		if sp.Name == spiderName {
			spiderConf = sp
		}
	}
	if spiderConf.Name == "" {
		logrus.Fatalf("not find spider config: %s", spiderName)
	}

	crawlChan := make(chan *model.Proxy, 100)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(index int) {
			log := logrus.WithFields(logrus.Fields{
				"index": index,
			})
			for proxy := range crawlChan {
				log := log.WithFields(logrus.Fields{
					"proxy": fmt.Sprintf("%s:%s", proxy.Host, proxy.Port),
				})
				err := filter.VerifyProxy(proxy)
				if err != nil {
					log.Warn("verify crawl proxy fail")
					continue
				}
				log.Info("verify crawl proxy pass")
			}
			wg.Done()
		}(i)
	}
	spider := common.NewSpider(spiderConf.Name, spiderConf.Url, spiderConf.Rule)
	spider.Crawl(crawlChan)
	close(crawlChan)
	wg.Wait()
}

func init() {
	crawlCmd.Flags().StringVar(&spiderName, "spider", "", "spider for test")
	rootCmd.AddCommand(crawlCmd)
}
