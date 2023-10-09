package seckill

import (
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"net"
	"seckill/conf"
	"seckill/dao/cache"
	"seckill/dao/db"
	"seckill/kitex_gen/seckill/seckill"
	. "seckill/pkg/log"
)

func main() {
	r, err := etcd.NewEtcdRegistry([]string{conf.EtcdDSN})
	if err != nil {
		Log.Fatalf("registry into etcd err:", err)
		return
	}
	addr, err := net.ResolveIPAddr("tcp", conf.ProductTcpAddr)
	if err != nil {
		Log.Fatalf("get product tcp addr err:", err)
	}
	seckillServer := &SeckillImpl{
		dao:   db.NewDao(),
		cache: cache.NewCache(),
	}
	svr := seckill.NewServer(seckillServer,
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.SeckillServiceName}),
	)
	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
