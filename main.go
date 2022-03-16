package brisklog

import (
	"brisklog/app/user"
	"brisklog/global"
	"brisklog/initialize"
	"fmt"
	"github.com/gin-gonic/gin"
)

/**
 * @author: xaohuihui
 * @Path: brisklog/main
 * @Description: 项目入口
 * @datetime: 2022/3/16 17:26:10
 * software: GoLand
**/

func main() {
	// 强制日志颜色化
	gin.ForceConsoleColor()

	// 初始化yaml配置
	initialize.InitConfig()

	// 初始化翻译
	if err := initialize.InitTrans("zh"); err != nil {
		panic(err)
	}

	// 加载多个APP的路由配置
	initialize.Include(user.Routers)
	// 初始化路由
	r := initialize.InitRouters()

	// 初始换日志信息
	initialize.InitLogger()

	// 初始化mysql
	initialize.InitMysqlDB()

	// 初始化redis
	initialize.InitRedis()

	if err := r.Run(fmt.Sprintf(":%d", global.Settings.Port)); err != nil {
		fmt.Printf("startup service failed, err: %v\n", err)
	}
}