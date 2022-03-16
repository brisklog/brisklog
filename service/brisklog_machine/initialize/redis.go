package initialize

import (
	"brisklog_machine/global"
	"fmt"
	"github.com/fatih/color"
	"github.com/go-redis/redis"
)

/**
 * @author: xaohuihui
 * @Path: brisklog/initialize/redis.go
 * @Description:
 * @datetime: 2022/3/16 18:26:56
 * software: GoLand
**/

func InitRedis() {
	add := fmt.Sprintf("%s:%d", global.Settings.RedisInfo.Host, global.Settings.RedisInfo.Port)
	// 生成redis客户端
	global.Redis = redis.NewClient(&redis.Options{
		Addr:     add,
		Password: global.Settings.RedisInfo.Password,
		DB:       global.Settings.RedisInfo.DB,
	})
	// 连接redis
	_, err := global.Redis.Ping().Result()
	if err != nil {
		color.Red("[InitRedis] 连接redis异常")
		color.Yellow(global.Settings.RedisInfo.Host)
		color.Yellow(err.Error())
	}
}
