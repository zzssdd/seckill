package seckill

import (
	"context"
	"seckill/conf"
	"seckill/dao/cache"
	"seckill/dao/db"
	"seckill/dao/mq"
	seckill "seckill/kitex_gen/seckill"
	"seckill/pkg/errmsg"
	. "seckill/pkg/log"
	"seckill/utils"
)

// SeckillImpl implements the last service interface defined in the IDL.
type SeckillImpl struct {
	dao   db.Dao
	cache cache.Cache
}

// DoSeckill implements the SeckillImpl interface.
func (s *SeckillImpl) DoSeckill(ctx context.Context, req *seckill.SeckillRequest) (resp *seckill.SeckillResponse, err error) {
	// TODO: Your code here...
	resp = new(seckill.SeckillResponse)
	if !s.cache.Prduct.PreSubNum(ctx, int(req.Pid), int(req.Num)) {
		resp.Code = errmsg.Sellout
		resp.Msg = errmsg.GetMsg(errmsg.Sellout)
		return
	}
	id := utils.GenID()
	s.cache.Orer.StoreOrderStatus(ctx, id, req.Uid, int(req.Pid))
	resp.Id = id
	resp.Code = errmsg.Success
	resp.Msg = errmsg.GetMsg(errmsg.Success)
	return
}

// Submit implements the SeckillImpl interface.
func (s *SeckillImpl) Submit(ctx context.Context, req *seckill.SubmitRequest) (resp *seckill.BaseResponse, err error) {
	// TODO: Your code here...
	resp = new(seckill.BaseResponse)

	var orderMQ *mq.Order
	err = orderMQ.SetUp(conf.RabbitmqDSN)
	if err != nil {
		Log.Errorln("set up order mq err:", err)
		return
	}
	msgChan, err := orderMQ.Consume()
	if err != nil {
		Log.Errorln("consume from order mq err:", err)
		return
	}
	for msg := range msgChan {
		msg.Ack(true)
		msg.Body
	}
	return
}
