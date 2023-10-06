// Code generated by Kitex v0.7.2. DO NOT EDIT.

package product

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	product "seckill/kitex_gen/product"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	AddProduct(ctx context.Context, req *product.ProductInfo, callOptions ...callopt.Option) (r *product.BaseResponse, err error)
	ListProduct(ctx context.Context, req *product.ListRequest, callOptions ...callopt.Option) (r *product.ListResponse, err error)
	ProductInfo(ctx context.Context, req *product.IdRequest, callOptions ...callopt.Option) (r *product.ProductResponse, err error)
	Try(ctx context.Context, req *product.BuyRequest, callOptions ...callopt.Option) (r *product.BaseResponse, err error)
	Commit(ctx context.Context, req *product.BuyRequest, callOptions ...callopt.Option) (r *product.BaseResponse, err error)
	Cancel(ctx context.Context, req *product.BuyRequest, callOptions ...callopt.Option) (r *product.BaseResponse, err error)
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
	return &kProductClient{
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

type kProductClient struct {
	*kClient
}

func (p *kProductClient) AddProduct(ctx context.Context, req *product.ProductInfo, callOptions ...callopt.Option) (r *product.BaseResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AddProduct(ctx, req)
}

func (p *kProductClient) ListProduct(ctx context.Context, req *product.ListRequest, callOptions ...callopt.Option) (r *product.ListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ListProduct(ctx, req)
}

func (p *kProductClient) ProductInfo(ctx context.Context, req *product.IdRequest, callOptions ...callopt.Option) (r *product.ProductResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ProductInfo(ctx, req)
}

func (p *kProductClient) Try(ctx context.Context, req *product.BuyRequest, callOptions ...callopt.Option) (r *product.BaseResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Try(ctx, req)
}

func (p *kProductClient) Commit(ctx context.Context, req *product.BuyRequest, callOptions ...callopt.Option) (r *product.BaseResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Commit(ctx, req)
}

func (p *kProductClient) Cancel(ctx context.Context, req *product.BuyRequest, callOptions ...callopt.Option) (r *product.BaseResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Cancel(ctx, req)
}
