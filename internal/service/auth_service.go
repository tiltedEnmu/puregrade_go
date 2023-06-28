package service

import (
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"
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
	UserId int64 `json:"user_id"`
	jwt.StandardClaims
}

var jwtSecretKey string = viper.GetString("jwtsecretkey")

func (s *AuthService) CreateUser(user puregrade.User) (int64, error) {
	user.Banned = false
	user.CreatedAt = time.Now()
	user.Password = generatePasswordHash(user.Password)
	user.BanReason = ""
	user.Avatar = uuid.New().String()
	return s.repos.User.Create(user)
}

func (s *AuthService) ParseAccessToken(token string) (int64, error) {
	parsedToken, err := jwt.ParseWithClaims(token, &jwtClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected singing method: %v", t.Header["alg"])
		}
		return []byte(jwtSecretKey), nil
	})

	if err != nil {
		return 0, err
	}

	if claims, ok := parsedToken.Claims.(*jwtClaims); ok && parsedToken.Valid {
		return claims.UserId, nil
	}

	return 0, errors.New("invalid access token")
}

func (s *AuthService) GenerateTokens(username, password string) (string, string, error) {

	user, err := s.repos.User.Get(username)
	if (err != nil) || (user.Password != password) {
		return "", "", err
	}

	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims{
		user.Id,
		jwt.StandardClaims{
			ExpiresAt: int64(24 * time.Hour),
		},
	}).SignedString([]byte(jwtSecretKey))

	var refreshToken string = uuid.New().String()

	if err = s.repos.Auth.UpsertRefreshToken(strconv.FormatInt(user.Id, 10), refreshToken); err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, err
}

func generatePasswordHash(password string) string {
	bytes := sha1.New()
	bytes.Write([]byte(password))
	return hex.EncodeToString(bytes.Sum(nil))
}
