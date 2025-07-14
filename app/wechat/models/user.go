package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	OpenID      string `gorm:"uniqueIndex"`
	NickName    string
	AvatarURL   string
	PhoneNumber string
	FaceCount   uint16
	TongueCount uint16
}

// TableName 指定了模型对应的数据库表名
func (User) TableName() string {
	return "wechat_users"
}
