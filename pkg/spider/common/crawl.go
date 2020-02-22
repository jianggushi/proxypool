package common

import (
	"fmt"
	"math/rand"
	"net/url"
	"time"

	"github.com/jianggushi/proxypool/pkg/model"
	"github.com/jianggushi/proxypool/pkg/request"
	"github.com/sirupsen/logrus"
)

func NewSpider(name string, rawurl string, rules map[string]int) *Common {
	url, err := url.Parse(rawurl)
	if err != nil {
		return nil
	}
	return &Common{
		name:     name,
		url:      rawurl,
		domain:   url.Host,
		duration: 1 * time.Minute,
		rules:    rules,
	}
}

var log = logrus.WithFields(logrus.Fields{})

type Common struct {
	name     string
	url      string
	domain   string
	duration time.Duration
	rules    map[string]int
}

func (s *Common) Name() string {
	return s.name
}

func (s *Common) Crawl(crawlChan chan<- *model.Proxy) {
	log := log.WithFields(logrus.Fields{
		"spider": s.name,
	})
	// random delay ?
	// 一个网站可能要爬取多个页面，通过 random delay 进行速率限制
	randomDelay := time.Duration(rand.Int63n(int64(3 * time.Second)))
	time.Sleep(randomDelay)
	// send http request
	resp, err := request.Get(s.url)
	if err != nil {
		log.Errorf("crawl requset get: %v", err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Errorf("crawl response status: %s", resp.Status)
		return
	}
	// extract proxy
	proxies, err := extractProxy(resp.Body, s.rules)
	if err != nil {
		return
	}
	// send proxy to crawl channel
	for _, proxy := range proxies {
		proxy.From = s.domain
		proxy.Created = time.Now()
		crawlChan <- proxy
		log.WithFields(logrus.Fields{
			"proxy": fmt.Sprintf("%s:%s", proxy.Host, proxy.Port),
		}).Info("crawl proxy")
	}
}
