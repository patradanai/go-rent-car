package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name        string
	Description string
	Users       []*User      `gorm:"many2many:user_roles;"`
	Permission  []Permission `gorm:"many2many:role_permissions;"`
}
