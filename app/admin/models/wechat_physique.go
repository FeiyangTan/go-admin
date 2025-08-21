package models

import (
	"go-admin/common/models"
)

type WechatPhysique struct {
	models.Model

	PhysiqueName      string `json:"physiqueName" gorm:"type:varchar(128);comment:体质名称"`
	AcupunctureMethod string `json:"acupunctureMethod" gorm:"type:varchar(2000);comment:针灸法"`
	ProductIds        string `json:"productIds" gorm:"type:text;comment:商品ID数组（以逗号分隔的字符串）"`
	models.ModelTime
	models.ControlBy
}

func (WechatPhysique) TableName() string {
	return "wechat_physique"
}

func (e *WechatPhysique) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *WechatPhysique) GetId() interface{} {
	return e.Id
}
