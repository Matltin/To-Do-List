package models

import (
	"time"
)

type Todo struct {
	ID          uint      `json:"id" gorm:"primarykey"`
	UserID      uint      `json:"user_id" gorm:"not null"`
	CreatedAt   time.Time `json:"create_at"`
	UpdatedAt   time.Time `json:"update_at"`
	FinishedAt  time.Time `json:"finished_at" gorm:"type:date"`
	Title       string    `json:"title" gorm:"not null"`
	Description string    `json:"description" gorm:"not null"`
	IsDone      bool      `json:"is_done" gorm:"Defualt:false"`
}
