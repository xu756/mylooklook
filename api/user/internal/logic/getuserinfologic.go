package logic

import (
	"context"
	"log"

	"mylooklook/api/user/internal/svc"
	"mylooklook/api/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo() (resp *types.Res, err error) {
	log.Print("GetUserInfo")
	return &types.Res{
		Code: 200,
		Msg:  "success",
		Data: "hello",
	}, nil
}
