package order

import (
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"net"
	"seckill/conf"
	"seckill/dao/cache"
	"seckill/dao/db"
	"seckill/kitex_gen/order/order"
	. "seckill/pkg/log"
)

func main() {
	r, err := etcd.NewEtcdRegistry([]string{conf.EtcdDSN})
	if err != nil {
		Log.Fatalf("registry into etcd err:", err)
		return
	}
	addr, err := net.ResolveIPAddr("tcp", conf.OrderTcpAddr)
	if err != nil {
		Log.Fatalf("get order tcp addr err:", err)
	}
	orderServer := &OrderImpl{
		dao:   db.NewDao(),
		cache: cache.NewCache(),
	}
	svr := order.NewServer(orderServer,
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.OrderServiceName}),
	)
	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
