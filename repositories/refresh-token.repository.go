package repositories

import (
	"car-booking/models"

	"gorm.io/gorm"
)

type IRefreshToken interface {
	CreateRecord(entity *models.RefreshToken) (*models.RefreshToken, error)
	UpdateRecord(id uint, entity *models.RefreshToken) (*models.RefreshToken, error)
}

type RefreshTokenRepository struct {
	DatabaseRepository
}

func NewRefreshTokenRepository(db *gorm.DB) IRefreshToken {
	return &RefreshTokenRepository{
		DatabaseRepository{DB: db},
	}
}

func (c *RefreshTokenRepository) CreateRecord(entity *models.RefreshToken) (*models.RefreshToken, error) {
	if result := c.Create(entity); result.Error != nil {
		return nil, result.Error
	}

	return entity, nil
}

func (c *RefreshTokenRepository) UpdateRecord(id uint, entity *models.RefreshToken) (*models.RefreshToken, error) {
	if err := c.Model(&models.RefreshToken{}).Where("id = ?", id).Updates(entity).Error; err != nil {
		return nil, err
	}

	return entity, nil
}
