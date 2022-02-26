package services

import "car-booking/utils"

type IJWTService interface {
	ValidateToken() (bool, error)
	GenerateToken() (string, error)
}

type jwtService struct {
	secretKey string
	issue     string
}

func JWTService() IJWTService {
	return &jwtService{
		secretKey: utils.LoadEnv("JWT_SECRET_KEY"),
		issue:     utils.LoadEnv("JWT_ISSUE"),
	}
}

func (t *jwtService) ValidateToken() (bool, error) {
	return false, nil
}

func (t *jwtService) GenerateToken() (string, error) {
	return "", nil
}
