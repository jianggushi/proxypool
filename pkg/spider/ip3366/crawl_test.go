package ip3366

import (
	"os"
	"testing"
)

func Test_Find(t *testing.T) {
	file, err := os.Open("ip3366-20200219.html")
	if err != nil {
		t.Fatal(err)
	}
	proxies, err := SpiderIP3366.find(file)
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
		SpiderIP3366.find(file)
	}
}
