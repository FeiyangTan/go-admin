package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Diagnosis struct {
	gorm.Model
	OpenID          string         `gorm:"size:64;not null;index"`
	User            User           `gorm:"foreignKey:OpenID;references:OpenID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"` // 外键约束：OnUpdate 级联更新，OnDelete 置空或级联删除可按需调整
	DiagnosisType   string         `json:"diagnosis_type" gorm:"size:64;not null;index"`
	DiagnosisResult datatypes.JSON `json:"diagnosis_result" gorm:"type:longtext;serializer:json;not null"`
}

// TableName 指定了模型对应的数据库表名
func (Diagnosis) TableName() string {
	return "wechat_diagnosis"
}
