package models

import "time"

type User struct {
	Id       int
	Email    string
	Name     string
	Password string
	Created  time.Time
	Avatar   string
}

func (u User) Date() string {
	return u.Created.Format("02-01-2006 15:04")
}
