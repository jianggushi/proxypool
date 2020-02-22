package filter

import (
	"testing"

	"github.com/jianggushi/proxypool/pkg/model"
)

func Test_RequestBaidu(t *testing.T) {
	proxy := &model.Proxy{
		IP:       "114.106.211.167",
		Port:     "8119",
		Protocol: model.Http,
	}
	tr, err := RequestBaidu(proxy)
	if err != nil {
		t.Error(err)
	}
	t.Logf("transfer time: %d", tr)
}
