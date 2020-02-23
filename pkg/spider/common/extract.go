package common

import (
	"errors"
	"fmt"
	"io"
	"regexp"
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
// <table>
//   <thead>
//     <tr></tr>
//   </thead>
//   <tbody>
//     <tr></tr>
//   </tbody>
// </table
func extractFromTable(doc *goquery.Document) ([][]string, error) {
	records := make([][]string, 0)
	// tbody
	doc.Find("table tbody tr").Each(func(i int, s *goquery.Selection) {
		item := make([]string, 0)
		s.Find("td").Each(func(i int, s *goquery.Selection) {
			// remove hidden element that confuse
			// see http://www.goubanjia.com/
			s.Find("[style]").Each(func(i int, s *goquery.Selection) {
				style, _ := s.Attr("style")
				matched, _ := regexp.MatchString(`display:\s*none`, style)
				if matched {
					s.Remove()
				}
			})
			// format text
			text := s.Text()
			text = regexp.MustCompile(`\s{2,}`).ReplaceAllString(text, " ")
			item = append(item, strings.TrimSpace(text))
		})
		// exclude empty item
		if len(item) != 0 {
			records = append(records, item)
		}
	})
	if len(records) == 0 {
		return nil, errors.New("not found proxy record")
	}
	return records, nil
}
