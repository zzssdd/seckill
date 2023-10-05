package conf

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var conf *viper.Viper

func getAllField() {
	if err := conf.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic("找不到配置文件")
		} else {
			panic(err)
		}
	}
	ApiServiceName = conf.GetString("ApiServiceName")
	UserServiceName = conf.GetString("UserServiceName")
	ProductServiceName = conf.GetString("ProductServiceName")
	SeckillServiceName = conf.GetString("SeckillServiceName")
	OrderServiceName = conf.GetString("OrderServiceName")
	ApiTcpAddr = conf.GetString("ApiTcpAddr")
	UserTcpAddr = conf.GetString("UserTcpAddr")
	ProductTcpAddr = conf.GetString("ProductTcpAddr")
	SeckillTcpAddr = conf.GetString("SeckillTcpAddr")
	OrderTcpAddr = conf.GetString("OrderTcpAddr")
	MysqlUserDSN = conf.GetString("MysqlUserDSN")
	MysqlProductDSN = conf.GetString("MysqlProductDSN")
	MysqlOrderStatusDSN = conf.GetString("MysqlOrderStatusDSN")
	RedisDSN = conf.GetString("RedisDSN")
	EtcdDSN = conf.GetString("EtcdDSN")
	RabbitmqDSN = conf.GetString("RabbitmqDSN")
	IpLimitCount = conf.GetInt("IpLimitCount")
	Salt = conf.GetString("Salt")
	TokenExpireTime = conf.GetInt("TokenExpireTime")
}

func LoadConf() {
	conf = viper.New()
	conf.AddConfigPath("./conf")
	conf.SetConfigName("base")
	conf.SetConfigType("json")
	getAllField()
	conf.WatchConfig()
	conf.OnConfigChange(func(in fsnotify.Event) {
		getAllField()
	})
}
