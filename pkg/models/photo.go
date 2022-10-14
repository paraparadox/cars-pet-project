package models

type Photo struct {
	Identifier
	Path  string
	Order int

	// belongsTo
	CarID uint
	Car   Car

	Timestamps
}
