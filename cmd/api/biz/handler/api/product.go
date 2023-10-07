package api

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"seckill/cmd/rpc"
	"seckill/kitex_gen/product"
	rpcProduct "seckill/kitex_gen/product"
	. "seckill/pkg/log"
)

// ProductAdd .
// @router /product [POST]
func ProductAdd(ctx context.Context, c *app.RequestContext) {
	var err error
	var req product.ProductInfo
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(product.BaseResponse)

	rpcReq := &rpcProduct.ProductInfo{
		Name:  req.Name,
		Price: req.Price,
		Pic:   req.Pic,
		Des:   req.Des,
		Num:   req.Num,
	}

	rpcResp, err := rpc.ProductClient.AddProduct(ctx, rpcReq)
	resp.Code = rpcResp.Code
	resp.Msg = rpcResp.Msg
	if err != nil {
		Log.Errorln("product service's productAdd return err:", err)
		return
	}
	c.JSON(consts.StatusOK, resp)
}

// ProductInfo .
// @router /product [GET]
func ProductInfo(ctx context.Context, c *app.RequestContext) {
	var err error
	var req product.IdRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(product.ProductResponse)

	rpcReq := &product.IdRequest{
		Id: req.Id,
	}

	rpcResp, err := rpc.ProductClient.ProductInfo(ctx, rpcReq)
	resp.Code = rpcResp.Code
	resp.Msg = rpcResp.Msg
	resp.Product = rpcResp.Product
	if err != nil {
		Log.Errorln("product service's productInfo return err:", err)
		return
	}
	c.JSON(consts.StatusOK, resp)
}

// ProductList .
// @router /list [GET]
func ProductList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req product.ListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(product.ListResponse)

	rpcReq := &rpcProduct.ListRequest{
		Offset: req.Offset,
		Limit:  req.Limit,
	}
	rpcResp, err := rpc.ProductClient.ListProduct(ctx, rpcReq)
	resp.Code = rpcResp.Code
	resp.Msg = rpcResp.Msg
	resp.Products = rpcResp.Products
	if err != nil {
		Log.Errorln("product service's productList return err:", err)
		return
	}
	c.JSON(consts.StatusOK, resp)
}
