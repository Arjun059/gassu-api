package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID          uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	UserID      int       `json:"userId"`
	User        User      `json:"user" gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"` // Automatically managed by GORM
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"` // Automatically managed by GORM

	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"` // Enables soft delete
}
