package models

import "time"

type Comment struct {
	Id      int
	PostId  int
	UserId  int
	Content string
	Created time.Time
}

type CommentReaction struct {
	CommentId int
	UserId    int
	Liked     bool
}

type CommentData struct {
	Comment      Comment
	User         User
	LikeCount    int
	DislikeCount int
	Liked        int
}

func (c Comment) Date() string {
	return c.Created.Format("02-01-2006 15:04")
}
