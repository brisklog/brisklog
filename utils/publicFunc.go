package utils

import (
	"fmt"
	"time"
)

/**
 * @author: xaohuihui
 * @Path: brisklog/utils/publicFunc.go
 * @Description:
 * @datetime: 2022/3/16 17:55:20
 * software: GoLand
**/

// GetNowFormatTodayTime 获取当天年月日函数
func GetNowFormatTodayTime() string {
	now := time.Now()
	dateStr := fmt.Sprintf("%02d-%02d-%02d", now.Year(), int(now.Month()), now.Day())
	return dateStr
}
