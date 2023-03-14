package todo

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	ID          uint           `gorm:"primarykey"`
	CreatedAt   time.Time      `gorm:"datetime:timestamp;autoCreateTime"`
	UpdatedAt   time.Time      `gorm:"datetime:timestamp;autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Title       string         `gorm:"not null"`
	Category    string
	Description string
}
