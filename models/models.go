package models

import "fmt"

type User struct {
	Email    string
	Name     string
	Password string
}

type Post struct {
	Id      int
	Created string
	Title   string
	Content string
}

type Category struct {
	Id   int
	Name string
}

func (p Post) String() string {
	return fmt.Sprintf("** %s **\n%s\n%s\n", p.Title, p.Content, p.Created)
}
