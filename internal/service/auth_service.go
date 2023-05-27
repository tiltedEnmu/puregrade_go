package service

import (
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"fmt"
	"time"

	"github.com/ZaiPeeKann/puregrade"
	"github.com/ZaiPeeKann/puregrade/internal/repository"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/spf13/viper"
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

var jwtSecretKey string = viper.GetString("jwtsecretkey")

func (s *AuthService) CreateUser(user puregrade.User) (int, error) {
	user.Banned = false
	user.CreatedAt = time.Now()
	user.Password = generatePasswordHash(user.Password)
	user.BanReason = ""
	user.Avatar = uuid.New().String()
	return s.repos.User.Create(user)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	fmt.Print(generatePasswordHash(password))
	user, err := s.repos.User.Get(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}
	var token *jwt.Token = jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims{
		user.Id,
		jwt.StandardClaims{
			ExpiresAt: 24 * 60 * 60 * 1000 * 1000000, // 24h
		},
	})
	return token.SignedString([]byte(jwtSecretKey))
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &jwtClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected singing method: %v", t.Header["alg"])
		}
		return []byte(jwtSecretKey), nil
	})

	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(*jwtClaims); ok && token.Valid {
		return claims.UserId, nil
	}

	return 0, errors.New("invalid access token")
}

func generatePasswordHash(password string) string {
	bytes := sha1.New()
	bytes.Write([]byte(password))
	return hex.EncodeToString(bytes.Sum(nil))
}
