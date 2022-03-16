package global

import (
	"brisklog_machine/config"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

/**
 * @author: xaohuihui
 * @Path: brisklog/global/globalVar.go
 * @Description:
 * @datetime: 2022/3/16 17:28:50
 * software: GoLand
**/

var (
	Settings config.ServerConfig
	Lg       *zap.Logger
	Trans    ut.Translator
	DB       *gorm.DB
	Redis    *redis.Client
)
