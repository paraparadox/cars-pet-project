package models

import (
	"gorm.io/gorm"
	"time"
)

type ID struct {
	ID uint `json:"id" gorm:"primary_key"`
}

type Timestamps struct {
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
