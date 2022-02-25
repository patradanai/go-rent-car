package models

import "gorm.io/gorm"

type DiscountType struct {
	gorm.Model
	Unit     string
	Discount []Discount
}
