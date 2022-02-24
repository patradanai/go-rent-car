package models

import (
	"time"

	"gorm.io/gorm"
)

type Booking struct {
	gorm.Model
	UserID      uint
	VehicleID   uint
	BillPayment []BillPayment
	BookingAt   time.Time
	ReturnAt    time.Time
	DropPoint   string
	PickUpPoint string
	Status      string
}
