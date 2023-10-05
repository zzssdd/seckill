package order

import (
	"context"
	order "seckill/kitex_gen/order"
)

// OrderImpl implements the last service interface defined in the IDL.
type OrderImpl struct{}

// OrderAdd implements the OrderImpl interface.
func (s *OrderImpl) OrderAdd(ctx context.Context, req *order.OrderInfo) (resp *order.BaseResponse, err error) {
	// TODO: Your code here...
	return
}

// OrderStatusAdd implements the OrderImpl interface.
func (s *OrderImpl) OrderStatusAdd(ctx context.Context, req *order.OrderInfo) (resp *order.BaseResponse, err error) {
	// TODO: Your code here...
	return
}

// Try implements the OrderImpl interface.
func (s *OrderImpl) Try(ctx context.Context, req *order.OrderInfo) (resp *order.BaseResponse, err error) {
	// TODO: Your code here...
	return
}

// Commit implements the OrderImpl interface.
func (s *OrderImpl) Commit(ctx context.Context, req *order.OrderInfo) (resp *order.BaseResponse, err error) {
	// TODO: Your code here...
	return
}

// Cancal implements the OrderImpl interface.
func (s *OrderImpl) Cancal(ctx context.Context, req *order.OrderInfo) (resp *order.BaseResponse, err error) {
	// TODO: Your code here...
	return
}
