package models

import (
	"go-admin/common/models"
)

type WechatProducts struct {
	models.Model

	ProductName   string `json:"productName" gorm:"type:varchar(128);comment:商品名称"`
	ImageUrl      string `json:"imageUrl" gorm:"type:varchar(512);comment:商品图片 URL"`
	Description   string `json:"description" gorm:"type:varchar(2000);comment:产品介绍"`
	Ingredients   string `json:"ingredients" gorm:"type:varchar(2000);comment:产品成分"`
	UsageMethod   string `json:"usageMethod" gorm:"type:varchar(2000);comment:使用方法"`
	Price         string `json:"price" gorm:"type:decimal(10,2);comment:商品价格（单位：元）"`
	MallProductId string `json:"mallProductId" gorm:"type:bigint(20);comment:商城商品编号"`
	models.ModelTime
	models.ControlBy
}

func (WechatProducts) TableName() string {
	return "wechat_products"
}

func (e *WechatProducts) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *WechatProducts) GetId() interface{} {
	return e.Id
}
