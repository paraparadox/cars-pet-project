package models

type Engine struct {
	ID
	Type              string
	NumberOfCylinders int
	FuelType          string

	// todo: complete belongsTo
	CarID uint

	Timestamps
}
