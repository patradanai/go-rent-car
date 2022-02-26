package repositories

import (
	"car-booking/models"

	"gorm.io/gorm"
)

type IUserRepository interface {
	FindOneByUser(username string) (*models.User, error)
}

type DatabaseRepository struct {
	DB *gorm.DB
}

func UserRepository(db *gorm.DB) IUserRepository {
	return &DatabaseRepository{DB: db}
}

func (c *DatabaseRepository) FindOneByUser(username string) (*models.User, error) {
	user := models.User{}

	if result := c.DB.First(&user, "username = ?", username); result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (c *DatabaseRepository) CreateUser(userModel models.User) (*models.User, error) {
	if result := c.DB.Create(&userModel); result.Error != nil {
		return nil, result.Error
	}
	return &userModel, nil
}
