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
