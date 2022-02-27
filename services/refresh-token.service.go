package services

import (
	"car-booking/models"
	"car-booking/repositories"
	"time"
)

type IRefreshTokenService interface {
	CreateRefreshToken(userId uint, token string, expiredAt time.Time, revoke bool) (*models.RefreshToken, error)
}

type refreshTokenService struct {
	repository repositories.IRefreshToken
}

func NewRefreshTokenService(r repositories.IRefreshToken) IRefreshTokenService {
	return &refreshTokenService{repository: r}
}

func (r *refreshTokenService) CreateRefreshToken(userId uint, token string, expiredAt time.Time, revoke bool) (*models.RefreshToken, error) {
	refreshToken := models.RefreshToken{UserID: userId, Token: token, ExpiredAt: expiredAt, Revoke: revoke}

	result, err := r.repository.CreateRecord(&refreshToken)
	if err != nil {
		return nil, err
	}

	return result, nil
}
