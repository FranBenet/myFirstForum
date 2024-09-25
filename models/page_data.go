package models

// Used to compile all the dynamic data for the HTML
type MainPage struct {
	Categories  []Category
	Posts       []PostData
	Trending    []PostData
	Pagination  []int
	Metadata    Metadata
	CurrentPage int
	TotalPages  int
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
