package repository

import (
	"fmt"
	"time"

	"github.com/ZaiPeeKann/puregrade"
	"github.com/jmoiron/sqlx"
)

type ReviewPostgres struct {
	db *sqlx.DB
}

func NewReviewPostgres(db *sqlx.DB) *ReviewPostgres {
	return &ReviewPostgres{db: db}
}

func (r *ReviewPostgres) Create(review puregrade.Review) (int, error) {
	var query string = `insert into reviews (body, author_id, product_id, rate, created_at, updated_at)
						values ($1, $2, $3, $4, $5, $6) returning id`
	var id int
	row := r.db.QueryRow(query, review.Title, review.Body, review.Author.Id, review.Product.Id, review.Rate, time.Now())
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *ReviewPostgres) GetAll(page int, productId int) ([]puregrade.Review, error) {
	var reviews []puregrade.Review
	var query string = `select * from reviews
						inner join reviews_products as p on p.review_id = reviews.id
						inner join products on p.product_id = products.id`
	if productId != 0 {
		query += fmt.Sprintf("where products.id = %d", productId)
	}
	query += fmt.Sprintf("limit %d offset %d", Limit, Limit*(page-1))
	err := r.db.Select(&reviews, query)
	return reviews, err
}

func (r *ReviewPostgres) GetOneByID(id int) (puregrade.Review, error) {
	var review puregrade.Review
	var query string = `select * from reviews
						inner join reviews_products as p on p.review_id = reviews.id
						inner join products on p.product_id = products.id
						where id = $1`
	err := r.db.Select(&review, query, id)
	return review, err
}

func (r *ReviewPostgres) Update(id int, title, body string) error {
	var query string = `update reviews set title = $1, body = $2 where id = $3`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *ReviewPostgres) Delete(id int) error {
	var query string = `delete from reviews where id = $1`
	_, err := r.db.Exec(query, id)
	return err
}
