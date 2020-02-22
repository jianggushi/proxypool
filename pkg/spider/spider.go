package spider

import (
	"github.com/jianggushi/proxypool/conf"
	"github.com/jianggushi/proxypool/pkg/model"
	"github.com/jianggushi/proxypool/pkg/spider/common"
	"github.com/sirupsen/logrus"
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
	for _, sp := range conf.Conf.Spiders {
		Register(common.NewSpider(sp.Name, sp.Url, sp.Rule))
		logrus.Infof("register spider: %s", sp.Name)
	}
}
