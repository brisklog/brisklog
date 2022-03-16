package utils

import (
	"brisklog/global"
	customResponse "brisklog/utils/Response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strings"
)

/**
 * @author: xaohuihui
 * @Path: brisklog/utils/validator.go
 * @Description:
 * @datetime: 2022/3/16 17:55:20
 * software: GoLand
**/

// HandleValidatorError 处理字段校验异常
func HandleValidatorError(c *gin.Context, err error) {
	// 如何返回错误信息
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		customResponse.Err(c, http.StatusInternalServerError, 500, "字段校验错误", err.Error())
	}
	msg := removeTopStruct(errs.Translate(global.Trans))
	customResponse.Err(c, http.StatusBadRequest, 400, "字段校验错误", msg)
	return
}

// removeTopStruct 定义一个去掉结构体名称前缀的自定义方法
func removeTopStruct(fields map[string]string) map[string]string {
	rsp := map[string]string{}
	for field, err := range fields {
		// 从文本的 . 开始切分 处理后"mobile": "mobile为必填字段" 处理前："PasswordLoginForm.mobile": "mobile为必填字段"
		rsp[field[strings.Index(field, ".")+1:]] = err
	}
	return rsp
}
