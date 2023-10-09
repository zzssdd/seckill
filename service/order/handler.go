package order

import (
	"context"
	"encoding/json"
	"seckill/conf"
	"seckill/dao/cache"
	"seckill/dao/db"
	"seckill/dao/mq"
	"seckill/dao/mq/model"
	order "seckill/kitex_gen/order"
	"seckill/pkg/errmsg"
	. "seckill/pkg/log"
	"time"
)

// OrderImpl implements the last service interface defined in the IDL.
type OrderImpl struct {
	dao   db.Dao
	cache cache.Cache
}

// OrderAdd implements the OrderImpl interface.
func (s *OrderImpl) OrderAdd(ctx context.Context, req *order.OrderInfo) (resp *order.BaseResponse, err error) {
	// TODO: Your code here...
	resp = new(order.BaseResponse)
	var orderMQ *mq.Order
	msgChan, err := orderMQ.Consume()
	if err != nil {
		Log.Errorln("get order consumer err:", err)
		return
	}
	var orderInfo *model.Order
	for v := range msgChan {
		err = json.Unmarshal(v.Body, &orderInfo)
		if err != nil {
			Log.Errorln("consume order mq err:", err)
			return
		}
		if time.Now().Unix()-orderInfo.TimeStamp >= 300000 {
			err = s.cache.Orer.StoreOrder(ctx, orderInfo.Id, orderInfo.Uid, orderInfo.Pid, orderInfo.TimeStamp)
			if err != nil {
				Log.Errorln("store order info into redis err:", err)
			}
		} else {
			affect, err := s.dao.Order.OrderAdd(orderInfo.Uid, orderInfo.Pid, orderInfo.TimeStamp)
			if err != nil {
				Log.Errorln("insert order info into mysql err:", err)
			} else if affect == 0 {
				Log.Infoln("repeat create order")
			}
		}
	}
	return
}

// OrderStatusAdd implements the OrderImpl interface.
func (s *OrderImpl) OrderStatusAdd(ctx context.Context, req *order.OrderInfo) (resp *order.BaseResponse, err error) {
	// TODO: Your code here...
	resp = new(order.BaseResponse)
	var orderStatusMQ *mq.Order
	err = orderStatusMQ.SetUp(conf.RabbitmqDSN)
	if err != nil {
		Log.Errorln("set up orderStatusMQ err:", err)
		return
	}
	err = s.cache.Orer.StoreOrder(ctx, req.Id, req.Uid, int(req.Pid), req.TimeStamp)
	if err != nil {
		Log.Errorln("store order status failed,err:", err)
		resp.Code = errmsg.Error
		resp.Msg = errmsg.GetMsg(errmsg.Error)
		return
	}
	resp.Code = errmsg.Success
	resp.Msg = errmsg.GetMsg(errmsg.Success)
	return
}

// OrderCancel implements the OrderImpl interface.
func (s *OrderImpl) OrderCancel(ctx context.Context, req *order.OrderCancelRequest) (resp *order.OrderCancelResponse, err error) {
	// TODO: Your code here...
	var orderMQ *mq.OrderStatus
	msgChan, err := orderMQ.DelayConsume()
	if err != nil {
		Log.Errorln("get order consume err:", err)
		return
	}
	var id int64
	for msg := range msgChan {
		id = int64(msg.Body[0])
		_, _, status := s.cache.Orer.GetOrderStatus(ctx, id)
		if status != "finished" {
			s.cache.Orer.DelOrderStatus(ctx, id)
		}
	}
	return
}

// Try implements the OrderImpl interface.
func (s *OrderImpl) Try(ctx context.Context, req *order.OrderInfo) (resp *order.BaseResponse, err error) {
	// TODO: Your code here...
	resp = new(order.BaseResponse)
	err = s.cache.Orer.DoingOrderStatus(ctx, req.Id, req.Uid, int(req.Pid))
	if err != nil {
		Log.Errorln("update order status into doing failed,err:", err)
		resp.Code = errmsg.Error
		resp.Msg = errmsg.GetMsg(errmsg.Error)
		return
	}
	resp.Code = errmsg.Success
	resp.Msg = errmsg.GetMsg(errmsg.Success)
	return
}

// Commit implements the OrderImpl interface.
func (s *OrderImpl) Commit(ctx context.Context, req *order.OrderInfo) (resp *order.BaseResponse, err error) {
	// TODO: Your code here...
	resp = new(order.BaseResponse)
	err = s.cache.Orer.DelOrderStatus(ctx, req.Id)
	if err != nil {
		Log.Errorln("delete order status into doing failed,err:", err)
		resp.Code = errmsg.Error
		resp.Msg = errmsg.GetMsg(errmsg.Error)
		return
	}
	resp.Code = errmsg.Success
	resp.Msg = errmsg.GetMsg(errmsg.Success)
	return
}

// Cancal implements the OrderImpl interface.
func (s *OrderImpl) Cancal(ctx context.Context, req *order.OrderInfo) (resp *order.BaseResponse, err error) {
	// TODO: Your code here...
	resp = new(order.BaseResponse)
	err = s.cache.Orer.StoreOrderStatus(ctx, req.Id, req.Uid, int(req.Pid))
	if err != nil {
		Log.Errorln("update order status into created failed,err:", err)
		resp.Code = errmsg.Error
		resp.Msg = errmsg.GetMsg(errmsg.Error)
		return
	}
	resp.Code = errmsg.Success
	resp.Msg = errmsg.GetMsg(errmsg.Success)
	return
}
