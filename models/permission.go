package models

import "gorm.io/gorm"

type Permission struct {
	gorm.Model
	Name        string
	Description string
	MenuID      uint
}
