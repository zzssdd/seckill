package errmsg

const (
	Success = 200

	UserExist   = 300
	LoginFailed = 301
	NotLogin    = 302
	TokenError  = 303
	LoginAgain  = 304

	NumNotEnough = 400
	OrderExist   = 401
	Sellout      = 402

	Error = 500
)

var codeMsg = map[int]string{
	Success: "成功！",

	UserExist:   "用户已被注册",
	LoginFailed: "邮箱或密码错误",
	NotLogin:    "请先登陆",
	TokenError:  "Token格式不正确",
	LoginAgain:  "请重新登陆",

	NumNotEnough: "库存量不足",
	OrderExist:   "请勿重复提交订单",
	Sellout:      "商品已被抢空",

	Error: "失败",
}

func GetMsg(code int) string {
	return codeMsg[code]
}
