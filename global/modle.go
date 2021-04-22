package global

import (
	"gorm.io/gorm"
	"time"
)

// 数据结构基本项
type GVA_MODEL struct {
	ID          uint `gorm:"primarykey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedTime gorm.DeletedAt `gorm:"index" json:"-"`
}
