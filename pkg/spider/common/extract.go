package common

import (
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/jianggushi/proxypool/pkg/model"
)

func extractProxy(r io.Reader, rule map[string]int) ([]*model.Proxy, error) {
	// load html
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return nil, err
	}
	records, err := extractFromTable(doc)
	if err != nil {
		return nil, err
	}
	return mapping(records, rule), nil
}

func mapping(records [][]string, rule map[string]int) []*model.Proxy {
	proxies := make([]*model.Proxy, 0, len(records))
	for _, record := range records {
		proxy := &model.Proxy{}
		if k, ok := rule["host"]; ok {
			proxy.Host = record[k]
		}
		if k, ok := rule["port"]; ok {
			proxy.Port = record[k]
		}
		if k, ok := rule["scheme"]; ok {
			proxy.Scheme = model.ParseScheme(record[k])
		}
		if k, ok := rule["proxy"]; ok {
			proxy.Proxy = record[k]
		}
		if k, ok := rule["anonymity"]; ok {
			proxy.Anonymity = model.ParseAnonymity(record[k])
		}

		if proxy.Host == "" &&
			proxy.Port == "" &&
			proxy.Proxy != "" {
			items := strings.Split(proxy.Proxy, ":")
			proxy.Host = items[0]
			proxy.Port = items[len(items)-1]
		}
		if proxy.Host != "" && proxy.Port != "" &&
			proxy.Proxy == "" {
			proxy.Proxy = fmt.Sprintf("%s:%s", proxy.Host, proxy.Port)
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
