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
