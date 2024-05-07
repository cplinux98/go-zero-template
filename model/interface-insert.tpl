Trans(ctx context.Context, fn func(context context.Context, session sqlx.Session) error) error
SelectBuilder() squirrel.SelectBuilder
InsertBuilder() squirrel.InsertBuilder
Insert(ctx context.Context, session sqlx.Session, data *{{.upperStartCamelObject}}) (sql.Result,error)
InsertMany(ctx context.Context, session sqlx.Session, list []*{{.upperStartCamelObject}}) (sql.Result, error)