package models

import "gorm.io/gorm"

type Photo struct {
	gorm.Model
	Path  string
	Order int

	// belongsTo
	CarID uint
	Car   Car
}
