package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/ZaiPeeKann/puregrade"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

var DefaultUserRoleId int64 = 0

func (r *UserPostgres) Create(user puregrade.User) (int64, error) {
	if len(user.Roles) == 0 {
		user.Roles = append(user.Roles, DefaultUserRoleId)
	}

	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	defer tx.Rollback()

	var id int64
	createUserQuery := "insert into users (email, username, password, avatar, created_at) values ($1, $2, $3, $4, $5) returning id"

	row := tx.QueryRow(createUserQuery, user.Email, user.Username, user.Password, user.Avatar, time.Now())
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	createUserRoleQuery := "insert into users_roles (user_id, role_id) values ($1, $2)"
	for _, value := range user.Roles {
		_, err = tx.Exec(createUserRoleQuery, id, value)
		if err != nil {
			return 0, err
		}
	}

	if err = tx.Commit(); err != nil {
		return 0, err
	}

	return id, err
}

func (r *UserPostgres) Get(username string) (puregrade.User, error) {
	var user puregrade.User
	var q string = `select users.id, users.email, users.avatar, users.banned, users.ban_reason, users.status,
					array(
						select uf.follower_id from users_followers as uf
						where users.id = uf.publisher_id
					) as followers,
					array(
						select ur.role_id from users_roles as ur
						where users.id = ur.user_id
					) as roles
					from users  
					where users.username = $1`
	err := r.db.QueryRow(q, username).Scan(&user.Id, &user.Email, &user.Avatar, &user.Banned,
		&user.BanReason, &user.Status, pq.Array(&user.Followers), pq.Array(&user.Roles))
	fmt.Print(user.Id)
	fmt.Print(user.Roles)
	if err == sql.ErrNoRows {
		return user, errors.New("invalid username")
	}
	return user, err
}

func (r *UserPostgres) GetById(id int64) (puregrade.Profile, error) {
	var p puregrade.Profile
	var q string = `select users.id, users.username, users.avatar, users.status,
					array(
						select uf.follower_id from users_followers as uf
						where users.id = uf.publisher_id
					) as followers,
					array(
						select ur.role_id from users_roles as ur
						where users.id = ur.user_id
					) as roles,
					users.created_at
					from users  
					where users.id = $1`
	err := r.db.QueryRow(q, id).Scan(&p.Id, &p.Username, &p.Avatar, &p.Status, pq.Array(&p.Followers), pq.Array(&p.Roles), &p.CreatedAt)
	if err == sql.ErrNoRows {
		return p, errors.New("invalid password or username")
	}
	return p, err
}

func (r *UserPostgres) CheckUserRole(id, role int64) (bool, error) {
	var ok bool
	var q string = `select exists (
		select id from users_roles where user_id = $1 and role_id = $2
	)`
	err := r.db.Get(&ok, q, id, role)
	return ok, err
}

func (r *UserPostgres) AddFollower(id, publisherId int64) error {
	var q string = `insert into users_followers (follower_id, publisher_id) values ($1, $2)`
	_, err := r.db.Exec(q, id, publisherId)
	return err
}

func (r *UserPostgres) DeleteFollower(id, publisherId int64) error {
	var q string = `delete from users_followers where follower_id = $1 and publisher_id = $2`
	_, err := r.db.Exec(q, id, publisherId)
	return err
}

func (r *UserPostgres) Delete(id int64, password string) error {
	var q string = `delete from users where id = $1 and password = $2`
	_, err := r.db.Exec(q, id, password)
	return err
}
