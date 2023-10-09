package api

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"seckill/cmd/api/biz/model/seckill"
	"seckill/cmd/rpc"
	"seckill/conf"
	"seckill/dao/mq"
	rpcSeckill "seckill/kitex_gen/seckill"
	"seckill/pkg/errmsg"
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
	var orderStatusMQ *mq.OrderStatus
	err = orderStatusMQ.SetUp(conf.RabbitmqDSN)
	if err != nil {
		Log.Errorln("Set up order status failed:", err)
	}
	err = orderStatusMQ.DelayPublish(rpcResp.Id)
	if err != nil {
		Log.Errorln("Publish order status failed:", err)
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
	var orderMQ *mq.Order
	err = orderMQ.SetUp(conf.RabbitmqDSN)
	if err != nil {
		Log.Errorln("Set up order status failed:", err)
	}
	err = orderMQ.Publish(req.ID, req.UID, int(req.Pid), req.ReqTime)
	if err != nil {
		Log.Errorln("Publish order info failed:", err)
	}
	resp.Code = errmsg.Success
	resp.Msg = errmsg.GetMsg(errmsg.Success)
	c.JSON(consts.StatusOK, resp)
}
