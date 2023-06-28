package service

import (
	"github.com/ZaiPeeKann/puregrade"
	"github.com/ZaiPeeKann/puregrade/internal/repository"
)

type Authorization interface {
	CreateUser(puregrade.User) (int64, error)
	GenerateTokens(username, password string) (string, string, error) // access, refresh, err
	ParseAccessToken(token string) (int64, error)
}

type User interface {
	GetProfile(id int64) (puregrade.Profile, error)
	Delete(id int64, password string) error
	FollowUser(id, publisherId int64) error
	UnfollowUser(id, publisherId int64) error
}

type Review interface {
	GetAll(page int, productId int64) ([]puregrade.Review, error)
	GetOneByID(id int64) (puregrade.Review, error)
	Create(review puregrade.Review) (int64, error)
	Update(id int64, title, body string) error
	Delete(id, userId int64) error
}

type Product interface {
	GetAll(filter puregrade.ProductFilter) ([]puregrade.Product, error)
	GetOneByID(id int64) (puregrade.Product, error)
	Create(product puregrade.CreateProductDTO) (int64, error)
	AddGenres(id int64, g []int64) error
	AddPlatforms(id int64, p []int64) error
	DeleteGenres(id int64, g []int64) error
	DeletePlatforms(id int64, p []int64) error
	Delete(id int64) error
}

type Service struct {
	Authorization
	User
	Review
	Product
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos),
		User:          NewUserService(repos),
		Review:        NewReviewService(repos),
		Product:       NewProductService(repos),
	}
}
