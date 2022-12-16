package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"mylooklook/api/user/internal/logic"
	"mylooklook/api/user/internal/svc"
	"mylooklook/api/user/internal/types"
)

func userLoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserLogin
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUserLoginLogic(r.Context(), svcCtx)
		resp, err := l.UserLogin(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
