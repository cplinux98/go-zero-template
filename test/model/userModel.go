package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Masterminds/squirrel"
	pkgErrors "github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserModel = (*customUserModel)(nil)

type (
	// UserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserModel.
	UserModel interface {
		userModel
		Trans(ctx context.Context, fn func(context context.Context, session sqlx.Session) error) error
		InsertBuilder() squirrel.InsertBuilder
		DeleteBuilder() squirrel.DeleteBuilder
		UpdateBuilder() squirrel.UpdateBuilder
		SelectBuilder() squirrel.SelectBuilder
		InsertByBuilderNoCache(ctx context.Context, session sqlx.Session, builder squirrel.InsertBuilder) (sql.Result, error)
		InsertByBuilderWithCache(ctx context.Context, session sqlx.Session, list []*User) (sql.Result, error)
		DeleteByBuilderNoCache(ctx context.Context, session sqlx.Session, builder squirrel.DeleteBuilder) (sql.Result, error)
		DeleteManyByIds(ctx context.Context, session sqlx.Session, ids []int64) (sql.Result, error)
		UpdateByBuilderNoCache(ctx context.Context, session sqlx.Session, builder squirrel.UpdateBuilder) (sql.Result, error)
		FindManyByBuilderNoCache(ctx context.Context, builder squirrel.SelectBuilder, orderBy string) ([]*User, error)
		FindCountByBuilderNoCache(ctx context.Context, builder squirrel.SelectBuilder, field string) (uint64, error)
		FindPageListByPageWithTotal(ctx context.Context, builder squirrel.SelectBuilder, page, pageSize uint64, orderBy string) ([]*User, uint64, error)
		FindOneByMobile(ctx context.Context, mobile string) (*User, error)
	}

	customUserModel struct {
		*defaultUserModel
	}
)

func (c *customUserModel) FindOneByMobile(ctx context.Context, mobile string) (*User, error) {
	userMobileKey := fmt.Sprintf("%s%v", cacheUserMobilePrefix, mobile)
	var resp User
	err := c.QueryRowIndexCtx(ctx, &resp, userMobileKey, c.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (any, error) {
		query := fmt.Sprintf("select %s from %s where mobile = ? limit 1", userRows, c.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, mobile); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, c.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (c *customUserModel) Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {

	return c.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})

}

// InsertBuilder
//
//	@Description: 构建插入语句
//	@Author cplinux98 2024-05-06 00:01:27
//	@receiver c
//	@return squirrel.InsertBuilder
func (c *customUserModel) InsertBuilder() squirrel.InsertBuilder {
	return squirrel.Insert(c.table)
}

// DeleteBuilder
//
//	@Description: 构建删除语句
//	@Author cplinux98 2024-05-06 00:02:47
//	@receiver c
//	@return squirrel.DeleteBuilder
func (c *customUserModel) DeleteBuilder() squirrel.DeleteBuilder {
	return squirrel.Delete(c.table)
}

// UpdateBuilder
//
//	@Description: 构建更新语句
//	@Author cplinux98 2024-05-06 00:03:07
//	@receiver c
//	@return squirrel.UpdateBuilder
func (c *customUserModel) UpdateBuilder() squirrel.UpdateBuilder {
	return squirrel.Update(c.table)
}

// SelectBuilder
//
//	@Description: 构建查询语句
//	@Author cplinux98 2024-05-06 00:04:41
//	@receiver c
//	@return squirrel.SelectBuilder
func (c *customUserModel) SelectBuilder() squirrel.SelectBuilder {
	return squirrel.Select().From(c.table)
}

// InsertByBuilderNoCache
//
//	@Description: 使用builder构建插入语句，没有使用缓存
//	@Author cplinux98 2024-05-06 10:56:44
//	@receiver m
//	@param ctx
//	@param session
//	@param insertBuilder
//	@return sql.Result
//	@return error
func (c *customUserModel) InsertByBuilderNoCache(ctx context.Context, session sqlx.Session, builder squirrel.InsertBuilder) (sql.Result, error) {
	query, values, err := builder.ToSql()
	if err != nil {
		return nil, err
	}
	if session != nil {
		return session.ExecCtx(ctx, query, values...)
	}

	return c.ExecNoCacheCtx(ctx, query, values...)
}

func (c *customUserModel) InsertByBuilderWithCache(ctx context.Context, session sqlx.Session, list []*User) (sql.Result, error) {
	// 构建要清理的缓存和构建插入语句
	keys := make([]string, 0)
	insertBuilder := c.InsertBuilder().Columns(userRowsExpectAutoSet)

	for _, data := range list {
		keys = append(keys, fmt.Sprintf("%s%v", cacheUserIdPrefix, data.Id), fmt.Sprintf("%s%v", cacheUserMobilePrefix, data.Mobile))
		insertBuilder = insertBuilder.Values(data.Mobile, data.Password, data.Nickname, data.Sex, data.Avatar, data.Info)
	}
	query, values, err2 := insertBuilder.ToSql()
	if err2 != nil {
		return nil, err2
	}

	ret, err := c.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		if session != nil {
			return session.ExecCtx(ctx, query, values...)
		}
		return conn.ExecCtx(ctx, query, values...)
	}, keys...)
	return ret, err
}

