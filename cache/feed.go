package cache

import (
	"github.com/bluele/gcache"
	log "github.com/sirupsen/logrus"
	"github.com/zkqiang/kuaizhi-feeder/enum"
	"github.com/zkqiang/kuaizhi-feeder/model"
	"time"
)

var Feed = &feedCache{
	gcache.New(32).LFU().Expiration(time.Minute).Build(),
}

type feedCache struct {
	holder gcache.Cache
}

func (cache *feedCache) Set(source enum.Source, cards *[]model.Card) {
	if err := cache.holder.Set(source.String(), cards); nil != err {
		log.Errorf("set source %s into cache failed: %s", source, err)
	}
}

func (cache *feedCache) Get(source enum.Source) *[]model.Card {
	ret, err := cache.holder.Get(source.String())
	if err != nil && err != gcache.KeyNotFoundError {
		return nil
	}
	if ret == nil {
		return nil
	}
	return ret.(*[]model.Card)
}
