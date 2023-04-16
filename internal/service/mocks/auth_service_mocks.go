package service_mocks

import (
	"errors"
	"fmt"

	puregrade "github.com/ZaiPeeKann/auth-service_pg"
	"github.com/dgrijalva/jwt-go"
)

type AuthService struct {
	something string
}

func NewAuthService() *AuthService {
	return &AuthService{something: "i don't know"}
}

func (s *AuthService) CreateUser(user puregrade.User) (int, error) {
	return 1, nil
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	var token *jwt.Token = jwt.NewWithClaims(jwt.SigningMethodHS256, struct {
		userId int
		jwt.StandardClaims
	}{
		1,
		jwt.StandardClaims{
			ExpiresAt: 60 * 60 * 1000, // 1h
		},
	})

	return token.SignedString([]byte("jwtSecretString"))
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, struct {
		userId int
		jwt.StandardClaims
	}{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected singing method: %v", t.Header["alg"])
		}
		return "jwtSecretKey", nil
	})

	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(*struct {
		userId int
		jwt.StandardClaims
	}); ok && token.Valid {
		return claims.userId, nil
	}

	return 0, errors.New("invalid access token")
}
