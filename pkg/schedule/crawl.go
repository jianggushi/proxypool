package schedule

import (
	"fmt"
	"sync"
	"time"

	"github.com/jianggushi/proxypool/pkg/model"
	"github.com/jianggushi/proxypool/pkg/spider"
	"github.com/sirupsen/logrus"
)

var (
	crawlChan     chan *model.Proxy
	crawlChanCap  = 20
	crawlDuration = 10 * time.Second

	numSpiders int

	log = logrus.WithFields(logrus.Fields{})
)

// ScheduleCrawl 定时任务，爬取 proxy
func ScheduleCrawl() {
	numSpiders = len(spider.Spiders)
	crawlChan = make(chan *model.Proxy, numSpiders*crawlChanCap)
	for {
		Crawl()
		log.Infof("schedule crawl proxy next: %v", time.Now().Add(crawlDuration))
		<-time.After(crawlDuration)
	}
}

// Crawl 爬取 proxy
func Crawl() {
	var wg sync.WaitGroup
	for index, sp := range spider.Spiders {
		wg.Add(1)
		go func(index int, sp spider.Spider) {
			defer wg.Done()
			log := log.WithFields(logrus.Fields{
				"spider": sp.Name(),
				"index":  fmt.Sprintf("%d/%d", index+1, numSpiders),
			})
			log.Info("schedule crawl proxy start")
			sp.Crawl(crawlChan)
			log.Info("schedule crawl proxy end")
		}(index, sp)
	}
	wg.Wait()
}
