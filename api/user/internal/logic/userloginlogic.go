package logic

import (
	"context"
	"errors"
	"mylooklook/api/user/internal/svc"
	"mylooklook/api/user/internal/types"
	"strings"

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
	if len(strings.TrimSpace(req.Username)) == 0 || len(strings.TrimSpace(req.Password)) == 0 {
		return nil, errors.New("参数错误")
	}
	user, er := l.svcCtx.UserModel.FindUserByName(l.ctx, req.Username)
	if er != nil {
		return nil, errors.New("用户不存在")
	}
	resp = &types.Res{
		Code: 200,
		Msg:  "success",
		Data: user,
	}
	return resp, nil
}
