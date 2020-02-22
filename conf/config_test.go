package conf

import (
	"testing"

	"github.com/BurntSushi/toml"
)

func Test_SpiderTomlDecode(t *testing.T) {
	var conf Config
	_, err := toml.DecodeFile("spider.toml", &conf)
	if err != nil {
		t.Fatal(err)
	}
	for _, sp := range conf.Spider {
		t.Log(sp)
	}
}
