package models

import (
	"go-admin/common/models"
)

type WechatDiagnosis struct {
	models.Model

	OpenId          string `json:"openId" gorm:"type:varchar(64);comment:微信 OpenID"`
	DiagnosisType   string `json:"diagnosisType" gorm:"type:varchar(64);comment:诊疗类型"`
	DiagnosisResult string `json:"diagnosisResult" gorm:"type:longtext;comment:诊疗结果 JSON 字符串"`
	models.ModelTime
	models.ControlBy
}

func (WechatDiagnosis) TableName() string {
	return "wechat_diagnosis"
}

func (e *WechatDiagnosis) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *WechatDiagnosis) GetId() interface{} {
	return e.Id
}
