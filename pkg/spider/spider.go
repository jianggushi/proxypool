package spider

import (
	"github.com/jianggushi/proxypool/pkg/model"
	"github.com/jianggushi/proxypool/pkg/spider/common"
)

var Spiders []Spider

type Spider interface {
	Crawl(c chan<- *model.Proxy)
	Name() string
}

func Register(sp Spider) {
	Spiders = append(Spiders, sp)
}

func init() {
	// Register(kuaidaili.SpiderKuaiDaiLi)
	// Register(ip3366.SpiderIP3366)
	// Register(freeip.SpiderFreeIP)
	var name, url string

	rules1 := map[string]int{
		"host":      0,
		"port":      1,
		"anonymity": 2,
		"scheme":    3,
	}
	rules2 := map[string]int{
		"host": 0,
		"port": 1,
	}
	name = "kuaidaili"
	url = "https://www.kuaidaili.com/free/inha/1/" // 国内高匿代理
	Register(common.NewSpider(name, url, rules1))

	name = "kuaidaili-2"
	url = "https://www.kuaidaili.com/free/intr/1/" // 国内普通代理
	Register(common.NewSpider(name, url, rules1))

	name = "ip3366"
	url = "http://www.ip3366.net/"
	Register(common.NewSpider(name, url, rules1))

	name = "freeip"
	url = "https://www.freeip.top/?page=1"
	Register(common.NewSpider(name, url, rules1))

	name = "89ip"
	url = "http://www.89ip.cn/index.html"
	Register(common.NewSpider(name, url, rules2))
}
