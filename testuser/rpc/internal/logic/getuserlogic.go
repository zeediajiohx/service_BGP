package logic

import (
	"context"

	"bgp-zero/testuser/rpc/internal/svc"
	"bgp-zero/testuser/rpc/types/testuser"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *testuser.IdRequest) (*testuser.UserResponse, error) {
	// todo: add your logic here and delete this line

	return &testuser.UserResponse{
		Id:     "my",
		Name:   "name",
		Gender: true,
	}, nil
}
