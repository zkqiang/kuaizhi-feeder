package cache

import (
	"github.com/bluele/gcache"
	log "github.com/sirupsen/logrus"
	"github.com/zkqiang/kuaizhi-feeder/model"
)

var Card = &cardCache{
	gcache.New(128).LFU().Build(),
}

type cardCache struct {
	holder gcache.Cache
}

func (cache *cardCache) Set(jobId string, card *model.Card) {
	if err := cache.holder.Set(jobId, card); nil != err {
		log.Errorf("Set jobId=%s into cache failed: %s", jobId, err)
	}
}

func (cache *cardCache) Get(jobId string) *model.Card {
	ret, err := cache.holder.Get(jobId)
	if err != nil && err != gcache.KeyNotFoundError {
		return nil
	}
	if ret == nil {
		return nil
	}
	return ret.(*model.Card)
}
