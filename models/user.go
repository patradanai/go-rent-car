package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserInfo     UserInfo
	RefreshToken []RefreshToken
	Booking      []Booking
	BillPayment  []BillPayment
	Role         []*Role `gorm:"many2many:user_roles;"`
	Username     string  `json:"username"`
	Password     string  `json:"password"`
	Email        string  `json:"email"`
	Status       bool    `json:"status"`
}
