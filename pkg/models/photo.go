package models

type Photo struct {
	ID
	Path  string
	Order int

	// belongsTo
	CarID uint
	Car   Car

	Timestamps
}
