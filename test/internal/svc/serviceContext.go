package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"test/internal/config"
	"test/model"
)

type ServiceContext struct {
	Config      config.Config
	RedisClient *redis.Redis

	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {

	sqlConn := sqlx.NewMysql(c.DB.DataSource)

	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(sqlConn, c.Cache),
	}
}
