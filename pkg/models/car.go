package models

type Car struct {
	ID
	ModelName         string
	TopSpeed          int
	ZeroToHundredTime int
	FuelTankCapacity  int
	Length            int
	Width             int
	Height            int
	Weight            int

	// belongsTo
	ManufacturerID uint
	Manufacturer   Manufacturer

	// hasOne
	Engine Engine

	// hasMany
	Photos []Photo

	Timestamps
}
