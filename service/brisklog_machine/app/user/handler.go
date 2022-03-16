package user

import (
	"brisklog_machine/forms"
	"brisklog_machine/global"
	"brisklog_machine/models"
	"brisklog_machine/utils"
	customResponse "brisklog_machine/utils/Response"
	"fmt"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

func PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Pong!",
	})
}

func LoginHandler(c *gin.Context) {
	var PasswordLoginForm forms.PasswordLoginForm

	fmt.Println(&store)

	if err := c.ShouldBind(&PasswordLoginForm); err != nil {
		utils.HandleValidatorError(c, err)
		return
	}

	color.Yellow(store.Get(PasswordLoginForm.CaptchaId, false))
	fmt.Println(store.Verify(PasswordLoginForm.CaptchaId, PasswordLoginForm.Captcha, false))
	fmt.Println(PasswordLoginForm)
	if !store.Verify(PasswordLoginForm.CaptchaId, PasswordLoginForm.Captcha, true) {
		customResponse.Err(c, http.StatusBadRequest, 400, "验证码错误", "")
		return
	}
	user, ok := FindUserInfo(PasswordLoginForm.Username, PasswordLoginForm.Password)
	if !ok {
		customResponse.Err(c, http.StatusUnauthorized, 401, "该用户未注册", "")
		return
	}

	token := utils.CreateToken(c, user.ID, user.NickName, user.Role)
	userInfoMap := HandleUserModelToMap(user)
	userInfoMap["token"] = token
	c.SetCookie("x-token", token, 3600, "/", "", false, true)

	customResponse.Success(c, 200, "success", userInfoMap)
	return

}

func GetUserList(c *gin.Context) {
	// 获取参数
	UserListFrom := forms.UserListForm{}
	if err := c.ShouldBind(&UserListFrom); err != nil {
		utils.HandleValidatorError(c, err)
		return
	}

	total, userList := GetUserListBusiness(UserListFrom.Page, UserListFrom.Size)

	if (total + len(userList)) == 0 {
		customResponse.Err(c, http.StatusBadRequest, 400, "未获取到数据", gin.H{
			"total":    total,
			"userlist": userList,
		})
		return
	}
	customResponse.Success(c, http.StatusOK, "获取用户列表成功", map[string]interface{}{
		"total":    total,
		"userlist": userList,
	})
}

// base64Captcha 缓存对象
// var store = base64Captcha.DefaultMemStore
var store = base64Captcha.NewMemoryStore(1024, time.Second*5)

// GetCaptcha 获取验证码
func GetCaptcha(c *gin.Context) {
	fmt.Println(&store)

	driver := base64Captcha.NewDriverDigit(80, 240, 5, 0.7, 80)

	cp := base64Captcha.NewCaptcha(driver, store)
	// b64s是图片的bs64编码
	id, b64s, err := cp.Generate()

	if err != nil {
		zap.S().Error("生成验证码错误：%s", err.Error())
		customResponse.Err(c, http.StatusInternalServerError, 500, "生成验证码错误", "")
		return
	}
	customResponse.Success(c, http.StatusOK, "生成验证成功", gin.H{
		"captchaId": id,
		"picPath":   b64s,
	})
}

func HandleUserModelToMap(user *models.User) (userInfoMap map[string]interface{}) {
	birthday := ""
	if user.Birthday == nil {
		birthday = ""
	} else {
		birthday = user.Birthday.Format("2006-01-02")
	}
	userInfoMap = map[string]interface{}{
		"id":        user.ID,
		"password":  user.Password,
		"nick_name": user.NickName,
		"head_url":  user.HeadUrl,
		"birthday":  birthday,
		"address":   user.Address,
		"gender":    user.Gender,
		"role":      user.Role,
		"mobile":    user.Mobile,
	}
	return userInfoMap
}

// PutHeaderImage 上传用户头像
func PutHeaderImage(c *gin.Context) {
	file, _ := c.FormFile("file")
	fileObj, err := file.Open()
	if err != nil {
		zap.L().Error(err.Error())
		return
	}
	// 把文件保存在本地
	// 创建一个空文件
	headerUrl := strconv.FormatInt(time.Now().Unix(), 10) + "-" + file.Filename
	savefile, err := os.Create("static/uploadfiles/" + headerUrl)
	if err != nil {
		global.Lg.Error(err.Error())
		customResponse.Err(c, http.StatusUnauthorized, 401, "头像上传失败", "")
		return
	}
	defer savefile.Close()
	// 复制给空文件
	_, err = io.Copy(savefile, fileObj)
	if err != nil {
		global.Lg.Error(err.Error())
		customResponse.Err(c, http.StatusUnauthorized, 401, "头像上传失败", "")
		return
	}

	// TODO 把用户头像地址存入到对应的user表中head_url中
	userId := c.GetInt("userId")
	UpdateParam := map[string]interface{}{
		"head_url": headerUrl,
	}
	UpdateUserInfo(uint(userId), &UpdateParam)

	customResponse.Success(c, http.StatusOK, "头像上传成功", gin.H{
		"userheaderUrl": headerUrl,
	})
}

// GetHeaderImage 获取头像
func GetHeaderImage(c *gin.Context) {
	filename := "16262666091133707.jpg"
	//c.Header("Content-Type", "image/jpeg")
	c.Header("Content-Disposition", "attachment; filename="+filename)
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Cache-Control", "no-cache")
	filepath := "static/uploadfiles/" + filename
	c.File(filepath)
}
