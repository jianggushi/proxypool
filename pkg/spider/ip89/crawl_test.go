package ip89

import (
	"os"
	"testing"
	"time"

	"github.com/jianggushi/proxypool/pkg/model"
)

func Test_Crawl(t *testing.T) {
	ch := make(chan *model.Proxy, 100)
	go func() {
		for proxy := range ch {
			t.Log(proxy)
		}
	}()
	SpiderIP89.Crawl(ch)
	time.Sleep(5 * time.Second)
}

func Test_Find(t *testing.T) {
	file, err := os.Open("ip3366-20200219.html")
	if err != nil {
		t.Fatal(err)
	}
	proxies, err := SpiderIP89.find(file)
	if err != nil {
		t.Fatal(err)
	}
	for index, proxy := range proxies {
		t.Error(index, proxy)
	}
}

func Benchmark_Find(b *testing.B) {
	file, err := os.Open("ip3366-20200219.html")
	if err != nil {
		b.Fatal(err)
	}
	for i := 0; i < b.N; i++ {
		SpiderIP89.find(file)
	}
}
