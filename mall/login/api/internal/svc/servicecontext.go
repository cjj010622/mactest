package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"zero-study/mall/login/api/internal/config"
	"zero-study/mall/regist/rpc/regist"
)

type ServiceContext struct {
	Config    config.Config
	RegistRpc regist.Regist
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		RegistRpc: regist.NewRegist(zrpc.MustNewClient(c.RegistRpc)),
	}
}
