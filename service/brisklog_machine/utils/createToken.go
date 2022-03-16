package utils

import (
	"brisklog_machine/middlewares"
	customResponse "brisklog_machine/utils/Response"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

/**
 * @author: xaohuihui
 * @Path: brisklog/utils/createToken.go
 * @Description:
 * @datetime: 2022/3/16 17:55:20
 * software: GoLand
**/

// CreateToken 创建token
func CreateToken(c *gin.Context, Id uint, NickName string, Role uint) string {
	j := middlewares.NewJWT()
	claims := middlewares.CustomClaims{
		ID: Id,
		NickName: NickName,
		AuthorityId: Role,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),
			// TODO 设置token过期时间
			ExpiresAt: time.Now().Unix() + 60 * 60 * 5, // token 5小时过期
			Issuer: "xaohuihui",
		},
	}
	// 生成token
	token, err := j.CreateToken(claims)
	if err != nil {
		customResponse.Err(c, http.StatusUnauthorized, 401, "token生成失败，重新再试", "")
		return ""
	}
	return token
}
