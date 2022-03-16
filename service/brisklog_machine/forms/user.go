package forms

/**
 * @author: xaohuihui
 * @Path: brisklog/forms/user.go
 * @Description:
 * @datetime: 2022/3/16 18:10:59
 * software: GoLand
**/

type PasswordLoginForm struct {
	// 密码 binding:“required” 为必填字段，长度大于3小于20
	Password  string `form:"password" json:"password" binding:"required,min=3,max=20,passwordverify"`
	Username  string `form:"username" json:"username" binding:"required,userverify"`
	Captcha   string `form:"captcha" json:"captcha" binding:"required,min=5,max=5"`
	CaptchaId string `form:"captcha_id" json:"captcha_id" binding:"required,required"`
}

type UserListForm struct {
	Page int `form:"page" json:"page" binding:"required,min=1"`
	Size int `form:"size" json:"size" binding:"required,min=10,max=100"`
}
