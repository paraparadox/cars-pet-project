package models

import "gorm.io/gorm"

type Manufacturer struct {
	gorm.Model
	Title          string
	FoundationYear int
	Logo           string

	// hasMany
	Cars []Car
}
