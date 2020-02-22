package conf

import (
	log "github.com/sirupsen/logrus"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Spiders []Spider `toml:"spider"`
}

type Spider struct {
	Name string         `toml:"name"`
	Url  string         `toml:"url"`
	Rule map[string]int `toml:"rule"`
}

var (
	Conf Config
)

func init() {
	file := "conf/spider.toml"
	_, err := toml.DecodeFile(file, &Conf)
	if err != nil {
		log.Fatalf("decode %s: %v", file, err)
	}
}
