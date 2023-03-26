package repository

import (
	"time"

	puregrade "github.com/ZaiPeeKann/auth-service_pg/internal/models"
	"github.com/jmoiron/sqlx"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

var DefaultUserRoleId int = 0

func (r *UserPostgres) CreateUser(user puregrade.User) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createUserQuery := "insert into users (email, username, password, avatar, created_at) values ($1, $2, $3, $4, $5) returning id"

	row := tx.QueryRow(createUserQuery, user.Email, user.Username, user.Password, user.Avatar, time.Now())
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUserRoleQuery := "insert into users_roles (user_id, role_id) values ($1, $2)"
	_, err = tx.Exec(createUserRoleQuery, id, DefaultUserRoleId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *UserPostgres) GetUser(username, password string) (puregrade.User, error) {
	var user puregrade.User
	query := "select * from users inner join users_follows as f on users.user_id = f.follower_id inner join users_roles as r on users.user_id = r.role_id where users.username = $1 and users.password = $2"
	err := r.db.Select(&user, query, username, password)
	return user, err
}
