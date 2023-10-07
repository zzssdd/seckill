package cache

import (
	"context"
)

type Ip struct {
}

const IpTag = "ip"

func (i *Ip) IpLimit(ctx context.Context, ip string) bool {
	return rds.SIsMember(ctx, IpTag, ip).Val()
}
