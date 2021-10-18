package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        string         `json:"id" gorm:"type: uuid; primaryKey"`
	Name      string         `json:"name" binding:"required" gorm:"type:varchar(100); not null"`
	Email     string         `json:"email" binding:"required" gorm:"type:varchar(100); not null"`
	Password  string         `json:"password" binding:"required" gorm:"type:varchar; not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type UserWithOutPassword struct {
	ID        string         `json:"id"`
	Name      string         `json:"name" binding:"required"`
	Email     string         `json:"email" binding:"required"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
