package request

import (
	"net/http"
	"time"
)

var (
	client = &http.Client{
		Timeout: 5 * time.Second,
	}
)

func Get(url string) (resp *http.Response, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	// req.Header.Set("Accept", "*/*")
	// req.Header.Set("Accept-Encoding", "gzip")
	req.Header.Set("User-Agent", RandomUA())
	return client.Do(req)
}
