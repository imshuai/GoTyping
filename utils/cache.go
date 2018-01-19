package utils

import (
	"sync"
	"time"

	"github.com/imshuai/kvcache"
)

type (
	CacheBucket struct {
		kvcache.Bucket
		keys map[string]time.Time
		lk   *sync.RWMutex
	}
)

func (cb *CacheBucket) SetWithExpireTime(key string, obj kvcache.ObjectCacher, t time.Duration) {

}

func init() {

}
