package middlewares

import (
	"brisklog/global"
	"brisklog/models"
	customResponse "brisklog/utils/Response"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
 * @author: xaohuihui
 * @Path: brisklog/middlewares/admin.go
 * @Description:
 * @datetime: 2022/3/16 17:56:14
 * software: GoLand
**/

// IsAdminAuth 权限认证中间件：判断用户的角色和是否有权限
func IsAdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取token信息
		claims, _ := c.Get("claims")
		// 获取现在用户信息
		currentUser := claims.(*CustomClaims)

		// 判断role权限与用户权限是否相同
		userId := c.GetInt("userId")
		var user models.User
		rows := global.DB.Where(&models.User{ID: uint(userId)}).Find(&user)
		fmt.Println(&user)
		if rows.RowsAffected < 1 {
			customResponse.Err(c, http.StatusUnauthorized, 401, "用户不存在", "")
			c.Abort()
			return
		}

		if currentUser.AuthorityId != (&user).Role {
			customResponse.Err(c, http.StatusForbidden, 403, "用户没有权限", "")
			// 中断下面中间件
			c.Abort()
			return
		}
		// 继续执行下面中间件
		c.Next()
	}

}
