package user

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"test/common/xerr"
	"test/model"

	"test/internal/svc"
	"test/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取用户详情
func NewDetailUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailUserLogic {
	return &DetailUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailUserLogic) DetailUser(req *types.DetailUserRequest) (resp *types.DetailUserResponse, err error) {
	// 1.使用id查询是否存在
	one, err := l.svcCtx.UserModel.FindOne(l.ctx, req.Id)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "DetailUser FindOne err %v", err)
	}
	if one == nil {
		return nil, xerr.NewErrCode(xerr.USER_ID_NOT_EXISTS_ERROR)
	}
	// 2.如果存在则copy
	var respUser types.User
	_ = copier.Copy(&respUser, one)

	return &types.DetailUserResponse{
		BaseMsgResp: types.BaseMsgResp{
			Code: xerr.OK,
			Msg:  "success",
		},
		Data: respUser,
	}, nil
}
