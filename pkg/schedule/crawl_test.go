package schedule

import (
	"fmt"
	"testing"
)

func Test_Crawl(t *testing.T) {
	// go func() {
	// 	for proxy := range proxych {
	// 		t.Error(proxy)
	// 	}
	// }()
	// go check()
	go Crawl()
	// t.Error("")
	// time.Sleep(30 * time.Second)
	for proxy := range crawlChan {
		fmt.Println(proxy)
	}
}
