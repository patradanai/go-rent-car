package models

import "gorm.io/gorm"

type UserInfo struct {
	gorm.Model
	UserID       uint
	Firstname    string
	Lastname     string
	Email        string
	Phone        string
	ImageSrc     string
	AddressLine1 string
	AddressLine2 string
	PostCode     string
	Country      string
}
