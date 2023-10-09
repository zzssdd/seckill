package task

import (
	"context"
	"seckill/cmd/rpc"
	"seckill/kitex_gen/order"
)

func StartBackgroundTask(ctx context.Context) {
	req := &order.OrderInfo{}
	rpc.OrderClient.OrderAdd(ctx, req)
}
