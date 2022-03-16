package user

import (
	"brisklog_machine/middlewares"
	"github.com/gin-gonic/gin"
)

func Routers(r *gin.Engine) {
	userRouter := r.Group("/user")
	{
		userRouter.GET("/ping", PingHandler)
		userRouter.POST("/login", LoginHandler)
		userRouter.GET("/list", middlewares.JWTAuth(), middlewares.IsAdminAuth(), GetUserList)
		userRouter.GET("/captcha", GetCaptcha)
		userRouter.POST("/upload/header-image", middlewares.JWTAuth(), PutHeaderImage)
		userRouter.GET("/download/header-image", GetHeaderImage)
	}
}
