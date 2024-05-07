// Trans
//
//	@Description: 事务
//	@Author cplinux98 2024-05-07 18:45:47
//	@receiver m
//	@param ctx
//	@param fn
//	@return error
func (m *default{{.upperStartCamelObject}}Model) Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {

    return m.{{if not .withCache}}conn.{{end}}TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
        return fn(ctx, session)
    })

}

// SelectBuilder
//
//	@Description: 构建查询语句
//	@Author cplinux98 2024-05-06 00:04:41
//	@receiver m
//	@return squirrel.SelectBuilder
func (m *default{{.upperStartCamelObject}}Model) SelectBuilder() squirrel.SelectBuilder {
	return squirrel.Select().From(m.table)
}

// InsertBuilder
//
//	@Description: 构建插入语句
//	@Author cplinux98 2024-05-06 00:01:27
//	@receiver m
//	@return squirrel.InsertBuilder
func (m *default{{.upperStartCamelObject}}Model) InsertBuilder() squirrel.InsertBuilder {
	return squirrel.Insert(m.table)
}

func (m *default{{.upperStartCamelObject}}Model) Insert(ctx context.Context, session sqlx.Session, data *{{.upperStartCamelObject}}) (sql.Result,error) {
	{{if .withCache}}{{.keys}}
    return m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values ({{.expression}})", m.table, {{.lowerStartCamelObject}}RowsExpectAutoSet)
		if session != nil {
		    return session.ExecCtx(ctx, query, {{.expressionValues}})
		}
		return conn.ExecCtx(ctx, query, {{.expressionValues}})
	}, {{.keyValues}})

	{{else}}query := fmt.Sprintf("insert into %s (%s) values ({{.expression}})", m.table, {{.lowerStartCamelObject}}RowsExpectAutoSet)
	if session != nil {
	    return session.ExecCtx(ctx, query, {{.expressionValues}})
	}
    return m.conn.ExecCtx(ctx, query, {{.expressionValues}}){{end}}
}

// InsertMany
//  @Description: 批量新增
//  @Author cplinux98 2024-05-07 18:49:26
//  @receiver m
//  @param ctx
//  @param session
//  @param list
//  @return sql.Result
//  @return error
//
func (m *default{{.upperStartCamelObject}}Model) InsertMany(ctx context.Context, session sqlx.Session, list []*{{.upperStartCamelObject}}) (sql.Result, error) {
	{{if .withCache}}cacheKeys := make([]string, 0){{end}}
	insertBuilder := m.InsertBuilder().Columns({{.lowerStartCamelObject}}RowsExpectAutoSet)

    for _, data := range list {
        {{if .withCache}}
        {{.keys}}
        cacheKeys = append(cacheKeys, {{.keyValues}})
        {{end}}
        insertBuilder = insertBuilder.Values({{.expressionValues}})
    }


	query, values, err2 := insertBuilder.ToSql()
	if err2 != nil {
		return nil, err2
	}
    {{if .withCache}}return m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		if session != nil {
			return session.ExecCtx(ctx, query, values...)
		}
		return conn.ExecCtx(ctx, query, values...)
	}, cacheKeys...){{else}}
    if session != nil {
        return session.ExecCtx(ctx, query, values...)
    }
    return m.conn.ExecCtx(ctx, query, values...)
    {{end}}
}


