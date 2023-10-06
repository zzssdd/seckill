package user

import (
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"net"
	"seckill/conf"
	"seckill/dao/cache"
	"seckill/dao/db"
	"seckill/kitex_gen/user/user"
	. "seckill/pkg/log"
)

func main() {
	r, err := etcd.NewEtcdRegistry([]string{conf.EtcdDSN})
	if err != nil {
		Log.Fatalf("registry into etcd err:", err)
		return
	}
	addr, err := net.ResolveIPAddr("tcp", conf.UserTcpAddr)
	if err != nil {
		Log.Fatalf("get user tcp addr err:", err)
	}
	userServer := &UserImpl{
		dao:   db.NewDao(),
		cache: cache.NewCache(),
	}
	svr := user.NewServer(userServer,
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.UserServiceName}),
	)
	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
