package service

import (
	"github.com/ZaiPeeKann/puregrade"
	"github.com/ZaiPeeKann/puregrade/internal/repository"
)

type Authorization interface {
	CreateUser(puregrade.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type User interface {
	GetProfile(id int) (puregrade.Profile, error)
	Delete(id int, password string) error
	FollowUser(id, publisherId int) error
	UnfollowUser(id, publisherId int) error
}

type Review interface {
	GetAll(page int, productId int) ([]puregrade.Review, error)
	GetOneByID(id int) (puregrade.Review, error)
	Create(review puregrade.Review) (int, error)
	Update(id int, title, body string) error
	Delete(id, userId int) error
}

type Product interface {
	GetAll(page int, filter map[string]string) ([]puregrade.Product, error)
	GetOneByID(id int) (puregrade.Product, error)
	Create(product puregrade.CreateProductDTO) (int, error)
	AddGenres(id int, g []int) error
	AddPlatforms(id int, p []int) error
	DeleteGenres(id int, g []int) error
	DeletePlatforms(id int, p []int) error
	Delete(id int) error
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
