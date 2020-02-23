package request

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

var (
	client = &http.Client{
		Timeout: 5 * time.Second,
	}
)

func Get(rawurl string) (resp *http.Response, err error) {
	req, err := http.NewRequest("GET", rawurl, nil)
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(rawurl)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Host", u.Host)
	req.Header.Set("Referer", fmt.Sprintf("%s://%s", u.Scheme, u.Host))
	// req.Header.Set("Accept", "*/*")
	// req.Header.Set("Accept-Encoding", "gzip")
	req.Header.Set("User-Agent", RandomUA())
	return client.Do(req)
}
