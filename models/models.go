package models

import "fmt"

type User struct {
	Id       int
	Email    string
	Name     string
	Password string
}

type Post struct {
	Id      int
	UserId  int
	Title   string
	Content string
	Created string
}

type Category struct {
	Id   int
	Name string
}

func (p Post) String() string {
	return fmt.Sprintf("** %s **\n%s\n%s\n", p.Title, p.Content, p.Created)
}
