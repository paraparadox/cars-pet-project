package models

type Car struct {
	Identifier
	ModelName         string `binding:"required" json:"model_name"`
	TopSpeed          int    `binding:"required,numeric,min=0" json:"top_speed"`
	ZeroToHundredTime int    `binding:"required,numeric,min=0" json:"zero_to_hundred_time"`
	FuelTankCapacity  int    `binding:"required,numeric,min=0" json:"fuel_tank_capacity"`
	Length            int    `binding:"required,numeric,min=0" json:"length"`
	Width             int    `binding:"required,numeric,min=0" json:"width"`
	Height            int    `binding:"required,numeric,min=0" json:"height"`
	Weight            int    `binding:"required,numeric,min=0" json:"weight"`

	// belongsTo
	ManufacturerID uint          `binding:"omitempty,numeric,min=1" json:"manufacturer_id"`
	Manufacturer   *Manufacturer `json:"manufacturer,omitempty"`

	// hasOne
	Engine *Engine `json:"engine,omitempty"`

	// hasMany
	Photos []Photo `json:"photos,omitempty"`

	Timestamps
}
