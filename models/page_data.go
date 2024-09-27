package models

// Used to compile all the dynamic data for the HTML
type MainPage struct {
	Categories []Category
	Posts      []PostData
	Trending   []PostData
	Metadata   Metadata
	Pagination Pagination
	User       User
}

type PostPage struct {
	Post     PostData
	Comments []CommentData
	Metadata Metadata
	User     User
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

type ProfilePageData struct {
	User     User
	Metadata Metadata
}
