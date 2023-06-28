package repository

import (
	"fmt"
	"strings"
	"time"

	"github.com/ZaiPeeKann/puregrade"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type ProductPostgres struct {
	db *sqlx.DB
}

func NewProductPostgres(db *sqlx.DB) *ProductPostgres {
	return &ProductPostgres{db: db}
}

func (r *ProductPostgres) Create(product puregrade.CreateProductDTO) (int64, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	defer tx.Rollback()

	var id int64
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

func (r *ProductPostgres) GetAll(offset, limit int, genres, platforms []int64, orderBy string, isAsc bool) ([]puregrade.Product, error) {
	var products []puregrade.Product
	var query string = `select * from products
						inner join products_platforms as p on p.product_id = products.id
						inner join products_genres as g on g.product_id = products.id`
	if genres != nil || platforms != nil {
		query += "where "
	}

	if genres != nil {
		for _, v := range genres {
			query += fmt.Sprintf("g.id = %d and ", v)
		}
	}

	if platforms != nil {
		for _, v := range genres {
			query += fmt.Sprintf("p.id = %d and ", v)
		}
		query = query[:len(query)-5] // delete last "and"
	} else {
		query = query[:len(query)-5]
	}

	if limit > 0 {
		query += fmt.Sprintf(" limit %d", limit)
	}
	if offset > 0 {
		query += fmt.Sprintf(" offset %d", offset)
	}

	if orderBy != "" {
		query += fmt.Sprintf(" order by", orderBy)
		if isAsc == true {
			query += " asc"
		} else {
			query += " desc"
		}
	}

	err := r.db.Select(&products, query)

	return products, err
}

func (r *ProductPostgres) GetOneByID(id int64) (puregrade.Product, error) {
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

func (r *ProductPostgres) DeleteGenres(id int64, g []int64) error {
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

func (r *ProductPostgres) DeletePlatforms(id int64, p []int64) error {
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

func (r *ProductPostgres) AddGenres(id int64, g []int64) error {
	var sb strings.Builder
	var q string = `insert into products_genres (product_id, genre_id) values`
	sb.WriteString(q)
	for _, v := range g {
		sb.WriteString(fmt.Sprint(" (", id, ", ", v, "),"))
	}
	_, err := r.db.Exec(sb.String()[:sb.Len()-1])
	return err
}

func (r *ProductPostgres) AddPlatforms(id int64, p []int64) error {
	var sb strings.Builder
	var q string = `insert into products_platforms (product_id, platform_id) values`
	sb.WriteString(q)
	for _, v := range p {
		sb.WriteString(fmt.Sprint(" (", id, ", ", v, "),"))
	}
	_, err := r.db.Exec(sb.String()[:sb.Len()-1])
	return err
}

func (r *ProductPostgres) Delete(id int64) error {
	var query string = `delete from products where id = $1`
	_, err := r.db.Query(query, id)
	return err
}
