package user

import (
	"context"
	"database/sql"
	"github.com/Masterminds/squirrel"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"test/common/xerr"
	"test/model"

	"test/internal/svc"
	"test/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 修改用户信息
func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserLogic) UpdateUser(req *types.UpdateUserRequest) (resp *types.UpdateUserResponse, err error) {
	// 1.检查id是否存在
	// update里面会检查的

	// 2.检查更新的手机号是否已经存在了
	whereBuilder := l.svcCtx.UserModel.SelectBuilder().Where(squirrel.Eq{"mobile": req.Mobile}).Where(squirrel.NotEq{"id": req.Id})
	count, err := l.svcCtx.UserModel.FindCountByBuilderNoCache(l.ctx, whereBuilder, "id")
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "UpdateUser FindCountByBuilderNoCache err %v", err)
	}
	if count > 0 {
		return nil, xerr.NewErrCode(xerr.USER_MOBILE_ALREADY_EXISTS_ERROR)
	}
	// 3.如果检查没问题，则进行更新
	var newUser model.User
	err = copier.Copy(&newUser, req)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.SERVER_COMMON_ERROR), "copier.Copy err %v", err)
	}
	newUser.Id = req.Id

	update, err := l.svcCtx.UserModel.Update(l.ctx, nil, &newUser)
	if err != nil && err != sql.ErrNoRows {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "UpdateUser Update err %v", err)
	}
	if err == sql.ErrNoRows {
		return nil, xerr.NewErrCode(xerr.USER_ID_NOT_EXISTS_ERROR)
	}
	affected, err := update.RowsAffected()
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "UpdateUser RowsAffected err %v", err)
	}
	if affected == 0 {
		return nil, xerr.NewErrCode(xerr.DB_UPDATE_AFFECTED_ZERO_ERROR)
	}

	// 4.把更新的结果查询回来
	one, err := l.svcCtx.UserModel.FindOne(l.ctx, newUser.Id)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "UpdateUser After Find err %v", err)
	}
	// 5.copy到响应里面
	var respUser types.User
	err = copier.Copy(&respUser, one)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.SERVER_COMMON_ERROR), "copier.Copy err %v", err)
	}

	return &types.UpdateUserResponse{
		BaseMsgResp: types.BaseMsgResp{
			Code: xerr.OK,
			Msg:  "success",
		},
		Data: respUser,
	}, nil
}
