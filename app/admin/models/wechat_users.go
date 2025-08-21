package models

import (
	"go-admin/common/models"
)

type WechatUsers struct {
	models.Model

	OpenId      string `json:"openId" gorm:"type:varchar(64);comment:微信 OpenID"`
	NickName    string `json:"nickName" gorm:"type:varchar(64);comment:昵称"`
	AvatarUrl   string `json:"avatarUrl" gorm:"type:varchar(255);comment:头像 URL"`
	PhoneNumber string `json:"phoneNumber" gorm:"type:varchar(20);comment:手机号码"`
	FaceCount   string `json:"faceCount" gorm:"type:smallint(5) unsigned;comment:人脸计数"`
	TongueCount string `json:"tongueCount" gorm:"type:smallint(5) unsigned;comment:舌头计数"`
	models.ModelTime
	models.ControlBy
}

func (WechatUsers) TableName() string {
	return "wechat_users"
}

func (e *WechatUsers) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *WechatUsers) GetId() interface{} {
	return e.Id
}
