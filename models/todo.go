package models

import (
	"time"
)

type Todo struct {
	ID          uint `gorm:"primarykey"`
	UserID      uint `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	FinishedAt  time.Time `gorm:"type:date"`
	Title       string    `gorm:"not null"`
	Description string    `gorm:"not null"`
	IsDone      bool      `gorm:"Defualt:false"`
}
