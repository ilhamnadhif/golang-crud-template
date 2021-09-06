package entity

import "time"

type Post struct {
	Id       int32
	Title    string
	Slug     string
	Body     string
	Image    string
	CreateAt time.Time
}
