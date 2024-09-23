package models

type Category struct {
	Id   int
	Name string
}

type PostCategory struct {
	PostId     int
	CategoryId int
}
