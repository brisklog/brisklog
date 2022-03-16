package middlewares


import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
跨域是指浏览器不能执行其他网站的脚本。他是浏览器的同源策略造成的，是浏览器对JavaScript施加的安全策略
浏览器在什么情况下会发起options预检请求？
在非简单请求且跨域的情况下，浏览器挥发期options预检请求。
Preflighted Requests是CORS中一种透明服务器验证机制。
预检请求首先需要向另外一个域名的资源发送一个HTTP OPTIONS 请求头，
目的是为了判断实际发送的请求是否安全。
*/

// Cors 返回的Response添加约定的请求头
// 请求方式如果是OPTIONS直接返回204
// 跨域中间件： 后端开启跨域功能
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers",
			"Content-Type,AccessToken,X-CSRF-Token,Authorization,Token,x-token")
		c.Header("Access-Control-Allow-Methods", "POST,GET,OPTIONS,DELETE,PATCH,PUT")
		c.Header("Access-Control-Expose-Header",
			"Content-Length,Access-Control-Allow-Origin,Access-Control-Allow-Headers,Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
	}
}
