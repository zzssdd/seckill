package order

import (
	"log"
	order "seckill/kitex_gen/order/order"
)

func main() {
	svr := order.NewServer(new(OrderImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
