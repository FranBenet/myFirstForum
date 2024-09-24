package models

// Used to compile all the dynamic data for the HTML
type MainPage struct {
	Categories []Category
	Posts      []PostData
	Trending   []PostData
	LoggedIn   bool
	Pagination []int
	Metadata   Metadata
}

type PostPage struct {
	Post     PostData
	Comments []CommentData
	LoggedIn bool
}

type Metadata struct {
	Success  string
	Error    string
	LoggedIn bool
}
