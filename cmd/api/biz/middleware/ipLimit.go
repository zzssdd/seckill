package middleware

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"seckill/dao/cache"
	"seckill/utils"
)

var (
	bloom *utils.BloomFilter
	rds   cache.Cache
)

func IpLimit() app.HandlerFunc {
	if bloom == nil {
		bloom = utils.NewBloomFilter()
		rds = cache.NewCache()
	}
	return func(ctx context.Context, c *app.RequestContext) {
		clientIp := c.ClientIP()
		if !bloom.IsContains(clientIp) || !rds.IP.IpLimit(ctx, clientIp) {
			c.Next(ctx)
		} else {
			c.Abort()
		}
	}
}
