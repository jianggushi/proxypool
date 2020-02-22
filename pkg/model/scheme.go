package model

import "strings"

type Scheme int

func (s Scheme) String() string {
	switch s {
	case Http:
		return "http"
	case Https:
		return "https"
	case Socks4:
		return "socks4"
	case Socks5:
		return "socks5"
	default:
		return "unknown"
	}
}

const (
	UnknownS Scheme = iota
	Http
	Https
	Socks4
	Socks5
)

func ParseScheme(scheme string) Scheme {
	scheme = strings.ToUpper(strings.TrimSpace(scheme))
	switch scheme {
	case "HTTP":
		return Http
	case "HTTPS":
		return Https
	case "SOCKS4":
		return Socks4
	case "SOCKS5":
		return Socks5
	default:
		return UnknownS
	}
}
