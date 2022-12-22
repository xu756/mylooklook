package logic

import (
	"context"
	"log"
	"mylooklook/api/user/internal/svc"
	"mylooklook/api/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.UserLogin) (resp *types.Res, err error) {
	user, er := l.svcCtx.UserModel.FindOne(l.ctx, 5)
	if er != nil {
		log.Print(er)
		return nil, er
	}
	resp = &types.Res{
		Code: 200,
		Msg:  "success",
		Data: user,
	}
	return resp, nil
}
