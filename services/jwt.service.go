package services

import (
	"car-booking/utils"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type IJWTService interface {
	ValidateToken(token string) (*jwt.Token, error)
	GenerateToken(userId uint) (string, error)
}

type jwtService struct {
	secretKey string
	issue     string
}

func NewJWTService() IJWTService {
	return &jwtService{
		secretKey: utils.LoadEnv("JWT_SECRET_KEY"),
		issue:     utils.LoadEnv("JWT_ISSUE"),
	}
}

func (t *jwtService) ValidateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", token.Header["alg"])
		}

		return []byte(t.secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func (t *jwtService) GenerateToken(userId uint) (string, error) {
	atClaim := jwt.MapClaims{}
	atClaim["userId"] = userId
	atClaim["exp"] = time.Now().Add(time.Minute * 15).Unix()
	atClaim["issue"] = t.issue

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaim)

	token, err := at.SignedString([]byte(t.secretKey))
	if err != nil {
		return "", nil
	}

	return token, nil
}
