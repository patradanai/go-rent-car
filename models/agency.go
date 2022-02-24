package models

import "gorm.io/gorm"

type Agency struct {
	gorm.Model
	Name   string
	Branch []Branch
}
