package storage

import (
	"fmt"

	"github.com/BasbustDama/tiny-url/pkg/cache"
)

type Storage interface {
	Put(shortUrl, longUrl string) error
	Get(shortUrl string) (string, error)
}

type cacheStorage struct {
	cache *cache.Cache
}

func New(cache *cache.Cache) Storage {
	return &cacheStorage{cache: cache}
}

func (storage *cacheStorage) Get(shortUrl string) (string, error) {
	value, found := storage.cache.Get(shortUrl)
	if !found {
		return "", nil
	}

	longUrl, ok := value.(string)
	if !ok {
		return "", fmt.Errorf("%s url value in storage is broken", shortUrl)
	}

	return longUrl, nil
}

func (storage *cacheStorage) Put(shortUrl string, longUrl string) error {
	storage.cache.Set(shortUrl, longUrl, 0)
	return nil
}
