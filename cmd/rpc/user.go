package rpc

import (
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	etcd "github.com/kitex-contrib/registry-etcd"
	"seckill/conf"
	"seckill/kitex_gen/user/user"
	. "seckill/pkg/log"
)

var UserClient user.Client

func newUserClient() {
	var err error
	r, err := etcd.NewEtcdResolver([]string{conf.EtcdDSN})
	if err != nil {
		Log.Fatalln("get etcd registry err:", err)
	}
	UserClient, err = user.NewClient(
		conf.UserServiceName,
		client.WithResolver(r),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.UserServiceName}),
	)
	if err != nil {
		Log.Fatalln("get user client err:", err)
	}
}

func init() {
	if UserClient == nil {
		newUserClient()
	}
}