// DeleteByBuilderNoCache
//
//	@Description: 使用builder构建删除语句，没有使用缓存
//	@Author cplinux98 2024-05-06 11:03:29
//	@receiver c
//	@param ctx
//	@param session
//	@param builder
//	@return sql.Result
//	@return error
func (c *customUserModel) DeleteByBuilderNoCache(ctx context.Context, session sqlx.Session, builder squirrel.DeleteBuilder) (sql.Result, error) {
	query, values, err := builder.ToSql()
	if err != nil {
		return nil, err
	}
	if session != nil {
		return session.ExecCtx(ctx, query, values...)
	}

	return c.ExecNoCacheCtx(ctx, query, values...)
}

func (c *customUserModel) DeleteManyByIds(ctx context.Context, session sqlx.Session, ids []int64) (sql.Result, error) {
	idWhere := c.SelectBuilder().Where(squirrel.Eq{"id": ids})
	deleteRecords, err := c.FindManyByBuilderNoCache(ctx, idWhere, "id")
	if err != nil {
		return nil, err
	}
	cacheKeys := make([]string, 0)
	for _, deleteRecord := range deleteRecords {
		userIdKey := fmt.Sprintf("%s%v", cacheUserIdPrefix, deleteRecord.Id)
		userMobileKey := fmt.Sprintf("%s%v", cacheUserMobilePrefix, deleteRecord.Mobile)
		cacheKeys = append(cacheKeys, userIdKey, userMobileKey)
	}
	fmt.Println(cacheKeys)
	return c.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (sql.Result, error) {
		query, values, err2 := c.DeleteBuilder().Where(squirrel.Eq{"id": ids}).ToSql()
		if err2 != nil {
			return nil, err2
		}

		if session != nil {
			return session.ExecCtx(ctx, query, values...)
		}
		return conn.ExecCtx(ctx, query, values...)
	}, cacheKeys...)
}

// UpdateByBuilderNoCache
//
//	@Description: 使用builder构建更新语句，没有使用缓存
//	@Author cplinux98 2024-05-06 11:09:44
//	@receiver c
//	@param ctx
//	@param session
//	@param builder
//	@return sql.Result
//	@return error
func (c *customUserModel) UpdateByBuilderNoCache(ctx context.Context, session sqlx.Session, builder squirrel.UpdateBuilder) (sql.Result, error) {
	query, values, err := builder.ToSql()
	if err != nil {
		return nil, err
	}
	if session != nil {
		return session.ExecCtx(ctx, query, values...)
	}

	return c.ExecNoCacheCtx(ctx, query, values...)

}

// FindManyByBuilderNoCache
//
//	@Description: 根据builder进行查询，没有缓存
//	@Author cplinux98 2024-05-06 11:16:56
//	@receiver c
//	@param ctx
//	@param builder
//	@param orderBy  排序语句
//	@return []*User
//	@return error
func (c *customUserModel) FindManyByBuilderNoCache(ctx context.Context, builder squirrel.SelectBuilder, orderBy string) ([]*User, error) {
	builder = builder.Columns(userRows)

	if orderBy == "" {
		builder = builder.OrderBy("id DESC")
	} else {
		builder = builder.OrderBy(orderBy)
	}

	query, values, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*User
	err = c.QueryRowsNoCacheCtx(ctx, &resp, query, values...)

	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}

}

// FindCountByBuilderNoCache
//
//	@Description: 根据builder进行count，没有缓存
//	@Author cplinux98 2024-05-06 11:22:31
//	@receiver c
//	@param ctx
//	@param builder
//	@param field count的字段名
//	@return uint64
//	@return error
func (c *customUserModel) FindCountByBuilderNoCache(ctx context.Context, builder squirrel.SelectBuilder, field string) (uint64, error) {
	if len(field) == 0 {
		return 0, pkgErrors.Wrapf(pkgErrors.New("FindCountByBuilderNoCache Least One Field"), "FindCountByBuilderNoCache Least One Field")
	}

	builder = builder.Columns("COUNT(" + field + ")")
	query, values, err := builder.ToSql()
	//fmt.Println(err)
	if err != nil {
		return 0, err
	}

	var resp uint64
	err = c.QueryRowNoCacheCtx(ctx, &resp, query, values...)

	switch err {
	case nil:
		return resp, nil
	default:
		return 0, err
	}
}

// FindPageListByPageWithTotal
//
//	@Description: 根据分页参数查询
//	@Author cplinux98 2024-05-06 11:32:51
//	@receiver c
//	@param ctx
//	@param builder  查询参数构建
//	@param page     页数
//	@param pageSize 每页多少条
//	@param orderBy  orderBy语句
//	@return []*User
//	@return uint64  当前条件下的总数
//	@return error
func (c *customUserModel) FindPageListByPageWithTotal(ctx context.Context, builder squirrel.SelectBuilder, page, pageSize uint64, orderBy string) ([]*User, uint64, error) {
	total, err := c.FindCountByBuilderNoCache(ctx, builder, "id")

	if err != nil {
		return nil, 0, err
	}

	builder = builder.Columns(userRows)

	if orderBy == "" {
		builder = builder.OrderBy("id DESC")
	} else {
		builder = builder.OrderBy(orderBy)
	}

	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize
	query, values, err := builder.Offset(offset).Limit(pageSize).ToSql()

	if err != nil {
		return nil, total, err
	}

	var resp []*User
	err = c.QueryRowsNoCacheCtx(ctx, &resp, query, values...)

	switch err {
	case nil:
		return resp, total, nil
	default:
		return nil, total, err
	}
}

// NewUserModel returns a model for the database table.
func NewUserModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) UserModel {
	return &customUserModel{
		defaultUserModel: newUserModel(conn, c, opts...),
	}
}
