package user

import (
	"context"
	user "seckill/kitex_gen/user"
)

// UserImpl implements the last service interface defined in the IDL.
type UserImpl struct{}

// Registry implements the UserImpl interface.
func (s *UserImpl) Registry(ctx context.Context, req *user.BaseRequest) (resp *user.BaseResponse, err error) {
	// TODO: Your code here...
	resp = new(user.BaseResponse)

	return
}

// Login implements the UserImpl interface.
func (s *UserImpl) Login(ctx context.Context, req *user.BaseRequest) (resp *user.LoginResponse, err error) {
	// TODO: Your code here...
	return
}
