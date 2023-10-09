// Code generated by Kitex v0.7.2. DO NOT EDIT.

package order

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	order "seckill/kitex_gen/order"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	OrderAdd(ctx context.Context, req *order.OrderInfo, callOptions ...callopt.Option) (r *order.BaseResponse, err error)
	OrderStatusAdd(ctx context.Context, req *order.OrderInfo, callOptions ...callopt.Option) (r *order.BaseResponse, err error)
	OrderCancel(ctx context.Context, req *order.OrderCancelRequest, callOptions ...callopt.Option) (r *order.OrderCancelResponse, err error)
	Try(ctx context.Context, req *order.OrderInfo, callOptions ...callopt.Option) (r *order.BaseResponse, err error)
	Commit(ctx context.Context, req *order.OrderInfo, callOptions ...callopt.Option) (r *order.BaseResponse, err error)
	Cancal(ctx context.Context, req *order.OrderInfo, callOptions ...callopt.Option) (r *order.BaseResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kOrderClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kOrderClient struct {
	*kClient
}

func (p *kOrderClient) OrderAdd(ctx context.Context, req *order.OrderInfo, callOptions ...callopt.Option) (r *order.BaseResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.OrderAdd(ctx, req)
}

func (p *kOrderClient) OrderStatusAdd(ctx context.Context, req *order.OrderInfo, callOptions ...callopt.Option) (r *order.BaseResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.OrderStatusAdd(ctx, req)
}

func (p *kOrderClient) OrderCancel(ctx context.Context, req *order.OrderCancelRequest, callOptions ...callopt.Option) (r *order.OrderCancelResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.OrderCancel(ctx, req)
}

func (p *kOrderClient) Try(ctx context.Context, req *order.OrderInfo, callOptions ...callopt.Option) (r *order.BaseResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Try(ctx, req)
}

func (p *kOrderClient) Commit(ctx context.Context, req *order.OrderInfo, callOptions ...callopt.Option) (r *order.BaseResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Commit(ctx, req)
}

func (p *kOrderClient) Cancal(ctx context.Context, req *order.OrderInfo, callOptions ...callopt.Option) (r *order.BaseResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Cancal(ctx, req)
}
