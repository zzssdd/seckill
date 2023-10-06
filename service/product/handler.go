package product

import (
	"context"
	product "seckill/kitex_gen/product"
)

// ProductImpl implements the last service interface defined in the IDL.
type ProductImpl struct{}

// AddProduct implements the ProductImpl interface.
func (s *ProductImpl) AddProduct(ctx context.Context, req *product.ProductInfo) (resp *product.BaseResponse, err error) {
	// TODO: Your code here...
	return
}

// ListProduct implements the ProductImpl interface.
func (s *ProductImpl) ListProduct(ctx context.Context, req *product.ListRequest) (resp *product.ListResponse, err error) {
	// TODO: Your code here...
	return
}

// Try implements the ProductImpl interface.
func (s *ProductImpl) Try(ctx context.Context, req *product.BuyRequest) (resp *product.BaseResponse, err error) {
	// TODO: Your code here...
	return
}

// Commit implements the ProductImpl interface.
func (s *ProductImpl) Commit(ctx context.Context, req *product.BuyRequest) (resp *product.BaseResponse, err error) {
	// TODO: Your code here...
	return
}

// Cancel implements the ProductImpl interface.
func (s *ProductImpl) Cancel(ctx context.Context, req *product.BuyRequest) (resp *product.BaseResponse, err error) {
	// TODO: Your code here...
	return
}
