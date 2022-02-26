package services

import (
	"car-booking/models"
	"car-booking/repositories"
)

type userService struct {
	repository repositories.IUserRepository
}

type IUserService interface {
	FindUser(username string) (*models.User, error)
}

func NewUserService(r repositories.IUserRepository) IUserService {
	return &userService{repository: r}
}

func (r *userService) FindUser(username string) (*models.User, error) {
	user, err := r.repository.FindOneByUser(username)
	if err != nil {
		return nil, err
	}
	return user, nil
}
