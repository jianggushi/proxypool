package schedule

import (
	"sync"
	"time"

	"github.com/jianggushi/proxypool/pkg/db"
	"github.com/jianggushi/proxypool/pkg/model"

	"github.com/jianggushi/proxypool/pkg/filter"
	"github.com/sirupsen/logrus"
)

var (
	numVerifyCrawl = 10
)

// DaemonVerifyCrawl 后台任务，验证 spider 爬取的 proxy 有效性，并存入 db
func DaemonVerifyCrawl() {
	// 启动 10 个 goroutine，后台任务
	for i := 0; i < numVerifyCrawl; i++ {
		go VerifyCrawl(i)
	}
}

func VerifyCrawl(index int) {
	log := log.WithFields(logrus.Fields{
		"index": index,
	})
	for proxy := range crawlChan {
		log := log.WithFields(logrus.Fields{
			"proxy": proxy.Proxy,
		})
		err := filter.VerifyProxy(proxy)
		if err != nil {
			log.Warn("verify crawl proxy fail")
			continue
		}
		log.Info("verify crawl proxy pass")
		err = db.Put(proxy)
		if err != nil {
			log.Errorf("save crawl proxy failure: %v", err)
			continue
		}
		log.Debug("save crawl proxy success")
	}
}

var (
	dbChanCap      = 20
	verifyDuration = 1 * time.Minute
)

// ScheduleVerifyDB 定时任务，验证 db 中 proxy 的有效性
func ScheduleVerifyDB() {
	for {
		t1 := time.Now()
		VerifyDB()
		log.Infof("verify db proxy cost time: %v", time.Since(t1))
		log.Infof("schedule verify db proxy next: %v", time.Now().Add(verifyDuration))
		<-time.After(verifyDuration)
	}
}

// VerifyDB
func VerifyDB() {
	var wg sync.WaitGroup
	proxies, err := db.GetAll()
	if err != nil {
		log.Error(err)
	}
	dbChan := make(chan *model.Proxy, dbChanCap)
	// start 10 goroutine to verify proxy
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			verifyDB(index, dbChan)
		}(i)
	}
	// send proxy to dbChan
	for _, proxy := range proxies {
		dbChan <- proxy
	}
	// close dbChan notify verify goroutine exit
	close(dbChan)
	wg.Wait()
}

func verifyDB(index int, dbChan <-chan *model.Proxy) {
	log := log.WithFields(logrus.Fields{
		"index": index,
	})
	for proxy := range dbChan {
		log := log.WithFields(logrus.Fields{
			"proxy": proxy.Proxy,
		})
		err := filter.VerifyProxy(proxy)
		if err != nil {
			log.Warn("verify db proxy fail")
			db.Delete(proxy) // ignore error
			continue
		}
		log.Info("verify db proxy pass")
		err = db.Update(proxy)
		if err != nil {
			log.Errorf("update db proxy failure: %v", err)
			continue
		}
		log.Debug("update db proxy success")
	}
}
