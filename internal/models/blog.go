package models

import (
	"time"

	"gorm.io/gorm"
)

type Blog struct {
	gorm.Model
	ID        int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	UserID    int
	User      User
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"` // Automatically managed by GORM
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"` // Automatically managed by GORM
}
