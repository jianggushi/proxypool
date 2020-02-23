package common

import (
	"net/url"
	"time"

	"github.com/jianggushi/proxypool/pkg/model"
	"github.com/jianggushi/proxypool/pkg/request"
	"github.com/sirupsen/logrus"
)

func NewSpider(name string, urls []string, rule map[string]int) *Common {
	url, err := url.Parse(urls[0])
	if err != nil {
		return nil
	}
	return &Common{
		name:     name,
		urls:     urls,
		domain:   url.Host,
		duration: 1 * time.Minute,
		rule:     rule,
	}
}

var log = logrus.WithFields(logrus.Fields{})

type Common struct {
	name     string
	urls     []string
	domain   string
	duration time.Duration
	rule     map[string]int
}

func (s *Common) Name() string {
	return s.name
}

func (s *Common) Crawl(crawlChan chan<- *model.Proxy) {
	log := log.WithFields(logrus.Fields{
		"spider": s.name,
	})
	// iterate urls
	for _, url := range s.urls {
		// send http request
		resp, err := request.Get(url)
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
		proxies, err := extractProxy(resp.Body, s.rule)
		if err != nil {
			return
		}
		// send proxy to crawl channel
		for _, proxy := range proxies {
			proxy.From = s.domain
			proxy.Created = time.Now()
			crawlChan <- proxy
			log.WithFields(logrus.Fields{
				"proxy": proxy.Proxy,
			}).Info("crawl proxy")
		}
		// delay for next url
		time.Sleep(5 * time.Second)
	}
}
