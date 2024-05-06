package user

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"test/common/result"
	"test/internal/logic/user"
	"test/internal/svc"
	"test/internal/types"
)

// 批量删除用户
func BatchDeleteUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.BatchDeleteUserRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewBatchDeleteUserLogic(r.Context(), svcCtx)
		resp, err := l.BatchDeleteUser(&req)
		result.HttpResult(r, w, resp, err)
	}
}
