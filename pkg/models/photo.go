package models

import "gorm.io/gorm"

type Photo struct {
	gorm.Model
	Path  string
	Order int
}
