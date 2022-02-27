package repositories

import (
	"car-booking/models"

	"gorm.io/gorm"
)

type IUserRepository interface {
	FindOneByUser(username string) (*models.User, error)
	FindById(id uint) (*models.User, error)
	CreateOne(userModel *models.User) (*models.User, error)
}

type UserRepository struct {
	DatabaseRepository
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{
		DatabaseRepository{DB: db},
	}
}

func (c *UserRepository) FindOneByUser(username string) (*models.User, error) {
	user := models.User{}

	if err := c.Debug().Preload("Roles").Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (c *UserRepository) FindById(id uint) (*models.User, error) {
	user := models.User{}

	if err := c.Model(&user).Where("id = ?", id).Association("Roles").Find(&models.Role{}); err != nil {
		return nil, err
	}

	return &user, nil
}

func (c *UserRepository) CreateOne(userModel *models.User) (*models.User, error) {
	if result := c.Create(userModel); result.Error != nil {
		return nil, result.Error
	}
	return userModel, nil
}
