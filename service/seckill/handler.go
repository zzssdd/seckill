package seckill

import (
	"context"
	seckill "seckill/kitex_gen/seckill"
)

// SeckillImpl implements the last service interface defined in the IDL.
type SeckillImpl struct{}

// DoSeckill implements the SeckillImpl interface.
func (s *SeckillImpl) DoSeckill(ctx context.Context, req *seckill.SeckillRequest) (resp *seckill.SeckillResponse, err error) {
	// TODO: Your code here...
	return
}

// Submit implements the SeckillImpl interface.
func (s *SeckillImpl) Submit(ctx context.Context, req *seckill.SubmitRequest) (resp *seckill.BaseResponse, err error) {
	// TODO: Your code here...
	return
}
