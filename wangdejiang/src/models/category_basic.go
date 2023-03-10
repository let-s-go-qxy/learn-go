package models

import "gorm.io/gorm"

type CategoryBasic struct {
	gorm.Model
	Identity string `gorm:"type:varchar(36)" json:"identity"` // 标识
	Name     string `gorm:"type:varchar(32)" json:"name"`
	ParentId int    `gorm:"type:int(11)" json:"parent_id"`
}

func (receiver *CategoryBasic) TableName() string {
	return "category_basic"
}
