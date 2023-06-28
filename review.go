package puregrade

import "time"

type Review struct {
	Id        int64     `json:"id" db:"id"`
	Title     string    `json:"title" db:"title"`
	Body      string    `json:"body" db:"body"` // markdown text
	Author    User      `json:"author" db:"author"`
	Product   Product   `json:"product" db:"product"`
	Rate      int       `json:"rate" db:"rate"` // 1 - 100
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
}

type RewiewFilter struct {
	Page      int
	ProductId int64
}

type ReviewLike struct {
	Id       int64 `json:"id" db:"id"`
	ReviewId int64 `json:"reviewId" db:"review_id"`
	UserId   int64 `json:"userId" db:"user_id"`
}
