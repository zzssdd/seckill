package cache

import (
	"context"
	"strconv"
)

type Order struct {
}

const OrderPreffix = "order:"
const OrderStatusPreffix = "orderstatus:"

func getOrderTag(id int64) string {
	return OrderPreffix + strconv.FormatInt(id, 64)
}

func getOrderStatusTag(id int64) string {
	return OrderStatusPreffix + strconv.FormatInt(id, 64)
}

func (o *Order) StoreOrder(ctx context.Context, id int64, uid int64, pid int, timeStamp int64) error {
	data := map[string]interface{}{
		"pid":  pid,
		"uid":  uid,
		"time": timeStamp,
	}
	return rds.HMSet(ctx, getOrderTag(id), data).Err()
}

func (o *Order) ExistOrder(ctx context.Context, id int64) bool {
	return rds.Exists(ctx, getOrderTag(id)).Val() == 1
}

func (o *Order) StoreOrderStatus(ctx context.Context, id int64, uid int64, pid int) error {
	data := map[string]interface{}{
		"pid":    pid,
		"uid":    uid,
		"status": "Created",
	}
	return rds.HMSet(ctx, getOrderStatusTag(id), data).Err()
}

func (o *Order) ExistOrderStatus(ctx context.Context, id int64) bool {
	return rds.Exists(ctx, getOrderTag(id)).Val() == 1
}

func (o *Order) GetOrderStatus(ctx context.Context, id int64) (pid int, uid int64, status string) {
	val := rds.HGetAll(ctx, getOrderTag(id)).Val()
	pid, _ = strconv.Atoi(val["pid"])
	uid, _ = strconv.ParseInt(val["uid"], 10, 64)
	status = val["status"]
	return
}

func (o *Order) DoingOrderStatus(ctx context.Context, id int64, uid int64, pid int) error {
	data := map[string]interface{}{
		"pid":    pid,
		"uid":    uid,
		"status": "Doing",
	}
	return rds.HMSet(ctx, getOrderStatusTag(id), data).Err()
}

func (o *Order) DelOrderStatus(ctx context.Context, id int64) error {
	return rds.Del(ctx, getOrderTag(id)).Err()
}
