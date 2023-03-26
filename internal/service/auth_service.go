package service

import (
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"fmt"

	puregrade "github.com/ZaiPeeKann/auth-service_pg/internal/models"
	"github.com/ZaiPeeKann/auth-service_pg/internal/repository"
	"github.com/dgrijalva/jwt-go"
)

type AuthService struct {
	repos *repository.Repository
}

func NewAuthService(repos *repository.Repository) *AuthService {
	return &AuthService{repos: repos}
}

type jwtClaims struct {
	UserId int `json:"user_id"`
	jwt.StandardClaims
}

var jwtSecretKey string = "MySecretKey123"

func (s *AuthService) CreateUser(user puregrade.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repos.User.CreateUser(user)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repos.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	var token *jwt.Token = jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims{
		user.Id,
		jwt.StandardClaims{
			ExpiresAt: 60 * 60 * 1000, // 1h
		},
	})

	return token.SignedString([]byte(jwtSecretKey))
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, jwtClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected singing method: %v", t.Header["alg"])
		}
		return jwtSecretKey, nil
	})

	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(*jwtClaims); ok && token.Valid {
		return claims.UserId, nil
	}

	return 0, errors.New("Invalid access token")
}

func generatePasswordHash(password string) string {
	bytes := sha1.New()
	bytes.Write([]byte(password))
	return hex.EncodeToString(bytes.Sum(nil))
}
