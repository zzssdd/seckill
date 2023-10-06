package rpc

import (
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	etcd "github.com/kitex-contrib/registry-etcd"
	"seckill/conf"
	"seckill/kitex_gen/seckill/seckill"
	. "seckill/pkg/log"
)

var SeckillClient seckill.Client

func newSeckillClient() {
	var err error
	r, err := etcd.NewEtcdResolver([]string{conf.EtcdDSN})
	if err != nil {
		Log.Fatalln("get etcd registry err:", err)
	}
	SeckillClient, err = seckill.NewClient(
		conf.SeckillServiceName,
		client.WithResolver(r),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.SeckillServiceName}),
	)
	if err != nil {
		Log.Fatalln("get product client err:", err)
	}
}

func init() {
	if ProductClient == nil {
		newSeckillClient()
	}
}
