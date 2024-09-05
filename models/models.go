package models

import "fmt"

type User struct {
	Id       int
	Email    string
	Name     string
	Password string
	Avatar   string
}

// Used to compile all the data for each post
type Post struct {
	Id            int
	Title         string
	Content       string
	Created       string
	User          User
	Categories    []Category
	LikesCount    int
	CommentsCount int
	LikeState     bool
	DislikeState  bool
	Comments      []Comment
}

type Comment struct {
	Id           int
	Content      string
	Created      string
	User         User
	LikesCount   int
	LikeState    bool
	DislikeState bool
}

type Category struct {
	Id   int
	Name string
}

func (p Post) String() string {
	return fmt.Sprintf("** %s **\n%s\n%s\n", p.Title, p.Content, p.Created)
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
