package cache

import (
	"github.com/redis/go-redis/v9"
	"seckill/conf"
)

var rds *redis.Client

type Cache struct {
}

func NewCache() Cache {
	if rds == nil {
		rds = redis.NewClient(&redis.Options{
			Addr:     conf.RedisDSN,
			Password: "",
			DB:       0,
		})
	}
	return Cache{}
}