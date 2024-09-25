package models

// Used to compile all the dynamic data for the HTML
type MainPage struct {
	Categories  []Category
	Posts       []PostData
	Trending    []PostData
	LoggedIn    bool
	Pagination  []int
	CurrentPage int
	TotalPages  int
	Metadata    Metadata
}

type PostPage struct {
	Post     PostData
	Comments []CommentData
	LoggedIn bool
	Metadata Metadata
}

type Metadata struct {
	Success  string
	Error    string
	LoggedIn bool
}
