namespace go user

struct BaseRequest{
    1:string email
    2:string password
}

struct BaseResponse{
    1:i16 code;
    2:string msg;
}

struct LoginResponse{
    1:i16 code;
    2:string msg;
    3:string token;
}

service User{
    BaseResponse Registry(1:BaseRequest req)
    LoginResponse Login(1:BaseRequest req)
}