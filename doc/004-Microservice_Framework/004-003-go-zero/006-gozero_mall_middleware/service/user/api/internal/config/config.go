package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Mysql struct {
		DataSource string
	}

	Auth struct {
		AccessSecret string // jwt密钥
		AccessExpire int64  // 有效期,单位:s
	}
	CacheRedis cache.CacheConf
}
