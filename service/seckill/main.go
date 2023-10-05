package seckill

import (
	"log"
	seckill "seckill/kitex_gen/seckill/seckill"
)

func main() {
	svr := seckill.NewServer(new(SeckillImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
