namespace go api

include "order.thrift"
include "product.thrift"
include "seckill.thrift"
include "user.thrift"

service Api{
    user.BaseResponse Registry(1:user.BaseRequest req)(api.post="/registry")
    user.LoginResponse Login(1:user.BaseRequest req)(api.post="/login")

    product.BaseResponse ProductAdd(1:product.ProductInfo req)(api.post="/product")
    product.ListResponse ProductList(1:product.ListRequest req)(api.get="/list")


    seckill.SeckillResponse DoSeckill(1:seckill.SeckillRequest req)(api.post="/seckill")
    seckill.BaseResponse Submit(1:seckill.SubmitRequest req)(api.post="/submit")
}