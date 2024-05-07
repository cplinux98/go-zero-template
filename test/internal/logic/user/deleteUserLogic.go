package user

import (
	"context"
	"github.com/pkg/errors"
	"test/common/xerr"
	"test/model"

	"test/internal/svc"
	"test/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除用户
func NewDeleteUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserLogic {
	return &DeleteUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteUserLogic) DeleteUser(req *types.DeleteUserRequest) (resp *types.BaseMsgResp, err error) {
	result, err := l.svcCtx.UserModel.Delete(l.ctx, nil, req.Id)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "DeleteUser Delete err %v", err)
	}
	if err == model.ErrNotFound {
		return nil, xerr.NewErrCode(xerr.USER_ID_NOT_EXISTS_ERROR)
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "DeleteUser RowsAffected err %v", err)
	}
	if affected == 0 {
		return nil, xerr.NewErrCode(xerr.DB_DELETE_AFFECTED_ZERO_ERROR)
	}

	return &types.BaseMsgResp{
		Code: xerr.OK,
		Msg:  "success",
	}, nil
}
