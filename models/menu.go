package models

import "gorm.io/gorm"

type Menu struct {
	gorm.Model
	Name        string
	Action      string
	Description string
	Permission  []Permission
}
