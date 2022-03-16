package initialize

import (
	"brisklog/global"
	"brisklog/models"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/**
 * @author: xaohuihui
 * @Path: brisklog/initialize/mysql.go
 * @Description:
 * @datetime: 2022/3/16 18:25:46
 * software: GoLand
**/

// InitMysqlDB 初始化mysql数据库连接
func InitMysqlDB() {
	mysqlInfo := global.Settings.MysqlInfo
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlInfo.Name, mysqlInfo.Password, mysqlInfo.Host, mysqlInfo.Port, mysqlInfo.DBName)
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	global.DB = db

	// 自动创建user表
	global.DB.AutoMigrate(&models.User{})
}
