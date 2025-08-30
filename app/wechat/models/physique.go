package models

import (
	"time"

	"gorm.io/gorm"
)

type Physique struct {
	ID                uint64         `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	PhysiqueName      string         `gorm:"size:128;not null;uniqueIndex:uq_physique_name;column:physique_name" json:"physique_name"` // 体质名称
	AcupunctureMethod string         `gorm:"type:varchar(2000);column:acupuncture_method" json:"acupuncture_method"`                   // 针灸法
	EighteenMethod    string         `gorm:"type:varchar(2000);column:eighteen_method" json:"eighteen_method"`                         // 十八宝
	WellnessMethod    string         `gorm:"type:varchar(2000);column:wellness_method" json:"wellness_method"`                         // 日常养生
	ProductIDs        string         `gorm:"type:text;column:product_ids" json:"product_ids"`                                          // 以逗号分隔的商品ID字符串
	CreatedAt         time.Time      `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt         time.Time      `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
	DeletedAt         gorm.DeletedAt `gorm:"index;column:deleted_at" json:"deleted_at,omitempty"`
}

// TableName 绑定实际表名
func (Physique) TableName() string {
	return "wechat_physique"
}
