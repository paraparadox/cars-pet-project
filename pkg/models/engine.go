package models

type Engine struct {
	Identifier
	Type              string `binding:"required" json:"type"`
	NumberOfCylinders int    `binding:"required,numeric,min=1" json:"number_of_cylinders"`
	FuelType          string `binding:"required" json:"fuel_type"`

	// todo: complete belongsTo
	CarID uint

	Timestamps
}
