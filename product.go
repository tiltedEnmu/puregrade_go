package puregrade

import "time"

type Product struct {
	Id          int       `json:"id" db:"id"`
	Title       string    `json:"title" db:"title"`
	Body        string    `json:"body" db:"body"`
	Genres      []string  `json:"genres" db:"genres"`
	Platforms   []string  `json:"platforms" db:"platforms"`
	ReleaseDate time.Time `json:"releaseDate" db:"release_date"`
	CreatedAt   time.Time `db:"created_at"`
}

type CreateProductDTO struct {
	Title       string    `json:"title" db:"title"`
	Body        string    `json:"body" db:"body"`
	Genres      []int     `json:"genres" db:"genres"`
	Platforms   []int     `json:"platforms" db:"platforms"`
	ReleaseDate time.Time `json:"releaseDate" db:"release_date"`
}

type ProductRate struct {
	Id        int `json:"id" db:"id"`
	Rate      int `json:"rate" db:"rate"` // 1 - 100
	ProductId int `json:"productId" db:"product_id"`
	UserId    int `json:"userId" db:"user_id"`
}

type ProductFilter struct {
	Page     int
	Genre    string
	Platform string
}
