package api

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"seckill/cmd/api/biz/model/seckill"
	"seckill/cmd/rpc"
	rpcSeckill "seckill/kitex_gen/seckill"
	. "seckill/pkg/log"
)

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

	rpcReq := &rpcSeckill.SeckillRequest{
		Uid: req.UID,
		Pid: req.Pid,
		Num: req.Num,
	}

	rpcResp, err := rpc.SeckillClient.DoSeckill(ctx, rpcReq)

	resp.Code = rpcResp.Code
	resp.Msg = rpcResp.Msg
	resp.ID = rpcResp.Id
	if err != nil {
		Log.Errorln("Seckill service's Doseckill return err:", err)
		return
	}
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
