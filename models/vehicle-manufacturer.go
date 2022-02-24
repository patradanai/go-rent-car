package models

import "gorm.io/gorm"

type VehicleManufacturer struct {
	gorm.Model
	Name        string
	Description string
	Vehicle     []Vehicle
}
