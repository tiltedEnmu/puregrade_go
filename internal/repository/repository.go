package repository

import (
	"github.com/ZaiPeeKann/puregrade"
	"github.com/jmoiron/sqlx"
)

type User interface {
	Create(puregrade.User) (int, error)
	Get(username string) (puregrade.User, error)
	GetById(id int) (puregrade.Profile, error)
	AddFollower(id, publisherId int) error
	DeleteFollower(id, publisherId int) error
	Delete(id int, password string) error
}

type Review interface {
	GetAll(page int, productId int) ([]puregrade.Review, error)
	GetOneByID(id int) (puregrade.Review, error)
	Create(review puregrade.Review) (int, error)
	Update(id int, title, body string) error
	Delete(id int) error
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

type Repository struct {
	User
	Review
	Product
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User:    NewUserPostgres(db),
		Review:  NewReviewPostgres(db),
		Product: NewProductPostgres(db),
	}
}
