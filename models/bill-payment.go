package models

import "gorm.io/gorm"

type BillPayment struct {
	gorm.Model
	PaymentTransaction []PaymentTransaction
	UserID             uint
	BookingID          uint
	Status             string
	PromotionCode      string
	Discount           uint
	SubTotal           uint
	Total              uint
}
