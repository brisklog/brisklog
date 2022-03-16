package middlewares

import (
	"brisklog/global"
	customResponse "brisklog/utils/Response"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

/*
我对中间件的理解:
1.类似于koa的洋葱模型,Django的钩子函数,前端的生命周期
2.在请求接口前或者后做一些逻辑出来
中间件几个关键字:
1. c.Next() 进入下一个中间件
1. c.Abort() 中断中间件(return 不能中断中间件的调用)
*/

type CustomClaims struct {
	ID          uint
	NickName    string
	AuthorityId uint
	jwt.StandardClaims
}

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 我们这里jwt鉴权取头部信息 x-token 登录时间返回token信息
		// 这里前端需要把token存储到cookie或者本地localStorage中
		//不过需要跟后端协商过期时间 可以约定刷新令牌或者重新登录
		var token string
		cookieToken, err := c.Cookie("x-token")
		fmt.Println(cookieToken, err)
		if err == nil && cookieToken != "" {
			token = cookieToken
		} else {
			token = c.Request.Header.Get("x-token")
		}
		color.Yellow(token)
		if token == "" {
			customResponse.Err(c, http.StatusUnauthorized, 401, "请登录", "")
			c.Abort()
			return
		}
		j := NewJWT()
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == TokenExpired {
				customResponse.Err(c, http.StatusUnauthorized, 401, "授权已过期", "")
				c.Abort()
				return
			} else if err == TokenNotValidYet {
				customResponse.Err(c, http.StatusUnauthorized, 401, "授权此处无效", "")
			} else if err == TokenMalformed {
				customResponse.Err(c, http.StatusUnauthorized, 401, "token格式错误", "")
			} else {
				customResponse.Err(c, http.StatusUnauthorized, 401, "未登录", "")
			}
		}
		fmt.Println(c)
		// gin的上下文记录claims和userId的值
		c.Set("claims", claims)
		c.Set("userId", claims.ID)
		userId := c.GetInt("userId")
		color.Red("______________________")
		fmt.Println(userId)
		c.Next()
	}
}

type JWT struct {
	SigningKey []byte
}

var (
	TokenExpired     = errors.New("token is expired ")
	TokenNotValidYet = errors.New("token not active yet ")
	TokenMalformed   = errors.New("that's not even a token ")
	TokenInvalid     = errors.New("couldn't handle this token:")
)

func NewJWT() *JWT {
	return &JWT{
		[]byte(global.Settings.JWTKey.SigningKey),
	}
}

// CreateToken 创建一个token
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// ParseToken 解析token
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				// token 格式错误
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// token 过期
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				// token 是有效地，然而在此处是无效的
				return nil, TokenNotValidYet
			} else {
				// token 无效
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid
	} else {
		return nil, TokenInvalid
	}
}


// RefreshToken 更新token
func (j *JWT) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return j.CreateToken(*claims)
	}
	return "", TokenInvalid
}
