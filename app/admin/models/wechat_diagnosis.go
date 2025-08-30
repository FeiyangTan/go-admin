package models

import (
	"go-admin/common/models"
)

type WechatDiagnosis struct {
	models.Model

	OpenId             string `json:"openId" gorm:"type:varchar(64);comment:微信昵称"`
	DiagnosisType      string `json:"diagnosisType" gorm:"type:varchar(64);comment:诊疗类型"`
	DiagnosisPhysique1 string `json:"diagnosisPhysique1" gorm:"type:varchar(32);comment:体质1名称"`
	DiagnosisPhysique2 string `json:"diagnosisPhysique2" gorm:"type:varchar(32);comment:体质2名称"`
	DiagnosisResult    string `json:"diagnosisResult" gorm:"type:longtext;comment:诊疗结果 JSON 字符串"`
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
