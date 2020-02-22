package ip89

import (
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/jianggushi/proxypool/pkg/model"
	"github.com/jianggushi/proxypool/pkg/request"
	"github.com/jianggushi/proxypool/pkg/util"
	"github.com/sirupsen/logrus"
)

var log = logrus.WithFields(logrus.Fields{
	"spider": SpiderIP89.name,
})

var SpiderIP89 = &IP89{
	name:   "89ip",
	domain: "www.89ip.cn",
	url:    "http://www.89ip.cn/index.html",
}

type IP89 struct {
	name   string
	domain string
	url    string
}

func (s *IP89) Name() string {
	return s.name
}

func (s *IP89) Crawl(ch chan<- *model.Proxy) {
	resp, err := request.Get(s.url)
	if err != nil {
		log.Errorf("http get: %v", err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Errorf("status code: %d %s", resp.StatusCode, resp.Status)
		return
	}
	proxies, err := s.find(resp.Body)
	if err != nil {
		log.Errorf("find proxy: %v", err)
		return
	}
	for _, proxy := range proxies {
		ch <- proxy
		log.WithFields(logrus.Fields{
			"proxy": fmt.Sprintf("%s:%s", proxy.Host, proxy.Port),
		}).Info("crawl proxy")
	}
}

// Find parse the html find proxy
func (s *IP89) find(r io.Reader) ([]*model.Proxy, error) {
	// load html
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return nil, err
	}
	// find proxy
	records, err := util.ReadTable(doc.Find("table"), true)
	if err != nil {
		return nil, err
	}
	proxies := make([]*model.Proxy, 0)
	for _, r := range records[1:] {
		proxy := &model.Proxy{
			Host:      r[0],
			Port:      r[1],
			Anonymity: model.ParseAnonymity(r[2]),
			Scheme:    model.ParseScheme(r[3]),
			From:      s.domain,
			Created:   time.Now(),
		}
		proxies = append(proxies, proxy)
	}
	if len(proxies) == 0 {
		return nil, errors.New("not find proxy")
	}
	return proxies, nil
}
