package automod

import (
	"context"
	"time"

	"github.com/hashicorp/golang-lru/v2/expirable"
)

type CacheStore interface {
	Get(ctx context.Context, name, key string) (string, error)
	Set(ctx context.Context, name, key string, val string) error
}

type MemCacheStore struct {
	Data *expirable.LRU[string, string]
}

func NewMemCacheStore(capacity int, ttl time.Duration) MemCacheStore {
	return MemCacheStore{
		Data: expirable.NewLRU[string, string](capacity, nil, ttl),
	}
}

func (s MemCacheStore) Get(ctx context.Context, name, key string) (string, error) {
	v, ok := s.Data.Get(name + "/" + key)
	if !ok {
		return "", nil
	}
	return v, nil
}

func (s MemCacheStore) Set(ctx context.Context, name, key string, val string) error {
	s.Data.Add(name+"/"+key, val)
	return nil
}
