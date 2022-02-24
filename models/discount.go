package models

import (
	"time"

	"gorm.io/gorm"
)

type Discount struct {
	gorm.Model
	DiscountTypeID       uint
	DiscountAmount       uint
	CouponCode           string
	MinimunOrderValue    uint
	MaximunDiscountValue uint
	ExpiredAt            time.Time
}
