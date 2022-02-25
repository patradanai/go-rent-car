package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserInfo     UserInfo
	RefreshToken []RefreshToken
	Booking      []Booking
	BillPayment  []BillPayment
	Username     string
	Password     string
	Email        string
	Status       bool
}
