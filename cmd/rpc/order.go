package rpc

import (
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	etcd "github.com/kitex-contrib/registry-etcd"
	"seckill/conf"
	"seckill/kitex_gen/order/order"
	"seckill/kitex_gen/seckill/seckill"
	. "seckill/pkg/log"
)

var OrderClient order.Client

func newOrderClient() {
	var err error
	r, err := etcd.NewEtcdResolver([]string{conf.EtcdDSN})
	if err != nil {
		Log.Fatalln("get etcd registry err:", err)
	}
	SeckillClient, err = seckill.NewClient(
		conf.OrderServiceName,
		client.WithResolver(r),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.OrderServiceName}),
	)
	if err != nil {
		Log.Fatalln("get order client err:", err)
	}
}

func init() {
	if ProductClient == nil {
		newOrderClient()
	}
}
