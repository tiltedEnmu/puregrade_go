package repository

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/ZaiPeeKann/puregrade"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

// limit posts in GetAll method
const Limit int = 20

type ProductPostgres struct {
	db *sqlx.DB
}

func NewProductPostgres(db *sqlx.DB) *ProductPostgres {
	return &ProductPostgres{db: db}
}

func (r *ProductPostgres) Create(product puregrade.CreateProductDTO) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	defer tx.Rollback()

	var id int
	createProductQuery := "insert into products (title, body, release_date, created_at) values ($1, $2, $3, $4) returning id"
	row := tx.QueryRow(createProductQuery, product.Title, product.Body, product.ReleaseDate, time.Now())
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	createProductGenresQuery := "insert into products_genres (product_id, genre_id) values ($1, $2)"
	for _, value := range product.Genres {
		_, err = tx.Exec(createProductGenresQuery, id, value)
		if err != nil {
			return 0, err
		}
	}

	createProductPlatformsQuery := "insert into product_platforms (product_id, platform_id) values ($1, $2)"
	for _, value := range product.Platforms {
		_, err = tx.Exec(createProductPlatformsQuery, id, value)
		if err != nil {
			return 0, err
		}
	}

	if err = tx.Commit(); err != nil {
		return 0, err
	}

	return id, err
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
	var query string = `select p.id,
						p.title,
						p.body,
						p.release_date,
						p.created_at,
						array(
							select pl.name from products_platforms as pp
							join platforms as pl on pp.platform_id = pl.id
							where p.id = pp.product_id
						) as platforms,
						array(
							select g.name from products_genres as pg
							join genres as g on pg.genre_id = g.id
							where p.id = pg.product_id
						) as genres
						from products as p
						where p.id = $1`
	err := r.db.QueryRow(query, id).Scan(&product.Id, &product.Title, &product.Body, &product.ReleaseDate, &product.CreatedAt, pq.Array(&product.Platforms), pq.Array(&product.Genres))
	fmt.Println(err)
	fmt.Println(product)
	return product, err
}

func (r *ProductPostgres) DeleteGenres(id int, g []int) error {
	var sb strings.Builder
	var q string = `delete from products_genres where product_id = $1 and genre_id in(`
	sb.WriteString(q)
	for i, v := range g {
		if i < len(g)-1 {
			sb.WriteString(fmt.Sprint(v, ", "))
		}
		sb.WriteString(fmt.Sprint(v, ")"))
	}
	_, err := r.db.Exec(sb.String(), id)
	return err
}

func (r *ProductPostgres) DeletePlatforms(id int, p []int) error {
	var sb strings.Builder
	var q string = `delete from products_platforms where product_id = $1 and platform_id in(`
	sb.WriteString(q)
	for i, v := range p {
		if i < len(p)-1 {
			sb.WriteString(fmt.Sprint(v, ", "))
		}
		sb.WriteString(fmt.Sprint(v, ")"))
	}
	_, err := r.db.Exec(sb.String(), id)
	return err
}

func (r *ProductPostgres) AddGenres(id int, g []int) error {
	var sb strings.Builder
	var q string = `insert into products_genres (product_id, genre_id) values`
	sb.WriteString(q)
	for _, v := range g {
		sb.WriteString(fmt.Sprint(" (", id, ", ", v, "),"))
	}
	_, err := r.db.Exec(sb.String()[:sb.Len()-1])
	return err
}

func (r *ProductPostgres) AddPlatforms(id int, p []int) error {
	var sb strings.Builder
	var q string = `insert into products_platforms (product_id, platform_id) values`
	sb.WriteString(q)
	for _, v := range p {
		sb.WriteString(fmt.Sprint(" (", id, ", ", v, "),"))
	}
	_, err := r.db.Exec(sb.String()[:sb.Len()-1])
	return err
}

func (r *ProductPostgres) Delete(id int) error {
	var query string = `delete from products where id = $1`
	_, err := r.db.Query(query, id)
	return err
}
