package Response

/**
 * @author: xaohuihui
 * @Path: brisklog/utils/Response/response.go
 * @Description:
 * @datetime: 2022/3/16 17:54:32
 * software: GoLand
**/

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Success(c *gin.Context, code int, msg interface{}, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
	return
}

func Err(c *gin.Context, httpCode int, code int, msg string, jsonStr interface{}) {
	c.JSON(httpCode, gin.H{
		"code": code,
		"msg":  msg,
		"data": jsonStr,
	})
	return
}
