package user

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"test/common/xerr"

	"test/internal/svc"
	"test/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取用户详情
func NewListUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListUserLogic {
	return &ListUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListUserLogic) ListUser(req *types.ListUserRequest) (resp *types.ListUserResponse, err error) {
	// 组装查询条件
	whereBuilder := l.svcCtx.UserModel.SelectBuilder()
	if req.Keyword != "" {
		whereBuilder = whereBuilder.Where(squirrel.Eq{"nickname": req.Keyword})
	}

	// 执行查询
	list, total, err := l.svcCtx.UserModel.FindPageListByPageWithTotal(l.ctx, whereBuilder, req.Page, req.PageSize, "id DESC")
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "list user failed, err: %s;req:%+v", err, req)
	}
	var respList []types.User

	if len(list) > 0 {
		for _, v := range list {
			var user types.User
			_ = copier.Copy(&user, v)
			respList = append(respList, user)
		}
	}

	return &types.ListUserResponse{
		BaseMsgResp: types.BaseMsgResp{
			Code: xerr.OK,
			Msg:  "success",
		},
		Data: types.ListUserInfo{
			BaseListInfo: types.BaseListInfo{
				Total: total,
			},
			List: respList,
		},
	}, nil
}
