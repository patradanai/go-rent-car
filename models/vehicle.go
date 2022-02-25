package models

import "gorm.io/gorm"

type Vehicle struct {
	gorm.Model
	VehicleID             uint
	VehicleTypeID         uint
	VehicleManufacturerID uint
	BranchID              uint
	Booking               Booking
	RegisterNumber        string
	Slug                  string
	Color                 string
	Year                  string
	Price                 string
	Status                bool
}
