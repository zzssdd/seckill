package task

import (
	"context"
	"seckill/cmd/rpc"
	"seckill/kitex_gen/order"
)

func StartDelayTask(ctx context.Context) {
	req := &order.OrderCancelRequest{}
	rpc.OrderClient.OrderCancel(ctx, req)
}
