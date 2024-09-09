package models

import (
	"fmt"
	"time"
)

type User struct {
	Id       int
	Email    string
	Name     string
	Password string
	Created  time.Time
	Avatar   string
}

type Post struct {
	Id      int
	UserId  int
	Title   string
	Content string
	Created time.Time
}

// Liked will be either -1, 0 or 1, representing disliked, neutral or liked, respectively.
type PostReaction struct {
	PostId int
	UserId int
	Liked  int
}

// PostData holds the metadata for a particular post.
type PostData struct {
	Post       Post
	User       User
	Categories []Category
	LikeCount  int
	Comments   []Comment // Use len for number of comments.
	Liked      int
}

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
	Liked     int
}

type CommentData struct {
	Comment      Comment
	User         User
	LikeCount    int
	DislikeCount int
}

type Category struct {
	Id   int
	Name string
}

type Session struct {
	Id        int
	UserId    int
	Uuid      string
	ExpiresAt time.Time
}

func (p Post) String() string {
	return fmt.Sprintf("** %s **\n%s\n%s\n", p.Title, p.Content, p.Created.Format("02-01-2006 15:04"))
}

// Used to compile all the dynamic data for the HTML
type MainData struct {
	Categories  []string
	Posts       []Post
	TotalPages  int
	CurrentPage int
	MiniPosts   []MiniPost
}

// Used to compile all the dynamic data for each minipost in the sidebar
type MiniPost struct {
	Id         int
	User       User
	Title      string
	LikesCount int
}
