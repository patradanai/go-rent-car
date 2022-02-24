package models

import "gorm.io/gorm"

type BranchAddress struct {
	gorm.Model
	AddressLine1 string
	AddressLine2 string
	Street       string
	City         string
	Country      string
	PostCode     string
	PhoneNumber  string
	Branch       Branch
}
