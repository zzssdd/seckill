namespace go seckill

struct SeckillRequest{
    1:i64 uid;
    2:i32 pid;
    3:i32 num;
}

struct SeckillResponse{
    1:i16 code;
    2:string msg;
    3:i64 id;
}

struct SubmitRequest{
    1:i64 id;
    2:i64 uid;
    3:i32 pid;
    4:i32 num;
}

struct BaseResponse{
    1:i16 code;
    2:string msg;
}

service Seckill{
    SeckillResponse DoSeckill(1:SeckillRequest req)
    BaseResponse Submit(1:SubmitRequest req)
}