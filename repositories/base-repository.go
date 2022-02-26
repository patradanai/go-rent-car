package repositories

import "gorm.io/gorm"

type DatabaseRepository struct {
	*gorm.DB
}
