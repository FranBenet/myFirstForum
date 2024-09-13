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

type PostReaction struct {
	PostId int
	UserId int
	Liked  bool
}

// PostData holds the metadata for a particular post.
// Liked tells if the user requesting the data has liked a specific post or comment.
// Liked will be either -1, 0 or 1, representing disliked, neutral or liked, respectively.
type PostData struct {
	Post         Post
	User         User
	Categories   []Category
	LikeCount    int
	DislikeCount int
	Comments     []Comment // Use len for number of comments.
	Liked        int
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
	Liked     bool
}

type CommentData struct {
	Comment      Comment
	User         User
	LikeCount    int
	DislikeCount int
	Liked        int
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

func (pd PostData) String() string {
	return fmt.Sprintf(
		"\n%vUser: %v\nCategories: %v\nLikes: %d Dislikes: %d Number of comments: %d Like status: %d\nComments: %v\n",
		pd.Post, pd.User, pd.Categories, pd.LikeCount, pd.DislikeCount, len(pd.Comments), pd.Liked, pd.Comments,
	)
}

// Used to compile all the dynamic data for the HTML
type MainPage struct {
	Categories []Category
	Posts      []PostData
	Trending   []PostData
	LoggedIn   bool
}

type PostPage struct {
	Post     PostData
	Comments []CommentData
	LoggedIn bool
}
