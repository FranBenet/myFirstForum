package models

// Used to compile all the dynamic data for the HTML
type MainPage struct {
<<<<<<< HEAD
	Categories  []Category
	Posts       []PostData
	Trending    []PostData
	LoggedIn    bool
	Pagination  []int
	CurrentPage int
	TotalPages  int
	Metadata    Metadata
=======
	Categories []Category
	Posts      []PostData
	Trending   []PostData
	Pagination []int
	Metadata   Metadata
>>>>>>> 5ecbcd320c8e4fdd2b70b2627a97c65423f572f9
}

type PostPage struct {
	Post     PostData
	Comments []CommentData
<<<<<<< HEAD
	LoggedIn bool
=======
>>>>>>> 5ecbcd320c8e4fdd2b70b2627a97c65423f572f9
	Metadata Metadata
}

type Metadata struct {
	Success  string
	Error    string
	LoggedIn bool
}
