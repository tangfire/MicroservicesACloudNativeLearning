package svc

import (
	"MicroservicesACloudNativeLearning/doc/004-Microservice_Framework/004-003-go-zero/003-gozero_mall_cache/service/user/api/internal/config"
	"MicroservicesACloudNativeLearning/doc/004-Microservice_Framework/004-003-go-zero/003-gozero_mall_cache/service/user/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UserModel // 加入User表增删改查操作Model
}

func NewServiceContext(c config.Config) *ServiceContext {
	// UserModel -> 接口类型
	// *defaultUserModel 实现了接口
	// 调用构造函数得到*model_nocache.defaultUserModel
	// NewUserModel(conn sqlx.SqlConn)
	// 需要 sqlx.SqlCon 的数据库连接

	sqlxConn := sqlx.NewMysql(c.Mysql.DataSource)

	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(sqlxConn, c.CacheRedis),
	}
}
