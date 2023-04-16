package puregrade

import "time"

type Product struct {
	Id          int       ``
	Title       string    ``
	Body        string    ``
	Genres      []string  ``
	Platforms   []string  ``
	ReleaseDate time.Time ``
	CreatedAt   time.Time ``
}

type ProductFilter struct {
	Page     int
	Genre    string
	Platform string
}
