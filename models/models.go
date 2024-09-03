package models

import "fmt"

type User struct {
	Id       int
	Email    string
	Name     string
	Password string
}

// Used to compile all the data for each post
type Post struct {
	Id            int
	UserId        int
	Title         string
	Content       string
	Created       string
	Username      string
	Category      string
	LikesCount    int
	CommentsCount int
	LikeState     bool
	DislikeState  bool
	Avatar        string
}

type Category struct {
	Id   int
	Name string
}

func (p Post) String() string {
	return fmt.Sprintf("** %s **\n%s\n%s\n", p.Title, p.Content, p.Created)
}

// Used to compile all the dynamic data for the HTML
type Data struct {
	Categories  []string
	Posts       []Post
	TotalPages  int
	CurrentPage int
	MiniPosts   []MiniPost
}

// Used to compile all the dynamic data for each minipost in the sidebar
type MiniPost struct {
	Id         int
	Username   string
	Title      string
	LikesCount int
}
