package models

import "fmt"

type User struct {
	Email    string
	Name     string
	Password string
}

type Post struct {
	Created string
	Content string
}

type Category struct {
	Name string
}

func (p Post) String() string {
	return fmt.Sprintf("%s\n%s\n", p.Content, p.Created)
}
