package models

import "gorm.io/gorm"

type BillPayment struct {
	gorm.Model
	PaymentTransaction []PaymentTransaction
	BookingID          uint
	Status             string
}
