// DeleteBuilder
//
//	@Description: 构建删除语句
//	@Author cplinux98 2024-05-06 00:02:47
//	@receiver m
//	@return squirrel.DeleteBuilder
func (m *default{{.upperStartCamelObject}}Model) DeleteBuilder() squirrel.DeleteBuilder {
	return squirrel.Delete(m.table)
}

func (m *default{{.upperStartCamelObject}}Model) Delete(ctx context.Context, session sqlx.Session, {{.lowerStartCamelPrimaryKey}} {{.dataType}}) (sql.Result,error) {
	{{if .withCache}}{{if .containsIndexCache}}data, err:=m.FindOne(ctx, {{.lowerStartCamelPrimaryKey}})
	if err!=nil{
		return nil, err
	}

{{end}}	{{.keys}}
    return m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where {{.originalPrimaryKey}} = {{if .postgreSql}}$1{{else}}?{{end}}", m.table)

        if session != nil {
            return session.ExecCtx(ctx, query, {{.lowerStartCamelPrimaryKey}})
        }

		return conn.ExecCtx(ctx, query, {{.lowerStartCamelPrimaryKey}})

	}, {{.keyValues}}){{else}}

	query := fmt.Sprintf("delete from %s where {{.originalPrimaryKey}} = {{if .postgreSql}}$1{{else}}?{{end}}", m.table)

    if session != nil {
        return session.ExecCtx(ctx, query, {{.lowerStartCamelPrimaryKey}})
    }

	return m.conn.ExecCtx(ctx, query, {{.lowerStartCamelPrimaryKey}}){{end}}
}


