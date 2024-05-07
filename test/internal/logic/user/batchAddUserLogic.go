package user

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"test/common/xerr"
	"test/model"

	"test/internal/svc"
	"test/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BatchAddUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 批量新增用户
func NewBatchAddUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BatchAddUserLogic {
	return &BatchAddUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BatchAddUserLogic) BatchAddUser(req *types.BatchAddUserRequest) (resp *types.BatchAddUserResponse, err error) {
	// 查询手机号是否都存在过，如果不存在则可以进行插入
	mobileList := make([]string, 0)
	mobileSet := make(map[string]string)

	for _, v := range req.Data {
		mobileList = append(mobileList, v.Mobile)
		mobileSet[v.Mobile] = ""
	}
	//fmt.Println(len(mobileSet), mobileSet)
	//fmt.Println(len(mobileList), mobileList)
	if len(mobileList) != len(mobileSet) {
		return nil, xerr.NewErrCode(xerr.USER_BATCH_ADD_HAS_SAME_MOBILE_ERROR)
	}

	selectBuilder := l.svcCtx.UserModel.SelectBuilder().Where(squirrel.Eq{"`mobile`": mobileList})
	count, err := l.svcCtx.UserModel.FindCountByBuilderNoCache(l.ctx, selectBuilder, "id")
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "BatchAddUser FindCountByBuilderNoCache err: %s", err)
	}
	if count > 0 {
		return nil, xerr.NewErrCode(xerr.USER_BATCH_ADD_HAS_MOBILE_ALREADY_EXISTS_ERROR)
	}
	// 开启事务插入
	err = l.svcCtx.UserModel.Trans(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		//insertBuilder := l.svcCtx.UserModel.InsertBuilder()
		//for _, data := range req.Data {
		//	// data.Mobile, data.Password, data.Nickname, data.Sex, data.Avatar, data.Info
		//	defaultPassword := "123456"
		//	insertBuilder = insertBuilder.Values(data.Mobile, defaultPassword, data.Nickname, data.Sex, data.Avatar, data.Info)
		//}
		//
		//insertResult, err2 := l.svcCtx.UserModel.InsertByBuilderNoCache(ctx, session, insertBuilder)
		//if err2 != nil {
		//	return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "BatchAddUser InsertByBuilderNoCache err: %s", err)
		//}
		insertDatas := make([]*model.User, 0)
		for _, data := range req.Data {
			var _user model.User
			err = copier.Copy(&_user, data)
			if err != nil {
				return err
			}
			// 设置默认密码
			_user.Password = "123456"
			insertDatas = append(insertDatas, &_user)
		}
		insertResult, err := l.svcCtx.UserModel.InsertMany(ctx, session, insertDatas)
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "BatchAddUser InsertByBuilderWithCache err: %s", err)
		}

		affected, err2 := insertResult.RowsAffected()
		if err2 != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "BatchAddUser RowsAffected err: %s", err)
		}
		if affected != int64(len(req.Data)) {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "BatchAddUser RowsAffected %d, data rows %d", affected, len(req.Data))
		}

		return nil
	})
	if err != nil {
		return nil, err
	}
	// 根据唯一手机号查询出来
	datas, err := l.svcCtx.UserModel.FindManyByBuilderNoCache(l.ctx, selectBuilder, "id")
	if err != nil && err != sqlx.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "BatchAddUser after FindManyByBuilderNoCache err: %s", err)
	}
	//fmt.Println(datas)
	//fmt.Println(len(datas))
	data := make([]types.User, 0)
	if len(datas) >= 0 {
		for _, v := range datas {
			//fmt.Println(v.Nickname)
			var _user types.User
			_ = copier.Copy(&_user, v)
			data = append(data, _user)
			// 添加缓存
		}
	}

	return &types.BatchAddUserResponse{
		BaseMsgResp: types.BaseMsgResp{
			Code: xerr.OK,
			Msg:  "success",
		},
		Data: data,
	}, nil
}
