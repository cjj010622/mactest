package logic

import (
	"context"
	"errors"
	"zero-study/mall/login/api/internal/svc"
	"zero-study/mall/login/api/internal/types"
	"zero-study/mall/regist/rpc/regist"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLoginLogic {
	return &GetLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLoginLogic) GetLogin(req *types.LoginReq) (resp *types.LoginReply, err error) {
	regist, err := l.svcCtx.RegistRpc.GetRegist(l.ctx, &regist.IdRequest{Id: "cjj"})
	if err != nil {
		return nil, err
	}
	if regist.Account != "yourdad" {
		return nil, errors.New("用户不存在")
	}
	account := "yourdad"
	return &types.LoginReply{
		Id:      req.Id,
		Account: account,
		Pwd:     "bbbb",
	}, nil

}
