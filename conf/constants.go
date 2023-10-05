package conf

var (
	ApiServiceName      string
	UserServiceName     string
	ProductServiceName  string
	SeckillServiceName  string
	OrderServiceName    string
	ApiTcpAddr          string
	UserTcpAddr         string
	ProductTcpAddr      string
	SeckillTcpAddr      string
	OrderTcpAddr        string
	MysqlUserDSN        string
	MysqlProductDSN     string
	MysqlOrderStatusDSN string
	RedisDSN            string
	EtcdDSN             string
	RabbitmqDSN         string
	IpLimitCount        int
	Salt                string
	TokenExpireTime     int
)
