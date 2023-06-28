package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type PGConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg PGConfig) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=%v", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.SSLMode))
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		fmt.Print("DB ping error! ")
		return nil, err
	}

	return db, nil
}
