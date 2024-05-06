package user

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"test/common/xerr"

	"test/internal/svc"
	"test/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BatchDeleteUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 批量删除用户
func NewBatchDeleteUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BatchDeleteUserLogic {
	return &BatchDeleteUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BatchDeleteUserLogic) BatchDeleteUser(req *types.BatchDeleteUserRequest) (resp *types.BaseMsgResp, err error) {
	// 查询id是否都存在
	whereBuilder := l.svcCtx.UserModel.SelectBuilder().Where(squirrel.Eq{"id": req.Ids})
	count, err := l.svcCtx.UserModel.FindCountByBuilderNoCache(l.ctx, whereBuilder, "id")
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "BatchDeleteUser FindCountByBuilderNoCache err %v", err)
	}
	if count != uint64(len(req.Ids)) {
		return nil, xerr.NewErrCode(xerr.USER_BATCH_DELETE_HAS_NOT_EXISTS_ERROR)
	}
	// 使用事务进行删除
	err = l.svcCtx.UserModel.Trans(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		deleteBuilder := l.svcCtx.UserModel.DeleteBuilder().Where(squirrel.Eq{"id": req.Ids})
		cache, err2 := l.svcCtx.UserModel.DeleteByBuilderNoCache(ctx, session, deleteBuilder)
		if err2 != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "BatchDeleteUser DeleteByBuilderNoCache err %v", err)
		}
		affected, err2 := cache.RowsAffected()
		if err2 != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "BatchDeleteUser RowsAffected err %v", err)
		}
		if affected != int64(len(req.Ids)) {
			return xerr.NewErrCode(xerr.USER_BATCH_DELETE_AFFECTED_ZERO_ERROR)
		}

		return nil
	})

	return &types.BaseMsgResp{
		Code: xerr.OK,
		Msg:  "batch delete user success"}, nil
}
