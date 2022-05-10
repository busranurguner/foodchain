package cache

import (
	"time"

	"github.com/patrickmn/go-cache"
)

func init() {

}

var MemCache = cache.New(5*time.Minute, 10*time.Minute)

func Get(key string, duration time.Duration, acquire func() interface{}) interface{} {
	cachable, isCached := MemCache.Get(key)
	if isCached {
		return cachable
	}
	t := acquire()
	MemCache.Set(key, t, duration)
	return t

}
