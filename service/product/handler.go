package main

import (
	"context"
	"seckill/dao/cache"
	"seckill/dao/db"
	"seckill/dao/db/model"
	product "seckill/kitex_gen/product"
	"seckill/pkg/errmsg"
	. "seckill/pkg/log"
)

// ProductImpl implements the last service interface defined in the IDL.
type ProductImpl struct {
	dao   db.Dao
	cache cache.Cache
}

// AddProduct implements the ProductImpl interface.
func (s *ProductImpl) AddProduct(ctx context.Context, req *product.ProductInfo) (resp *product.BaseResponse, err error) {
	// TODO: Your code here...
	resp = new(product.BaseResponse)
	err = s.dao.Product.ProductAdd(&model.ProductInfo{
		Name:  req.Name,
		Price: float32(req.Price),
		Pic:   req.Pic,
		Des:   req.Des,
		Num:   req.Num,
	})
	if err != nil {
		Log.Errorln("add product into mysql err:", err)
		resp.Code = errmsg.Error
		resp.Msg = errmsg.GetMsg(errmsg.Error)
		return
	}
	resp.Code = errmsg.Success
	resp.Msg = errmsg.GetMsg(errmsg.Success)
	return
}

// ListProduct implements the ProductImpl interface.
func (s *ProductImpl) ListProduct(ctx context.Context, req *product.ListRequest) (resp *product.ListResponse, err error) {
	// TODO: Your code here...
	resp = new(product.ListResponse)
	products, err := s.dao.Product.ProductList(int(req.Offset), int(req.Limit))
	if err != nil {
		Log.Errorln("get product list err:", err)
		resp.Code = errmsg.Error
		resp.Msg = errmsg.GetMsg(errmsg.Error)
		return
	}
	ret_products := []*product.ProductInfo{}
	for _, v := range products {
		s.cache.Prduct.StoreProductInfo(ctx, int(v.ID), v.Name, v.Price, v.Pic, v.Des, int(v.Num))
		ret_product := &product.ProductInfo{
			Id:    int32(v.ID),
			Name:  v.Name,
			Price: float64(v.Price),
			Pic:   v.Pic,
			Des:   v.Des,
			Num:   v.Num,
		}
		ret_products = append(ret_products, ret_product)
	}

	resp.Products = ret_products
	resp.Code = errmsg.Success
	resp.Msg = errmsg.GetMsg(errmsg.Success)
	return
}

// ProductInfo implements the ProductImpl interface.
func (s *ProductImpl) ProductInfo(ctx context.Context, req *product.IdRequest) (resp *product.ProductResponse, err error) {
	// TODO: Your code here...
	var productInfo *model.ProductInfo
	id := int(req.Id)
	if s.cache.Prduct.ExistProduct(ctx, id) {
		name, price, pic, des, num := s.cache.Prduct.GetProductInfo(ctx, id)
		productInfo.ID = id
		productInfo.Name = name
		productInfo.Price = float32(price)
		productInfo.Pic = pic
		productInfo.Des = des
		productInfo.Num = int32(num)
	} else {
		productInfo, err = s.dao.Product.ProductInfo(id)
		if err != nil {
			Log.Errorln("get product info from mysql err:", err)
			resp.Code = errmsg.Error
			resp.Msg = errmsg.GetMsg(errmsg.Error)
			return
		}
		err = s.cache.Prduct.StoreProductInfo(ctx, productInfo.ID, productInfo.Name, productInfo.Price, productInfo.Pic, productInfo.Des, int(productInfo.Num))
		if err != nil {
			Log.Errorln("store product info into cache err:", err)
			resp.Code = errmsg.Error
			resp.Msg = errmsg.GetMsg(errmsg.Error)
			return
		}
	}
	ret_product := &product.ProductInfo{
		Id:    req.Id,
		Name:  productInfo.Name,
		Price: float64(productInfo.Price),
		Pic:   productInfo.Pic,
		Des:   productInfo.Des,
		Num:   productInfo.Num,
	}
	resp.Code = errmsg.Success
	resp.Msg = errmsg.GetMsg(errmsg.Success)
	resp.Product = ret_product
	return
}

// Try implements the ProductImpl interface.
func (s *ProductImpl) Try(ctx context.Context, req *product.BuyRequest) (resp *product.BaseResponse, err error) {
	// TODO: Your code here...
	resp = new(product.BaseResponse)
	affect, err := s.dao.Product.Try(int(req.Id), int(req.Num))
	if err != nil || affect == 0 {
		Log.Errorln("try freeze product num err:", err)
		resp.Code = errmsg.NumNotEnough
		resp.Msg = errmsg.GetMsg(errmsg.NumNotEnough)
		return
	}
	resp.Code = errmsg.Success
	resp.Msg = errmsg.GetMsg(errmsg.Success)
	return
}

// Commit implements the ProductImpl interface.
func (s *ProductImpl) Commit(ctx context.Context, req *product.BuyRequest) (resp *product.BaseResponse, err error) {
	// TODO: Your code here...
	resp = new(product.BaseResponse)
	err = s.dao.Product.Commit(int(req.Id), int(req.Num))
	if err != nil {
		Log.Errorln("commit product stmt err:", err)
		resp.Code = errmsg.Error
		resp.Msg = errmsg.GetMsg(errmsg.Error)
		return
	}
	resp.Code = errmsg.Success
	resp.Msg = errmsg.GetMsg(errmsg.Success)
	return
}

// Cancel implements the ProductImpl interface.
func (s *ProductImpl) Cancel(ctx context.Context, req *product.BuyRequest) (resp *product.BaseResponse, err error) {
	// TODO: Your code here...
	resp = new(product.BaseResponse)
	err = s.dao.Product.Cancel(int(req.Id), int(req.Num))
	if err != nil {
		Log.Errorln("cancel product stmt err:", err)
		resp.Code = errmsg.Error
		resp.Msg = errmsg.GetMsg(errmsg.Error)
		return
	}
	resp.Code = errmsg.Success
	resp.Msg = errmsg.GetMsg(errmsg.Success)
	return
}
