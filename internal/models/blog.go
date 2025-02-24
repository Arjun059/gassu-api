package models

import (
	"time"

	"gorm.io/gorm"
)

type Blog struct {
	ID      uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Title   string `json:"title"`
	Content string `json:"content"`
	UserID  uint   `json:"userId"`
	User    *User  `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`

	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"` // Automatically managed by GORM
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"` // Automatically managed by GORM

	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"` // Enables soft delete
}
