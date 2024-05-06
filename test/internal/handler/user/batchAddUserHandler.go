package user

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"test/common/result"
	"test/internal/logic/user"
	"test/internal/svc"
	"test/internal/types"
)

// 批量新增用户
func BatchAddUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.BatchAddUserRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewBatchAddUserLogic(r.Context(), svcCtx)
		resp, err := l.BatchAddUser(&req)
		result.HttpResult(r, w, resp, err)
	}
}
