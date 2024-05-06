package user

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"test/common/result"
	"test/internal/logic/user"
	"test/internal/svc"
	"test/internal/types"
)

// 修改用户信息
func UpdateUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateUserRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewUpdateUserLogic(r.Context(), svcCtx)
		resp, err := l.UpdateUser(&req)
		result.HttpResult(r, w, resp, err)
	}
}
