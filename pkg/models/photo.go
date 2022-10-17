package models

type Photo struct {
	Identifier
	Path  string `json:"path"`
	Order int    `json:"order"`

	// belongsTo
	CarID uint `json:"car_id"`
	Car   *Car `json:"car,omitempty"`

	Timestamps
}
