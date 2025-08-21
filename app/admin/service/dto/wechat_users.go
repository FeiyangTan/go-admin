package dto

import (
	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type WechatUsersGetPageReq struct {
	dto.Pagination `search:"-"`
	OpenId         string `form:"openId"  search:"type:exact;column:open_id;table:wechat_users" comment:"微信 OpenID"`
	NickName       string `form:"nickName"  search:"type:exact;column:nick_name;table:wechat_users" comment:"昵称"`
	WechatUsersOrder
}

type WechatUsersOrder struct {
	Id          string `form:"idOrder"  search:"type:order;column:id;table:wechat_users"`
	OpenId      string `form:"openIdOrder"  search:"type:order;column:open_id;table:wechat_users"`
	NickName    string `form:"nickNameOrder"  search:"type:order;column:nick_name;table:wechat_users"`
	AvatarUrl   string `form:"avatarUrlOrder"  search:"type:order;column:avatar_url;table:wechat_users"`
	PhoneNumber string `form:"phoneNumberOrder"  search:"type:order;column:phone_number;table:wechat_users"`
	FaceCount   string `form:"faceCountOrder"  search:"type:order;column:face_count;table:wechat_users"`
	TongueCount string `form:"tongueCountOrder"  search:"type:order;column:tongue_count;table:wechat_users"`
	CreatedAt   string `form:"createdAtOrder"  search:"type:order;column:created_at;table:wechat_users"`
	UpdatedAt   string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:wechat_users"`
	DeletedAt   string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:wechat_users"`
}

func (m *WechatUsersGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type WechatUsersInsertReq struct {
	Id          int    `json:"-" comment:"自增主键"` // 自增主键
	OpenId      string `json:"openId" comment:"微信 OpenID"`
	NickName    string `json:"nickName" comment:"昵称"`
	AvatarUrl   string `json:"avatarUrl" comment:"头像 URL"`
	PhoneNumber string `json:"phoneNumber" comment:"手机号码"`
	FaceCount   string `json:"faceCount" comment:"人脸计数"`
	TongueCount string `json:"tongueCount" comment:"舌头计数"`
	common.ControlBy
}

func (s *WechatUsersInsertReq) Generate(model *models.WechatUsers) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.OpenId = s.OpenId
	model.NickName = s.NickName
	model.AvatarUrl = s.AvatarUrl
	model.PhoneNumber = s.PhoneNumber
	model.FaceCount = s.FaceCount
	model.TongueCount = s.TongueCount
}

func (s *WechatUsersInsertReq) GetId() interface{} {
	return s.Id
}

type WechatUsersUpdateReq struct {
	Id          int    `uri:"id" comment:"自增主键"` // 自增主键
	OpenId      string `json:"openId" comment:"微信 OpenID"`
	NickName    string `json:"nickName" comment:"昵称"`
	AvatarUrl   string `json:"avatarUrl" comment:"头像 URL"`
	PhoneNumber string `json:"phoneNumber" comment:"手机号码"`
	FaceCount   string `json:"faceCount" comment:"人脸计数"`
	TongueCount string `json:"tongueCount" comment:"舌头计数"`
	common.ControlBy
}

func (s *WechatUsersUpdateReq) Generate(model *models.WechatUsers) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.OpenId = s.OpenId
	model.NickName = s.NickName
	model.AvatarUrl = s.AvatarUrl
	model.PhoneNumber = s.PhoneNumber
	model.FaceCount = s.FaceCount
	model.TongueCount = s.TongueCount
}

func (s *WechatUsersUpdateReq) GetId() interface{} {
	return s.Id
}

// WechatUsersGetReq 功能获取请求参数
type WechatUsersGetReq struct {
	Id int `uri:"id"`
}

func (s *WechatUsersGetReq) GetId() interface{} {
	return s.Id
}

// WechatUsersDeleteReq 功能删除请求参数
type WechatUsersDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *WechatUsersDeleteReq) GetId() interface{} {
	return s.Ids
}
