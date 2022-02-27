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
	FindUserById(id uint) (*models.User, error)
	CreateUser(user *models.User) (*models.User, error)
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

func (r *userService) FindUserById(id uint) (*models.User, error) {
	user, err := r.repository.FindById(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userService) CreateUser(user *models.User) (*models.User, error) {
	user, err := r.repository.CreateOne(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
