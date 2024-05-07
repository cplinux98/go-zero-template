func (m *default{{.upperStartCamelObject}}Model) FindOne(ctx context.Context, {{.lowerStartCamelPrimaryKey}} {{.dataType}}) (*{{.upperStartCamelObject}}, error) {
	{{if .withCache}}{{.cacheKey}}
	var resp {{.upperStartCamelObject}}
	err := m.QueryRowCtx(ctx, &resp, {{.cacheKeyVariable}}, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query :=  fmt.Sprintf("select %s from %s where {{.originalPrimaryKey}} = {{if .postgreSql}}$1{{else}}?{{end}} limit 1", {{.lowerStartCamelObject}}Rows, m.table)
		return conn.QueryRowCtx(ctx, v, query, {{.lowerStartCamelPrimaryKey}})
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}{{else}}query := fmt.Sprintf("select %s from %s where {{.originalPrimaryKey}} = {{if .postgreSql}}$1{{else}}?{{end}} limit 1", {{.lowerStartCamelObject}}Rows, m.table)
	var resp {{.upperStartCamelObject}}
	err := m.conn.QueryRowCtx(ctx, &resp, query, {{.lowerStartCamelPrimaryKey}})
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}{{end}}
}


// FindManyByBuilderNoCache
//
//	@Description: 根据builder进行查询，没有使用缓存
//	@Author cplinux98 2024-05-06 11:16:56
//	@receiver m
//	@param ctx
//	@param builder
//	@param orderBy  排序语句
//	@return []*{{.upperStartCamelObject}}
//	@return error
func (m *default{{.upperStartCamelObject}}Model) FindManyByBuilderNoCache(ctx context.Context, builder squirrel.SelectBuilder, orderBy string) ([]*{{.upperStartCamelObject}}, error) {
	builder = builder.Columns({{.lowerStartCamelObject}}Rows)

	if orderBy == "" {
		builder = builder.OrderBy("id DESC")
	} else {
		builder = builder.OrderBy(orderBy)
	}

	query, values, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*{{.upperStartCamelObject}}
	{{if .withCache}}
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
    {{else}}
    err = m.conn.QueryRowsCtx(ctx, &resp, query, values...)
    {{end}}

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
//	@receiver m
//	@param ctx
//	@param builder
//	@param field count的字段名
//	@return uint64
//	@return error
func (m *default{{.upperStartCamelObject}}Model) FindCountByBuilderNoCache(ctx context.Context, builder squirrel.SelectBuilder, field string) (uint64, error) {
	if len(field) == 0 {
		return 0, pkgErrors.Wrapf(pkgErrors.New("FindCountByBuilderNoCache Least One Field"), "FindCountByBuilderNoCache Least One Field")
	}

	builder = builder.Columns("COUNT(" + field + ")")
	query, values, err := builder.ToSql()

	if err != nil {
		return 0, err
	}

	var resp uint64
    {{if .withCache}}
    err = m.QueryRowNoCacheCtx(ctx, &resp, query, values...)
    {{else}}
    err = m.conn.QueryRowCtx(ctx, &resp, query, values...)
    {{end}}

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
//	@receiver m
//	@param ctx
//	@param builder  查询参数构建
//	@param page     页数
//	@param pageSize 每页多少条
//	@param orderBy  orderBy语句
//	@return []*{{.upperStartCamelObject}}
//	@return uint64  当前条件下的总数
//	@return error
func (m *default{{.upperStartCamelObject}}Model) FindPageListByPageWithTotal(ctx context.Context, builder squirrel.SelectBuilder, page, pageSize uint64, orderBy string) ([]*{{.upperStartCamelObject}}, uint64, error) {
	total, err := m.FindCountByBuilderNoCache(ctx, builder, "id")

	if err != nil {
		return nil, 0, err
	}

	builder = builder.Columns({{.lowerStartCamelObject}}Rows)

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

	var resp []*{{.upperStartCamelObject}}
    {{if .withCache}}
    err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
    {{else}}
    err = m.conn.QueryRowsCtx(ctx, &resp, query, values...)
    {{end}}

	switch err {
	case nil:
		return resp, total, nil
	default:
		return nil, total, err
	}
}

