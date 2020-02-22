package common

import (
	"net/url"
	"testing"
	"time"

	"github.com/jianggushi/proxypool/pkg/model"
)

func Test_UrlParse(t *testing.T) {
	rawurl := "https://www.kuaidaili.com/free/inha/1/"
	url, err := url.Parse(rawurl)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(url.Host, url.Hostname())
}

func Test_Crawl(t *testing.T) {
	ch := make(chan *model.Proxy, 100)
	go func() {
		for proxy := range ch {
			t.Log(proxy)
		}
	}()
	rules := map[string]int{
		"host":      0,
		"port":      1,
		"anonymity": 2,
		"scheme":    3,
	}
	url := []string{"https://www.kuaidaili.com/free/inha/1/"}
	spider := NewSpider("kuaidaili", url, rules)
	spider.Crawl(ch)
	time.Sleep(5 * time.Second)
}
