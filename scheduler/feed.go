package scheduler

import (
	log "github.com/sirupsen/logrus"
	"github.com/zkqiang/kuaizhi-feeder/cache"
	"github.com/zkqiang/kuaizhi-feeder/config"
	"github.com/zkqiang/kuaizhi-feeder/enum"
)

func Feed(source enum.Source) bool {
	log.Debug("start feed")
	token, err := config.GetToken(source)
	if err != nil {
		log.Fatal(err)
	}
	log.Debugf("get token %s", token)
	job := getJob(token)
	if job == nil {
		log.Debugf("no job from token %s", token)
		return false
	}
	jobId := job.JobId
	log.Debugf("get job %s from token %s", jobId, token)
	feed := cache.Feed.Get(source)
	if feed != nil {
		card := cache.Card.Get(jobId)
		if (*feed)[0].Url == card.Url {
			log.Debugf("have no new card "+
				"from token %s, jobId %s", token, jobId)
			return false
		}
	}
	cards := crawl(source)
	if cards == nil {
		log.Errorf("crawl zero card from %s", source)
		return false
	}
	card := cache.Card.Get(jobId)
	if card != nil && (*cards)[0].Url == card.Url {
		log.Debugf("have no new card "+
			"from token %s, jobId %s", token, jobId)
		return false
	}
	cache.Feed.Set(source, cards)
	postJob(token, jobId, cards)
	log.Debugf("post job %s cards %s from token %s", jobId, cards, token)
	cache.Card.Set(jobId, &(*cards)[0])
	return true
}
