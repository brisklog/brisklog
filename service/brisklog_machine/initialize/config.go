package initialize

/**
 * @author: xaohuihui
 * @Path: brisklog/initialize/config.go
 * @Description: 初始化配置
 * @datetime: 2022/3/16 17:41:10
 * software: GoLand
**/

import (
	"brisklog_machine/config"
	"brisklog_machine/global"
	"github.com/fatih/color"
	"github.com/spf13/viper"
)

func InitConfig() {
	// 实例化viper
	v := viper.New()
	// 文件的路径设置
	v.SetConfigFile("./settings-dev.yaml")
	if err := v.ReadInConfig(); err != nil {
		if err := v.ReadInConfig(); err != nil {
			panic(err)
		}
	}

	serverConfig := config.ServerConfig{}
	// 给serverConfig 初始值
	if err := v.Unmarshal(&serverConfig); err != nil {
		panic(err)
	}
	// 传递全局变量
	global.Settings = serverConfig
	color.Blue("初始化环境变量", global.Settings.LogsAddress)
}
