package svc

import (
	"bgp-zero/testuser/rpc/usrt"
	"bgp-zero/video/api/internal/config"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc usrt.Usrt
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: usrt.NewUsrt(zrpc.MustNewClient(c.UserRpc)),
	}
}
