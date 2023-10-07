package api

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"seckill/cmd/rpc"
	"seckill/kitex_gen/user"
	rpcUser "seckill/kitex_gen/user"
	. "seckill/pkg/log"
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

	rpcReq := &rpcUser.BaseRequest{
		Email:    req.Email,
		Password: req.Password,
	}

	rpcResp, err := rpc.UserClient.Registry(ctx, rpcReq)
	resp.Code = rpcResp.Code
	resp.Msg = rpcResp.Msg
	if err != nil {
		Log.Errorln("user service's registry return err:", err)
		return
	}
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

	rpcReq := &rpcUser.BaseRequest{
		Email:    req.Email,
		Password: req.Password,
	}

	rpcResp, err := rpc.UserClient.Login(ctx, rpcReq)
	resp.Code = rpcResp.Code
	resp.Msg = rpcResp.Msg
	if err != nil {
		Log.Errorln("user service's login return err:", err)
		return
	}

	c.JSON(consts.StatusOK, resp)
}
