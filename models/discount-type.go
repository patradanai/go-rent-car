package models

import "gorm.io/gorm"

type DisocuntType struct {
	gorm.Model
	Unit     string
	Discount []Discount
}
