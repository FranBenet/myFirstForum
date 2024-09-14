package models

import "time"

type Session struct {
	Id        int
	UserId    int
	Uuid      string
	ExpiresAt time.Time
}
