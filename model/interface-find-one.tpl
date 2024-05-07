FindOne(ctx context.Context, {{.lowerStartCamelPrimaryKey}} {{.dataType}}) (*{{.upperStartCamelObject}}, error)
FindManyByBuilderNoCache(ctx context.Context, builder squirrel.SelectBuilder, orderBy string) ([]*{{.upperStartCamelObject}}, error)
FindCountByBuilderNoCache(ctx context.Context, builder squirrel.SelectBuilder, field string) (uint64, error)
FindPageListByPageWithTotal(ctx context.Context, builder squirrel.SelectBuilder, page, pageSize uint64, orderBy string) ([]*{{.upperStartCamelObject}}, uint64, error)