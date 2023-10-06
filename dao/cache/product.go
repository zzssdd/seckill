package cache

import (
	"context"
	"strconv"
)

const PreffixTag = "product:"

func getPreffix(id int) string {
	return PreffixTag + strconv.Itoa(id)
}

type Product struct {
}

func (p *Product) StoreProductInfo(ctx context.Context, id int, name string, price float32, pic string, des string, num int) error {
	data := map[string]interface{}{
		"name":  name,
		"price": price,
		"pic":   pic,
		"des":   des,
		"num":   num,
	}
	return rds.HMSet(ctx, getPreffix(id), data).Err()
}

func (p *Product) ExistProduct(ctx context.Context, id int) bool {
	return rds.Exists(ctx, getPreffix(id)).Val() == 1
}

func (p *Product) GetProductInfo(ctx context.Context, id int) (name string, price float64, pic string, des string, num int) {
	val := rds.HGetAll(ctx, getPreffix(id)).Val()
	price, _ = strconv.ParseFloat(val["price"], 32)
	num, _ = strconv.Atoi(val["num"])
	name, pic, des = val["name"], val["pic"], val["des"]
	return
}

func (p *Product) PreSubNum(ctx context.Context, id int, num int) bool {
	script := `if redis.call('hget',KEYS[1],num)< ARGV[1] then return 0 elseif redis.call('hincrby',KEYS[1],num,-ARGV[2]) return 1`
	return rds.Eval(ctx, script, []string{getPreffix(id)}, num).Val().(int) == 1
}

func (p *Product) AddNum(ctx context.Context, id int, num int) error {
	return rds.HIncrBy(ctx, getPreffix(id), "num", int64(num)).Err()
}
