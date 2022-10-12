package models

import "gorm.io/gorm"

type Car struct {
	gorm.Model
	ModelName         string
	TopSpeed          int
	ZeroToHundredTime int
	FuelTankCapacity  int
	Length            int
	Width             int
	Height            int
	Weight            int
}
