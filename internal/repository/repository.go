package repository

import (
	puregrade "github.com/ZaiPeeKann/auth-service_pg"
	"github.com/jmoiron/sqlx"
)

type User interface {
	CreateUser(puregrade.User) (int, error)
	GetUser(username, password string) (puregrade.User, error)
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
	Create(product puregrade.Product) (int, error)
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
