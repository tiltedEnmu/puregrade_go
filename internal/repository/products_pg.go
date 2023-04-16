package repository

import (
	"errors"
	"fmt"
	"time"

	puregrade "github.com/ZaiPeeKann/auth-service_pg"
	"github.com/jmoiron/sqlx"
)

// limit posts in GetAll method
const Limit int = 20

type ProductPostgres struct {
	db *sqlx.DB
}

func NewProductPostgres(db *sqlx.DB) *ProductPostgres {
	return &ProductPostgres{db: db}
}

func (r *ProductPostgres) Create(product puregrade.Product) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createProductQuery := "insert into products (title, body, release_date, created_at) values ($1, $2, $3, $4) returning id"
	row := tx.QueryRow(createProductQuery, product.Title, product.Body, product.ReleaseDate, time.Now())
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	createProductGenresQuery := "insert into product_genres (product_id, genre_id) values ($1, $2)"
	for _, value := range product.Genres {
		_, err = tx.Exec(createProductGenresQuery, id, value)
		if err != nil {
			tx.Rollback()
			return 0, err
		}
	}

	createProductPlatformsQuery := "insert into product_platforms (product_id, platform_id) values ($1, $2)"
	for _, value := range product.Platforms {
		_, err = tx.Exec(createProductPlatformsQuery, id, value)
		if err != nil {
			tx.Rollback()
			return 0, err
		}
	}

	return id, tx.Commit()
}

func (r *ProductPostgres) GetAll(page int, filter map[string]string) ([]puregrade.Product, error) {
	var products []puregrade.Product
	var query string = `select * from products
						inner join products_platforms as p on p.product_id = products.id
						inner join platforms on p.platform_id = platforms.id
						inner join products_genres as g on g.product_id = products.id
						inner join genres on g.genre_id = genres.id`
	if page <= 0 {
		return products, errors.New("page must not be negative")
	}
	if len(filter) != 0 {
		query += "where "
		for k, v := range filter {
			query += k + " = " + v + " and "
		}
		query = query[:len(query)-5]
	}

	query += fmt.Sprintf("limit %d offset %d", Limit, Limit*(page-1))

	err := r.db.Select(&products, query)
	return products, err
}

func (r *ProductPostgres) GetOneByID(id int) (puregrade.Product, error) {
	var product puregrade.Product
	var query string = `select * from products
						inner join products_platforms as p on p.product_id = products.id
						inner join platforms on p.platform_id = platforms.id
						inner join products_genres as g on g.product_id = products.id
						inner join genres on g.genre_id = genres.id
						where products.id = $1`
	err := r.db.Select(&product, query, id)
	return product, err
}

func (r *ProductPostgres) Delete(id int) error {
	var query string = `delete from products where id = $1`
	_, err := r.db.Query(query, id)
	return err
}
