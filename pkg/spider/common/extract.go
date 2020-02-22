package common

import (
	"errors"
	"io"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/jianggushi/proxypool/pkg/model"
)

func extractProxy(r io.Reader, rules map[string]int) ([]*model.Proxy, error) {
	// load html
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return nil, err
	}
	records, err := extractFromTable(doc)
	if err != nil {
		return nil, err
	}
	return mapping(records, rules), nil
}

func mapping(records [][]string, rules map[string]int) []*model.Proxy {
	proxies := make([]*model.Proxy, 0, len(records))
	for _, record := range records {
		proxy := &model.Proxy{}
		if k, ok := rules["host"]; ok {
			proxy.Host = record[k]
		}
		if k, ok := rules["port"]; ok {
			proxy.Port = record[k]
		}
		if k, ok := rules["scheme"]; ok {
			proxy.Scheme = model.ParseScheme(record[k])
		}
		if k, ok := rules["anonymity"]; ok {
			proxy.Anonymity = model.ParseAnonymity(record[k])
		}
		proxies = append(proxies, proxy)
	}
	return proxies
}

// Find parse the html find proxy
func extractFromTable(doc *goquery.Document) ([][]string, error) {
	records := make([][]string, 0)
	// tbody
	doc.Find("table tbody tr").Each(func(i int, s *goquery.Selection) {
		item := make([]string, 0)
		s.Find("td").Each(func(i int, s *goquery.Selection) {
			item = append(item, strings.TrimSpace(s.Text()))
		})
		records = append(records, item)
	})
	if len(records) == 0 {
		return nil, errors.New("not found proxy record")
	}
	return records, nil
}
