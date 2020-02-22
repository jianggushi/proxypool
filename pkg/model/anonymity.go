package model

import "strings"

type Anonymity int

func (a Anonymity) String() string {
	switch a {
	case Transparent:
		return "transparent"
	case Anonymous:
		return "anonymous"
	case HighAnonymity:
		return "highAnonymity"
	default:
		return "unknown"
	}
}

const (
	UnknownA      Anonymity = iota
	Transparent             // 透明
	Anonymous               // 匿名
	HighAnonymity           // 高匿名
)

func ParseAnonymity(anonymity string) Anonymity {
	anonymity = strings.TrimSpace(anonymity)
	switch {
	case strings.Contains(anonymity, "高匿"):
		return HighAnonymity
	case strings.Contains(anonymity, "匿名"):
		return Anonymous
	case strings.Contains(anonymity, "透明"):
		return Transparent
	default:
		return UnknownA
	}
}
