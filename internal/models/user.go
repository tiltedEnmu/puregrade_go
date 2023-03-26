package puregrade

import "time"

type User struct {
	Id        int       `json:"id" db:"id"`
	Username  string    `json:"username" db:"username"`
	Email     string    ``
	Password  string    ``
	Avatar    string    `json:"avatar"`
	Banned    bool      ``
	BanReason string    ``
	Status    string    ``
	Follows   []int     ``
	Roles     []string  ``
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}
