package models

/**
 * @author: xaohuihui
 * @Path: brisklog/models/user.go
 * @Description:
 * @datetime: 2022/3/16 17:57:32
 * software: GoLand
**/

import "time"

type User struct {
	ID       uint       `json:"id" gorm:"primaryKey"`
	Password string     `json:"password"`
	NickName string     `json:"nick_name"`
	HeadUrl  string     `json:"head_url"`
	Birthday *time.Time `json:"birthday" gorm:"type:date"`
	Address  string     `json:"address"`
	Desc     string     `json:"desc"`
	Gender   string     `json:"gender"`
	Role     uint       `json:"role"`
	Mobile   string     `json:"mobile"`
}

// TableName 设置User的表名为 profiles
func (User) TableName() string {
	return "profiles"
}
