func (m *default{{.upperStartCamelObject}}Model) Update(ctx context.Context, session sqlx.Session, {{if .containsIndexCache}}newData{{else}}data{{end}} *{{.upperStartCamelObject}}) (sql.Result,error) {
	{{if .withCache}}{{if .containsIndexCache}}data, err:=m.FindOne(ctx, newData.{{.upperStartCamelPrimaryKey}})
	if err!=nil{
		return nil, err
	}

{{end}}	{{.keys}}
    return m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where {{.originalPrimaryKey}} = {{if .postgreSql}}$1{{else}}?{{end}}", m.table, {{.lowerStartCamelObject}}RowsWithPlaceHolder)

		if session != nil {
		    return session.ExecCtx(ctx, query, {{.expressionValues}})
		}

		return conn.ExecCtx(ctx, query, {{.expressionValues}})

	}, {{.keyValues}}){{else}}query := fmt.Sprintf("update %s set %s where {{.originalPrimaryKey}} = {{if .postgreSql}}$1{{else}}?{{end}}", m.table, {{.lowerStartCamelObject}}RowsWithPlaceHolder)

    if session != nil {
        return session.ExecCtx(ctx, query, {{.expressionValues}})
    }

    return m.conn.ExecCtx(ctx, query, {{.expressionValues}}){{end}}
}

// DeleteManyByIds
//  @Description: 根据id列表删除
//  @Author cplinux98 2024-05-07 19:26:41
//  @receiver m
//  @param ctx
//  @param session
//  @param ids  传入对应的主键类型数组，类似[]int64, []string
//  @return sql.Result
//  @return error
//
func (m *default{{.upperStartCamelObject}}Model) DeleteManyByIds(ctx context.Context, session sqlx.Session, ids interface{}) (sql.Result, error) {
	idWhere := m.SelectBuilder().Where(squirrel.Eq{"id": ids})
	{{if .withCache}}records{{else}}_{{end}}, err := m.FindManyByBuilderNoCache(ctx, idWhere, "id")
	if err != nil {
		return nil, err
	}
	{{if .withCache}}cacheKeys := make([]string, 0)
	for _, data := range records {
        {{.keys}}
		cacheKeys = append(cacheKeys, {{.keyValues}})
	}{{end}}

    query, values, err2 := m.DeleteBuilder().Where(squirrel.Eq{"id": ids}).ToSql()
    if err2 != nil {
        return nil, err2
    }

    {{if .withCache}}
	return m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (sql.Result, error) {

		if session != nil {
			return session.ExecCtx(ctx, query, values...)
		}
		return conn.ExecCtx(ctx, query, values...)
	}, cacheKeys...)
	{{else}}
    if session != nil {
        return session.ExecCtx(ctx, query, values...)
    }
    return m.conn.ExecCtx(ctx, query, values...)
	{{end}}
}
