package models

// Used to compile all the dynamic data for the HTML
type MainPage struct {
	Categories []Category
	Posts      []PostData
	Trending   []PostData
	Metadata   Metadata
	Pagination Pagination
}

type PostPage struct {
	Post     PostData
	Comments []CommentData
	Metadata Metadata
}

type Metadata struct {
	Success  string
	Error    string
	LoggedIn bool
}

type Pagination struct {
	CurrentPage int
	TotalPages  int
}
