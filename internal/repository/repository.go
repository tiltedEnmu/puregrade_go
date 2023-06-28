package repository

import (
	"github.com/ZaiPeeKann/puregrade"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

type Auth interface {
	UpsertRefreshToken(userId, token string) error
	GetRefreshToken(userId string) (string, error)
}

type User interface {
	Create(user puregrade.User) (int64, error)
	Get(username string) (puregrade.User, error)
	GetById(id int64) (puregrade.Profile, error)
	AddFollower(id, publisherId int64) error
	DeleteFollower(id, publisherId int64) error
	Delete(id int64, password string) error
}

type Review interface {
	GetAll(page int, productId int64) ([]puregrade.Review, error)
	GetOneByID(id int64) (puregrade.Review, error)
	Create(review puregrade.Review) (int64, error)
	Update(id int64, title, body string) error
	Delete(id int64) error
}

type Product interface {
	GetAll(offset, limit int, genres, platforms []int64, orderBy string, isAsc bool) ([]puregrade.Product, error)
	GetOneByID(id int64) (puregrade.Product, error)
	Create(product puregrade.CreateProductDTO) (int64, error)
	AddGenres(id int64, genres []int64) error
	AddPlatforms(id int64, platforms []int64) error
	DeleteGenres(id int64, genres []int64) error
	DeletePlatforms(id int64, platforms []int64) error
	Delete(id int64) error
}

type Repository struct {
	Auth
	User
	Review
	Product
}

type Databases struct {
	Redis    *redis.Client
	Postgres *sqlx.DB
}

func NewRepository(dbs *Databases) *Repository {
	return &Repository{
		Auth:    NewAuthRedis(dbs.Redis),
		User:    NewUserPostgres(dbs.Postgres),
		Review:  NewReviewPostgres(dbs.Postgres),
		Product: NewProductPostgres(dbs.Postgres),
	}
}
