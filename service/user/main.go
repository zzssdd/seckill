package user

import (
	"log"
	user "seckill/kitex_gen/user/user"
)

func main() {
	svr := user.NewServer(new(UserImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
