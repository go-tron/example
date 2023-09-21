package config

import (
	"github.com/go-tron/config"
	"github.com/go-tron/iris/session"
	"github.com/go-tron/logger"
	"github.com/go-tron/mysql"
	"github.com/go-tron/nsq"
	"github.com/go-tron/redis"
	"github.com/go-tron/token"
	"github.com/kataras/iris/v12/middleware/basicauth"
)

var Config = config.New()

var BasicAuth = basicauth.Default(Config.GetStringMapString("basicAuth"))
var DB = mysql.NewWithConfig(Config)
var OrderLogger = logger.NewZapWithConfig(Config, "order", "info")
var NsqProducer = nsq.NewProducerWithConfig(Config)
var Redis = redis.NewWithConfig(Config)
var Sessions = session.NewSessionsWithConfig(Config, Redis.Client)
var UserToken = &token.IdBased{
	Redis:      Redis,
	Name:       "user-token",
	ExpireTime: 0,
	Repeatable: false,
	Disposable: false,
}
