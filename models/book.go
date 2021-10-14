package models

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID          string         `json:"id" gorm:"type:uuid; primaryKey"`
	Name        string         `json:"name" gorm:"type:varchar(100); not null"`
	Description string         `json:"description" gorm:"not null"`
	Author      string         `json:"author" gorm:"type:varchar(100); not null"`
	ImageUrl    string         `json:"image_url"`
	CreatedAt   time.Time      `json:"created_at" gorm:"default: NOW(); not null"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
