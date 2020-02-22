package spider

import (
	"github.com/jianggushi/proxypool/pkg/model"
	"github.com/jianggushi/proxypool/pkg/spider/freeip"
	"github.com/jianggushi/proxypool/pkg/spider/ip3366"
	"github.com/jianggushi/proxypool/pkg/spider/kuaidaili"
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
	Register(kuaidaili.SpiderKuaiDaiLi)
	Register(ip3366.SpiderIP3366)
	Register(freeip.SpiderFreeIP)
}
