package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	// mysql
	// redis

	Mysql struct {
		DataSource string
	}

	CacheRedis cache.CacheConf
}
