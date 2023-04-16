package puregrade

import "time"

type User struct {
	Id        int       `json:"id" db:"id"`
	Username  string    `json:"username" db:"username" binding:"required"`
	Email     string    `json:"email" db:"email" binding:"required"`
	Password  string    `json:"password" db:"password" binding:"required"`
	Avatar    string    `json:"avatar" db:"avatar"`
	Banned    bool      `json:"banned" db:"banned"`
	BanReason string    `json:"banReason" db:"ban_reason"`
	Status    string    `json:"status" db:"status"`
	Follows   []int     `json:"follows"`
	Roles     []string  `json:"roles"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}
