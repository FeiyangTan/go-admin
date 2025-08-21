package models

import (
	"time"

	"gorm.io/gorm"
)

// WechatProduct 对应表 wechat_products
type Product struct {
	ID            uint64         `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	ProductName   string         `gorm:"size:128;not null;index:idx_mall_products_name;column:product_name" json:"product_name"` // 商品名称
	ImageURL      string         `gorm:"size:512;not null;column:image_url" json:"image_url"`                                    // 商品图片 URL
	Description   string         `gorm:"type:varchar(2000);column:description" json:"description"`                               // 产品介绍
	Ingredients   string         `gorm:"type:varchar(2000);column:ingredients" json:"ingredients"`                               // 产品成分
	UsageMethod   string         `gorm:"type:varchar(2000);column:usage_method" json:"usage_method"`                             // 使用方法
	Price         float64        `gorm:"type:decimal(10,2);not null;default:0.00;column:price" json:"price"`                     // 商品价格
	MallProductID uint64         `gorm:"not null;column:mall_product_id" json:"mall_product_id"`                                 // 商城商品编号
	CreatedAt     time.Time      `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index:idx_mall_products_deleted_at;column:deleted_at" json:"deleted_at,omitempty"`
}

// TableName 覆盖默认表名
func (Product) TableName() string {
	return "wechat_products"
}
