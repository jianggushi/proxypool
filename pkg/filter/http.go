package filter

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/jianggushi/proxypool/pkg/model"
)

func RequestBaidu(proxy *model.Proxy) (int, error) {
	baiduurl := "http://www.baidu.com/"
	return requestSite(proxy, baiduurl)
}

func requestSite(proxy *model.Proxy, siteurl string) (int, error) {
	// build proxy url for http client
	proxyurl := func(*http.Request) (*url.URL, error) {
		rawurl := fmt.Sprintf("http://%s:%s", proxy.Host, proxy.Port)
		return url.Parse(rawurl)
	}
	// new http client, timeout 10s
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: proxyurl,
		},
		Timeout: 5 * time.Second,
	}
	req, err := http.NewRequest("GET", siteurl, nil)
	if err != nil {
		return 0, err
	}
	t1 := time.Now()
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	// check http status code is ok?
	if resp.StatusCode != http.StatusOK {
		return 0, errors.New(resp.Status)
	}
	t := int(time.Now().Sub(t1).Milliseconds())
	return t, nil
}
