package puregrade

import "time"

type Review struct {
	Id        int       ``
	Title     string    ``
	Body      string    ``
	Author    User      ``
	Product   Product   ``
	Rate      int       ``
	CreatedAt time.Time ``
	UpdatedAt time.Time ``
}

type RewiewFilter struct {
	Limit     int
	Offset    int
	AuthorId  int
	ProductId int
}

type Like struct {
	Id       int  ``
	Value    bool `` // 0 - Dislike, 1 - Like
	ReviewId int  ``
	UserId   int  ``
}
