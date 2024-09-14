package models

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
