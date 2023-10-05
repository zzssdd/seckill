package product

import (
	"log"
	product "seckill/kitex_gen/product/product"
)

func main() {
	svr := product.NewServer(new(ProductImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
