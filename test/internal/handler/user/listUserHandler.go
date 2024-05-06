package user

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"test/common/result"
	"test/internal/logic/user"
	"test/internal/svc"
	"test/internal/types"
)

// 获取用户详情
func ListUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListUserRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewListUserLogic(r.Context(), svcCtx)
		resp, err := l.ListUser(&req)
		result.HttpResult(r, w, resp, err)
	}
}
