package models

import "gorm.io/gorm"

type PaymentType struct {
	gorm.Model
	PaymentTransaction []PaymentTransaction
}
