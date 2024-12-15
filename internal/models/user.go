package models

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	UserType  string    `json:"userType"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"` // Automatically managed by GORM
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
