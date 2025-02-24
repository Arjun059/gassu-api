package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`
	UserType string `json:"userType"`
	Email    string `json:"email"`
	Password string `json:"password"`

	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"` // Automatically managed by GORM
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"`

	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"` // Enables soft delete

	Blogs []Blog `gorm:"foreignKey:UserID"`
}
