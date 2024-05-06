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

type AddUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 新增用户
func NewAddUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserLogic {
	return &AddUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddUserLogic) AddUser(req *types.AddUserRequest) (resp *types.AddUserResponse, err error) {
	// 1.检查用户是否存在
	userRecord, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, req.Mobile)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "AddUser FindOneByMobile err %v", err)
	}
	if userRecord != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.USER_EXISTS_ERROR), "AddUser mobile: %s exists", req.Mobile)
	}
	// 2.进行添加
	user := new(model.User)
	user.Mobile = req.Mobile
	user.Nickname = req.Nickname
	user.Avatar = req.Avatar
	user.Sex = req.Sex

	insert, err := l.svcCtx.UserModel.Insert(l.ctx, nil, user)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "AddUser Insert err: %v, user: %+v", err, user)
	}
	id, err := insert.LastInsertId()
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "AddUser LastInsertId err: %v, user: %+v", err, user)
	}
	// 3.进行查询数据库数据
	one, err := l.svcCtx.UserModel.FindOne(l.ctx, id)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "AddUser FindOne err: %v, id: %+v", err, id)
	}
	var respUser types.User

	_ = copier.Copy(&respUser, one)

	return &types.AddUserResponse{
		BaseMsgResp: types.BaseMsgResp{
			Code: xerr.OK,
			Msg:  "success",
		},
		Data: respUser,
	}, nil

}
