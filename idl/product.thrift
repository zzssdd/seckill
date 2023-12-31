namespace go product

struct ProductInfo{
    1:i32 id;
    2:string name;
    3:double price;
    4:string pic;
    5:string des;
    6:i32 num;
}

struct BaseResponse{
    1:i16 code;
    2:string msg;
}

struct BuyRequest{
    1:i32 id;
    2:i32 num;
}

struct ListRequest{
    1:i32 offset;
    2:i32 limit;
}

struct ListResponse{
    1:list<ProductInfo> products
    2:i16 code;
    3:string msg;
}

struct IdRequest{
    1:i32 id
}

struct ProductResponse{
        1:ProductInfo product
        2:i16 code;
        3:string msg;
}

service Product{
    BaseResponse AddProduct(ProductInfo req)
    ListResponse ListProduct(ListRequest req)
    ProductResponse ProductInfo(IdRequest req)
    BaseResponse Try(BuyRequest req)
    BaseResponse Commit(BuyRequest req)
    BaseResponse Cancel(BuyRequest req)
}