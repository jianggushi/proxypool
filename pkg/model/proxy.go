package model

import (
	"time"
)

// Proxy define the proxy ip
type Proxy struct {
	Host      string    `json:"host"`      // hostname or ip address
	Port      string    `json:"port"`      // port number
	Scheme    Scheme    `json:"scheme"`    // http or https
	Proxy     string    `json:"proxy"`     // proxy ip:port
	Anonymity Anonymity `json:"anonymity"` // anonymity
	Ping      int       `json:"ping"`      // ping time
	Transfer  int       `json:"transfer"`  // transfer time
	Country   string    `json:"country"`   // 国家
	Region    string    `json:"region"`    // 地区
	Isp       string    `json:"isp"`       // 运营商
	From      string    `json:"from"`      // from which one website
	Created   time.Time `json:"created"`   // create time
	Checked   time.Time `json:"checked"`   // last check time
}

func (p Proxy) String() string {
	return p.Proxy
}
