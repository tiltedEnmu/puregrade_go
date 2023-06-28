package service_mocks

import (
	"errors"
	"fmt"
	"time"

	"github.com/ZaiPeeKann/puregrade"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/spf13/viper"
)

type AuthService struct {
	something string
}

type jwtClaims struct {
	UserId int64 `json:"user_id"`
	jwt.StandardClaims
}

var jwtSecretKey string = viper.GetString("jwtsecretkey")

func NewAuthService() *AuthService {
	return &AuthService{something: "i don't know"}
}

func (s *AuthService) CreateUser(user puregrade.User) (int64, error) {
	return 1, nil
}

func (s *AuthService) GenerateTokens(username, password string) (string, string, error) {

	user := new(puregrade.User)
	user.Password = "mySuperSecretPass123"

	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims{
		user.Id,
		jwt.StandardClaims{
			ExpiresAt: int64(24 * time.Hour),
		},
	}).SignedString([]byte(jwtSecretKey))

	var refreshToken string = uuid.New().String()

	return accessToken, refreshToken, err
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
