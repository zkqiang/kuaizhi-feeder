package scheduler

import (
	log "github.com/sirupsen/logrus"
	"github.com/zkqiang/kuaizhi-feeder/enum"
	"github.com/zkqiang/kuaizhi-feeder/model"
	"github.com/zkqiang/kuaizhi-feeder/spider"
)

func crawl(source enum.Source) *[]model.Card {
	var cards *[]model.Card
	var err error
	switch source {
	case enum.Youyanshe:
		cards, err = spider.CrawlYouyanshe(0)
	}
	if err != nil {
		log.Errorf("%s crawl error: %s", source, err)
		return nil
	}
	if len(*cards) == 0 {
		log.Errorf("%s crawl nil data", source)
		return nil
	}
	return cards
}
