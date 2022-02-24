package models

import "gorm.io/gorm"

type Branch struct {
	gorm.Model
	BranchAddressID uint
	Vehicle         []Vehicle
}
