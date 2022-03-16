package user

import (
	"brisklog_machine/global"
	"brisklog_machine/models"
	"fmt"
)

var users []models.User

func GetUserListBusiness(page int, size int) (int, []interface{}) {
	// 分页用户列表数据
	userList := make([]interface{}, 0, len(users))
	// 计算偏移量
	offset := (page - 1) * size
	// 查询所有的user
	result := global.DB.Offset(offset).Limit(size).Find(&users)
	// 查不到数据是
	if result.RowsAffected == 0 {
		return 0, userList
	}
	// 获取user总数
	total := len(users)
	// 查询数据
	result.Offset(offset).Limit(size).Find(&users)

	for _, useSingle := range users {
		birthday := ""
		if useSingle.Birthday == nil {
			birthday = ""
		} else {
			birthday = useSingle.Birthday.Format("2006-01-02")
		}
		userItemMap := map[string]interface{}{
			"id":        useSingle.ID,
			"password":  useSingle.Password,
			"nick_name": useSingle.NickName,
			"head_url":  useSingle.HeadUrl,
			"birthday":  birthday,
			"address":   useSingle.Address,
			"gender":    useSingle.Gender,
			"role":      useSingle.Role,
			"mobile":    useSingle.Mobile,
		}
		userList = append(userList, userItemMap)
	}
	return total, userList
}

// FindUserInfo 通过username找到用户信息
func FindUserInfo(username string, password string) (*models.User, bool) {
	var user models.User
	// 查询用户
	rows := global.DB.Where(&models.User{NickName: username, Password: password}).Find(&user)
	fmt.Println(&user)
	if rows.RowsAffected < 1 {
		return &user, false
	}
	return &user, true
}

func UpdateUserInfo(userId uint, updateParam *map[string]interface{}) {
	var user models.User
	row := global.DB.Where(&models.User{ID: userId}).First(&user)
	row.Updates(updateParam)
}
