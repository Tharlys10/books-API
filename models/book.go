package models

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID          string         `json:"id" gorm:"type:uuid; primaryKey"`
	Name        string         `json:"name" binding:"required" gorm:"type:varchar(100); not null"`
	Description string         `json:"description" binding:"required" gorm:"not null"`
	Author      string         `json:"author" binding:"required" gorm:"type:varchar(100); not null"`
	ImageUrl    string         `json:"image_url"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
