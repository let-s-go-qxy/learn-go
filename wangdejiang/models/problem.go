package models

import "gorm.io/gorm"

type Problem struct {
	gorm.Model
	Identity string `gorm:"type:varchar(36)" json:"identity"` // 标识
	// TODO 看看这里的json是做什么的。如果有需要，将CategoryId的json改为cid
	CategoryId string `gorm:"column:cid;type:varchar(255)" json:"category_id"` // 分类id逗号分隔
	Title      string `gorm:"type:varchar(255)" json:"title"`                  // 文章title
	Content    string `gorm:"type:text" json:"content"`                        // 文章内容
	MaxMem     int    `gorm:"type:int(11)" json:"max_mem"`                     // 最大内存
	MaxRuntime int    `gorm:"type:int(11)" json:"max_runtime"`                 // 最大运行时间
	PassNumb   int    `gorm:"type:int(11)" json:"pass_numb"`                   // 问题通过数
	TotalNumb  int    `gorm:"type:int(11)" json:"total_numb"`                  // 问题总数
}

// TableName 获取表格名称
func (receiver *Problem) TableName() string {
	return "problem"
}
