package middleware

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
)

func IpLimit() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		clientIp := c.ClientIP()
		fmt.Println(clientIp)
	}
}
