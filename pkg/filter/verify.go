package filter

import (
	"time"

	"github.com/jianggushi/proxypool/pkg/model"
)

func VerifyProxy(proxy *model.Proxy) error {
	tr, err := RequestBaidu(proxy)
	if err != nil {
		return err
	}
	proxy.Transfer = tr
	proxy.Checked = time.Now()
	return nil
}
