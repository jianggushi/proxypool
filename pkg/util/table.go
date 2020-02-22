package util

import (
	"errors"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Find parse the html find proxy
func ReadTable(s *goquery.Selection, header bool) ([][]string, error) {
	records := make([][]string, 0)
	// thead
	if header == true {
		item := make([]string, 0)
		s.Find("thead tr th").Each(func(i int, s *goquery.Selection) {
			item = append(item, strings.TrimSpace(s.Text()))
		})
		if len(item) == 0 {
			return nil, errors.New("not found thead")
		}
		records = append(records, item)
	}
	// tbody
	s.Find("tbody tr").Each(func(i int, s *goquery.Selection) {
		item := make([]string, 0)
		s.Find("td").Each(func(i int, s *goquery.Selection) {
			item = append(item, strings.TrimSpace(s.Text()))
		})
		records = append(records, item)
	})
	if len(records) == 0 {
		return nil, errors.New("not found table record")
	}
	return records, nil
}
