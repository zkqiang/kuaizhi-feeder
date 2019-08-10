package enum

import "strings"

type Source int

const (
	Youyanshe Source = iota
)

func (source Source) String() string {
	switch source {
	case Youyanshe:
		return "Youyanshe"
	default:
		return "Unknow"
	}
}

func GetSource(srcStr string) Source {
	switch strings.ToLower(srcStr) {
	case "youyanshe":
		return Youyanshe
	default:
		return -1
	}
}
