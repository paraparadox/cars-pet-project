package models

type Engine struct {
	Identifier
	Type              string
	NumberOfCylinders int
	FuelType          string

	// todo: complete belongsTo
	CarID uint

	Timestamps
}
