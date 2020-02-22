package common

import (
	"os"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func Test_extractFromTable(t *testing.T) {
	file, err := os.Open("example/kuaidaili-20200219.html")
	if err != nil {
		t.Fatal(err)
	}
	doc, err := goquery.NewDocumentFromReader(file)
	if err != nil {
		t.Fatal(err)
	}
	records, err := extractFromTable(doc)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(records)
}

func Test_extractProxy(t *testing.T) {
	file, err := os.Open("example/kuaidaili-20200219.html")
	if err != nil {
		t.Fatal(err)
	}
	rules := map[string]int{
		"host":      0,
		"port":      1,
		"anonymity": 2,
		"scheme":    3,
	}
	proxies, err := extractProxy(file, rules)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(proxies)
}

func Benchmark_extractFromTable(b *testing.B) {
	file, err := os.Open("example/kuaidaili-20200219.html")
	if err != nil {
		b.Fatal(err)
	}
	doc, err := goquery.NewDocumentFromReader(file)
	if err != nil {
		b.Fatal(err)
	}
	for i := 0; i < b.N; i++ {
		extractFromTable(doc)
	}
}
