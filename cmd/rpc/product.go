package rpc

import (
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	etcd "github.com/kitex-contrib/registry-etcd"
	"seckill/conf"
	"seckill/kitex_gen/product/product"
	"seckill/kitex_gen/user/user"
	. "seckill/pkg/log"
)

var ProductClient product.Client

func newProductClient() {
	var err error
	r, err := etcd.NewEtcdResolver([]string{conf.EtcdDSN})
	if err != nil {
		Log.Fatalln("get etcd registry err:", err)
	}
	UserClient, err = user.NewClient(
		conf.ProductServiceName,
		client.WithResolver(r),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.ProductServiceName}),
	)
	if err != nil {
		Log.Fatalln("get product client err:", err)
	}
}

func init() {
	if ProductClient == nil {
		newProductClient()
	}
}
