package models

import "gorm.io/gorm"

type PaymentTransaction struct {
	gorm.Model
	PaymentTypeID uint
	BillPaymentID uint
}
