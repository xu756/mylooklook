package logic

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"mylooklook/api/user/internal/svc"
	"mylooklook/api/user/internal/types"
	"strconv"
	"strings"
	"time"

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
func (l *UserLoginLogic) getJwtToken(secretKey string, iat, seconds, username string) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["user"] = username
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}

func (l *UserLoginLogic) UserLogin(req *types.UserLogin) (resp *types.Res, err error) {
	if len(strings.TrimSpace(req.Username)) == 0 || len(strings.TrimSpace(req.Password)) == 0 {
		return nil, errors.New("参数错误")
	}
	user, er := l.svcCtx.UserModel.FindUserByName(l.ctx, req.Username)
	log.Print(user)
	if er != nil {
		return nil, errors.New("用户不存在")
	}
	now := time.Now().Unix()
	//accessExpire := l.svcCtx.Config.Auth.AccessExpire
	jwtToken, e := l.getJwtToken(l.svcCtx.Config.Auth.AccessSecret, strconv.FormatInt(now, 10), strconv.FormatInt(l.svcCtx.Config.Auth.AccessExpire, 10), req.Username)
	if e != nil {
		return nil, errors.New("没有用户")
	}
	resp = &types.Res{
		Code: 200,
		Msg:  "success",
		Data: jwtToken,
	}
	return resp, nil
}
