package logic

import (
	"context"

	"zero-study/mall/regist/rpc/internal/svc"
	"zero-study/mall/regist/rpc/types/regist"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRegistLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRegistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRegistLogic {
	return &GetRegistLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetRegistLogic) GetRegist(in *regist.IdRequest) (*regist.RegistRepose, error) {
	//testid := "cjj"
	//testacc := "yourdad"
	//testpwd := "123456"
	return &regist.RegistRepose{
		Id:      "cjj",
		Account: "yourdad",
		Pwd:     "123456",
	}, nil
}
