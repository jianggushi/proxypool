package kuaidaili

import (
	"io/ioutil"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/jianggushi/proxypool/pkg/model"
	"github.com/jianggushi/proxypool/pkg/request"
)

func Test_RequestGet2(t *testing.T) {
	resp, err := request.Get("https://www.kuaidaili.com/free/inha/1/")
	if err != nil {
		log.Errorf("http get: %v", err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Errorf("status code: %d %s", resp.StatusCode, resp.Status)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}
	t.Log(string(body))
}

func Test_Crawl(t *testing.T) {
	ch := make(chan *model.Proxy, 100)
	go func() {
		for proxy := range ch {
			t.Log(proxy)
		}
	}()
	SpiderKuaiDaiLi.Crawl(ch)
	time.Sleep(5 * time.Second)
}

func Test_Find(t *testing.T) {
	file, err := os.Open("kuaidaili-20200219.html")
	if err != nil {
		t.Fatal(err)
	}
	SpiderKuaiDaiLi.find(file)
}

func Benchmark_Find(b *testing.B) {
	file, err := os.Open("kuaidaili-20200219.html")
	if err != nil {
		b.Fatal(err)
	}
	for i := 0; i < b.N; i++ {
		SpiderKuaiDaiLi.find(file)
	}
}

func Test_Http(t *testing.T) {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	req, err := http.NewRequest("GET", "https://www.baidu.com", nil)
	if err != nil {
		log.Errorf("new request: %v", err)
		return
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Encoding", "gzip, deflate")
	req.Header.Set("User-Agent", request.RandomUA())
	resp, err := client.Do(req)
	if err != nil {
		log.Errorf("client do: %v", err)
		return
	}
	defer resp.Body.Close()
	t.Error(req.Header)
	t.Error(resp.Request.Header)
	if resp.StatusCode != 200 {
		log.Errorf("status code: %d %s", resp.StatusCode, resp.Status)
		return
	}
}
