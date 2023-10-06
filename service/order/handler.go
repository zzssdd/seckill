package order

import (
	"context"
	"seckill/dao/cache"
	"seckill/dao/db"
	order "seckill/kitex_gen/order"
	"seckill/pkg/errmsg"
	. "seckill/pkg/log"
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
	affect, err := s.dao.Order.OrderAdd(req.Uid, int(req.Pid), req.TimeStamp)
	if err != nil {
		Log.Errorln("add order failed,err:", err)
		resp.Code = errmsg.Error
		resp.Msg = errmsg.GetMsg(errmsg.Error)
		return
	}
	if affect == 0 {
		resp.Code = errmsg.OrderExist
		resp.Msg = errmsg.GetMsg(errmsg.OrderExist)
		return
	}
	resp.Code = errmsg.Success
	resp.Msg = errmsg.GetMsg(errmsg.Success)
	return
}

// OrderStatusAdd implements the OrderImpl interface.
func (s *OrderImpl) OrderStatusAdd(ctx context.Context, req *order.OrderInfo) (resp *order.BaseResponse, err error) {
	// TODO: Your code here...
	resp = new(order.BaseResponse)
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
