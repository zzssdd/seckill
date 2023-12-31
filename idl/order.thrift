namespace go order

struct OrderInfo{
    1:i64 id;
    2:i64 uid;
    3:i32 pid;
    4:i64 timeStamp
}

struct BaseResponse{
    1:i16 code;
    2:string msg;
}

struct OrderCancelRequest{

}

struct OrderCancelResponse{

}

service Order{
   BaseResponse OrderAdd(1:OrderInfo req)
   BaseResponse OrderStatusAdd(1:OrderInfo req)
   OrderCancelResponse OrderCancel(1:OrderCancelRequest req)
   BaseResponse Try(1:OrderInfo req)
   BaseResponse Commit(1:OrderInfo req)
   BaseResponse Cancal(1:OrderInfo req)
}