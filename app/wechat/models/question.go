package models

type Question struct {
	ID    uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name  string `gorm:"type:varchar(128);not null;uniqueIndex:uk_wechat_question_name"" json:"name"`
	Value string `gorm:"type:varchar(255)" json:"value"`
}

// TableName 指定表名
func (Question) TableName() string {
	return "wechat_question"
}
