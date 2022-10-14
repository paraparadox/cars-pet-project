package models

type Manufacturer struct {
	Identifier
	Title          string `binding:"required" json:"title"`
	FoundationYear int    `binding:"required,numeric,min=1" json:"foundation_year"`
	Logo           string `binding:"required,url" json:"logo"`

	// hasMany
	Cars []Car `json:"cars,omitempty"`

	Timestamps
}
