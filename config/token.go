package config

import (
	"errors"
	"github.com/zkqiang/kuaizhi-feeder/enum"
)

type Token string

const (
	youyanshe = "7a459520d17c081a6452b38617d44b55"
)

func GetToken(source enum.Source) (string, error) {
	switch source {
	case enum.Youyanshe:
		return youyanshe, nil
	default:
		return "", errors.New("Unkown source:" + source.String())
	}
}
