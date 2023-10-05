// Code generated by Kitex v0.7.2. DO NOT EDIT.

package order

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	order "seckill/kitex_gen/order"
)

func serviceInfo() *kitex.ServiceInfo {
	return orderServiceInfo
}

var orderServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "Order"
	handlerType := (*order.Order)(nil)
	methods := map[string]kitex.MethodInfo{
		"OrderAdd":       kitex.NewMethodInfo(orderAddHandler, newOrderOrderAddArgs, newOrderOrderAddResult, false),
		"OrderStatusAdd": kitex.NewMethodInfo(orderStatusAddHandler, newOrderOrderStatusAddArgs, newOrderOrderStatusAddResult, false),
		"Try":            kitex.NewMethodInfo(tryHandler, newOrderTryArgs, newOrderTryResult, false),
		"Commit":         kitex.NewMethodInfo(commitHandler, newOrderCommitArgs, newOrderCommitResult, false),
		"Cancal":         kitex.NewMethodInfo(cancalHandler, newOrderCancalArgs, newOrderCancalResult, false),
	}
	extra := map[string]interface{}{
		"PackageName":     "order",
		"ServiceFilePath": `idl/order.thrift`,
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.7.2",
		Extra:           extra,
	}
	return svcInfo
}

func orderAddHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*order.OrderOrderAddArgs)
	realResult := result.(*order.OrderOrderAddResult)
	success, err := handler.(order.Order).OrderAdd(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newOrderOrderAddArgs() interface{} {
	return order.NewOrderOrderAddArgs()
}

func newOrderOrderAddResult() interface{} {
	return order.NewOrderOrderAddResult()
}

func orderStatusAddHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*order.OrderOrderStatusAddArgs)
	realResult := result.(*order.OrderOrderStatusAddResult)
	success, err := handler.(order.Order).OrderStatusAdd(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newOrderOrderStatusAddArgs() interface{} {
	return order.NewOrderOrderStatusAddArgs()
}

func newOrderOrderStatusAddResult() interface{} {
	return order.NewOrderOrderStatusAddResult()
}

func tryHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*order.OrderTryArgs)
	realResult := result.(*order.OrderTryResult)
	success, err := handler.(order.Order).Try(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newOrderTryArgs() interface{} {
	return order.NewOrderTryArgs()
}

func newOrderTryResult() interface{} {
	return order.NewOrderTryResult()
}

func commitHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*order.OrderCommitArgs)
	realResult := result.(*order.OrderCommitResult)
	success, err := handler.(order.Order).Commit(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newOrderCommitArgs() interface{} {
	return order.NewOrderCommitArgs()
}

func newOrderCommitResult() interface{} {
	return order.NewOrderCommitResult()
}

func cancalHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*order.OrderCancalArgs)
	realResult := result.(*order.OrderCancalResult)
	success, err := handler.(order.Order).Cancal(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newOrderCancalArgs() interface{} {
	return order.NewOrderCancalArgs()
}

func newOrderCancalResult() interface{} {
	return order.NewOrderCancalResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) OrderAdd(ctx context.Context, req *order.OrderInfo) (r *order.BaseResponse, err error) {
	var _args order.OrderOrderAddArgs
	_args.Req = req
	var _result order.OrderOrderAddResult
	if err = p.c.Call(ctx, "OrderAdd", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) OrderStatusAdd(ctx context.Context, req *order.OrderInfo) (r *order.BaseResponse, err error) {
	var _args order.OrderOrderStatusAddArgs
	_args.Req = req
	var _result order.OrderOrderStatusAddResult
	if err = p.c.Call(ctx, "OrderStatusAdd", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Try(ctx context.Context, req *order.OrderInfo) (r *order.BaseResponse, err error) {
	var _args order.OrderTryArgs
	_args.Req = req
	var _result order.OrderTryResult
	if err = p.c.Call(ctx, "Try", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Commit(ctx context.Context, req *order.OrderInfo) (r *order.BaseResponse, err error) {
	var _args order.OrderCommitArgs
	_args.Req = req
	var _result order.OrderCommitResult
	if err = p.c.Call(ctx, "Commit", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Cancal(ctx context.Context, req *order.OrderInfo) (r *order.BaseResponse, err error) {
	var _args order.OrderCancalArgs
	_args.Req = req
	var _result order.OrderCancalResult
	if err = p.c.Call(ctx, "Cancal", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}