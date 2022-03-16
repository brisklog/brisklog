package initialize

import (
	"brisklog/global"
	"brisklog/validation"
	"fmt"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

/**
 * @author: xaohuihui
 * @Path: brisklog/initialize/validator.go
 * @Description:
 * @datetime: 2022/3/16 18:28:45
 * software: GoLand
**/

// InitTrans validator 信息翻译
func InitTrans(locale string) (err error) {
	color.Red("给InitTrans传递一个参数,判断加载什么语言包,然后获取到语言包赋值给全局翻译器")
	// 修改gin框架中的validator引擎属性，实现定制
	// 加载自定义验证注册
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		//// 注冊自定义验证方法
		//v.RegisterValidation("userverify", validation.UserVerify)
		//v.RegisterValidation("passwordverify", validation.PasswordVerify)

		zhT := zh.New() // 中文翻译器
		enT := en.New() // 英文翻译器
		// 第一个参数是备用的语言环境，后面的参数是应该支持的语言环境
		uni := ut.New(enT, zhT, enT)
		global.Trans, ok = uni.GetTranslator(locale)
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s)", locale)
		}

		switch locale {
		case "en":
			_ = en_translations.RegisterDefaultTranslations(v, global.Trans)
		case "zh":
			_ = zh_translations.RegisterDefaultTranslations(v, global.Trans)
		default:
			_ = en_translations.RegisterDefaultTranslations(v, global.Trans)
		}

		// 注册自定义校验器
		RegisterValidatorFunc(v, "userverify", "非法用户名", validation.UserVerify)
		RegisterValidatorFunc(v, "passwordverify", "非法密码", validation.PasswordVerify)

		return nil
	}
	return nil
}

// Func myvalidator.validate
type myFunc validator.Func

// RegisterValidatorFunc 注册自定义校验tag
func RegisterValidatorFunc(v *validator.Validate, tag string, msgStr string, fn myFunc) {
	// 注册tag 自定义校验
	_ = v.RegisterValidation(tag, validator.Func(fn))
	// 自定义错误内容
	_ = v.RegisterTranslation(tag, global.Trans, func(ut ut.Translator) error {
		return ut.Add(tag, "{0}"+msgStr, true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T(tag, fe.Field())
		return t
	})
	return
}
