package models

import "gorm.io/gorm"

type Engine struct {
	gorm.Model
	Type              string
	NumberOfCylinders int
	FuelType          string

	// todo: complete belongsTo
	CarID uint
}
