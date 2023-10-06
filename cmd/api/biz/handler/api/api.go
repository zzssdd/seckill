// Code generated by hertz generator.

package api

import (
	"context"
	"seckill/cmd/api/biz/model/product"
	"seckill/cmd/api/biz/model/seckill"
	"seckill/cmd/api/biz/model/user"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// Registry .
// @router /registry [POST]
func Registry(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.BaseRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(user.BaseResponse)

	c.JSON(consts.StatusOK, resp)
}

// Login .
// @router /login [POST]
func Login(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.BaseRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(user.LoginResponse)

	c.JSON(consts.StatusOK, resp)
}

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

	c.JSON(consts.StatusOK, resp)
}

// DoSeckill .
// @router /seckill [POST]
func DoSeckill(ctx context.Context, c *app.RequestContext) {
	var err error
	var req seckill.SeckillRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(seckill.SeckillResponse)

	c.JSON(consts.StatusOK, resp)
}

// Submit .
// @router /submit [POST]
func Submit(ctx context.Context, c *app.RequestContext) {
	var err error
	var req seckill.SubmitRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(seckill.BaseResponse)

	c.JSON(consts.StatusOK, resp)
}
