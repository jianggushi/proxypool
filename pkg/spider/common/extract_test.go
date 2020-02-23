package common

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func Test_extractFromTable(t *testing.T) {
	file, err := os.Open("example/goubanjia-20200223.html")
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
	for _, record := range records {
		t.Log(record)
	}
}

func Test_extractProxy(t *testing.T) {
	file, err := os.Open("example/xicidaili-20200223.html")
	if err != nil {
		t.Fatal(err)
	}
	rules := map[string]int{
		"host":      1,
		"port":      2,
		"anonymity": 4,
		"scheme":    5,
	}
	proxies, err := extractProxy(file, rules)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(proxies)
}

func Test_removeHiddenElement(t *testing.T) {
	html := `
	<html>
		<body>
			<table>
				<tr>
					<td>
						<p class="1" style="display:block;">1</p>
						<p class="2" style="display:none;">2</p>
					</td>
				</tr>
			</table>
		</body>
	</html>
	`
	doc, err := goquery.NewDocumentFromReader(bytes.NewBufferString(html))
	if err != nil {
		t.Fatal(err)
	}
	doc.Find("[style]").Each(func(i int, s *goquery.Selection) {
		style, _ := s.Attr("style")
		if strings.Contains(style, "display:none") {
			s.Remove()
		}
	})
	t.Log(goquery.OuterHtml(doc.Find("body")))
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
