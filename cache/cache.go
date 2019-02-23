package cache

import (
	"github.com/go-redis/redis"
	"github.com/godcong/wego/cache"
)

var caches = map[string]cache.Cache{}

// DefaultCache ...
func DefaultCache() cache.Cache {
	return RedisCache("localhost:6379", "2rXfzaNKqX1b", 1)
}

// RedisCache ...
func RedisCache(addr, pass string, db int) cache.Cache {
	redis := cache.NewRedisCache(&redis.Options{
		Addr:     addr,
		Password: pass, // no password set
		DB:       db,   // use default DB
	})
	return redis
}

// InitWegoCache ...
func InitWegoCache(c cache.Cache) {
	caches["wego"] = c
	cache.RegisterCache(c)
}
