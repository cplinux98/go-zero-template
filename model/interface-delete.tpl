DeleteBuilder() squirrel.DeleteBuilder
Delete(ctx context.Context, session sqlx.Session, {{.lowerStartCamelPrimaryKey}} {{.dataType}}) (sql.Result,error)
DeleteManyByIds(ctx context.Context, session sqlx.Session, ids interface{}) (sql.Result, error)