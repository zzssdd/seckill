package user

import (
	"context"
	"seckill/dao/cache"
	"seckill/dao/db"
	user "seckill/kitex_gen/user"
	"seckill/pkg/errmsg"
	. "seckill/pkg/log"
	"seckill/utils"
)

// UserImpl implements the last service interface defined in the IDL.
type UserImpl struct {
	dao   db.Dao
	cache cache.Cache
}

// Registry implements the UserImpl interface.
func (s *UserImpl) Registry(ctx context.Context, req *user.BaseRequest) (resp *user.BaseResponse, err error) {
	// TODO: Your code here...
	resp = new(user.BaseResponse)
	add, err := s.dao.User.UserAdd(req.Email, req.Password)
	if err != nil {
		Log.Error(err)
		return nil, err
	}
	if add == 0 {
		resp.Code = errmsg.UserExist
		resp.Msg = errmsg.GetMsg(errmsg.UserExist)
		return
	}
	resp.Code = errmsg.Success
	resp.Msg = errmsg.GetMsg(errmsg.Success)
	return
}

// Login implements the UserImpl interface.
func (s *UserImpl) Login(ctx context.Context, req *user.BaseRequest) (resp *user.LoginResponse, err error) {
	// TODO: Your code here...
	resp = new(user.LoginResponse)
	var is_ok bool
	var id int64
	if s.cache.User.ExistUserInfo(ctx, req.Email) {
		is_ok = s.cache.User.CheckLogin(ctx, req.Email, req.Password)
	} else {
		id, is_ok = s.dao.User.UserLogin(req.Email, req.Password)
		err = s.cache.User.StoreUserInfo(ctx, req.Email, req.Password, id)
		if err != nil {
			Log.Errorf("store user info into cache err", err)
		}
	}
	if !is_ok {
		resp.Code = errmsg.LoginFailed
		resp.Msg = errmsg.GetMsg(errmsg.LoginFailed)
		return
	}
	token, err := utils.GenToken(id, req.Email)
	if err != nil {
		Log.Error(err)
		resp.Code = errmsg.Error
		resp.Msg = errmsg.GetMsg(errmsg.Error)
		return
	}
	resp.Code = errmsg.Success
	resp.Msg = errmsg.GetMsg(errmsg.Success)
	resp.Token = token
	return
}
