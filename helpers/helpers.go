package helpers

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	"gitea.koodsisu.fi/josepfrancescbenetmorella/literary-lions/dbaser"
	"gitea.koodsisu.fi/josepfrancescbenetmorella/literary-lions/models"
)

var funcMap = template.FuncMap{
	"sub": func(a, b int) int {
		return a - b
	},
	"add": func(a, b int) int {
		return a + b
	},
}

func RenderTemplate(w http.ResponseWriter, name string, data interface{}) {
	htmlTemplates := []string{
		"web/templates/base.html",
		"web/templates/header.html",
		"web/templates/sidebar.html",
		"web/templates/breadcrumbs.html",
		"web/templates/filter.html",
		"web/templates/main-gallery.html",
		"web/templates/post-templates.html",
		"web/templates/pagination.html",
	}

	//	Adding to the Templates the needed html page to be sent for each specific page request.
	htmlTemplates = append(htmlTemplates, "web/templates/"+name+".html")

	tmpl := template.Must(template.New("base.html").Funcs(funcMap).ParseFiles(htmlTemplates...))

	err := tmpl.ExecuteTemplate(w, "base.html", data)
	if err != nil {
		log.Println("ERROR EXECUTING TEMPLATES")
		log.Printf("Error Executing Template: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func GetPostData(db *sql.DB, post models.Post, userId int) (models.PostData, error) {
	postUser, err := dbaser.UserById(db, post.UserId)
	if err != nil {
		return models.PostData{}, err
	}
	comments, err := dbaser.PostComments(db, post.Id)
	if err != nil {
		return models.PostData{}, err
	}
	likes, dislikes, err := dbaser.PostReactions(db, post.Id)
	if err != nil {
		return models.PostData{}, err
	}
	categories, err := dbaser.PostCategories(db, post.Id)
	if err != nil {
		return models.PostData{}, err
	}
	var likeStatus int
	if userId == 0 {
		likeStatus = 0
	} else {
		likeStatus, err = dbaser.PostLikeStatus(db, post.Id, userId)
		if err != nil {
			return models.PostData{}, err
		}
	}
	data := models.PostData{
		Post:         post,
		User:         postUser,
		Categories:   categories,
		Comments:     comments,
		LikeCount:    likes,
		DislikeCount: dislikes,
		Liked:        likeStatus,
	}
	return data, nil
}

func GetCommentData(db *sql.DB, comment models.Comment, sessionUser int) (models.CommentData, error) {
	commentUser, err := dbaser.UserById(db, comment.UserId)
	if err != nil {
		return models.CommentData{}, err
	}
	likes, dislikes, err := dbaser.CommentReactions(db, comment.Id)
	if err != nil {
		return models.CommentData{}, err
	}
	var likeStatus int
	if sessionUser == 0 {
		likeStatus = 0
	} else {
		likeStatus, err = dbaser.CommentLikeStatus(db, comment.Id, sessionUser)
		if err != nil {
			return models.CommentData{}, err
		}
	}
	data := models.CommentData{
		Comment:      comment,
		User:         commentUser,
		LikeCount:    likes,
		DislikeCount: dislikes,
		Liked:        likeStatus,
	}
	return data, nil
}

func MainPageData(db *sql.DB, userId, page int) (models.MainPage, error) {
	var mainData models.MainPage
	posts, err := dbaser.MainPagePosts(db, page)
	if err != nil {
		log.Print(err)
		mainData.Metadata.Error = err.Error()
		return mainData, err
	}
	var postData []models.PostData
	for _, p := range posts {
		data, err := GetPostData(db, p, userId)
		if err != nil {
			mainData.Metadata.Error = err.Error()
			return mainData, err
		}
		postData = append(postData, data)
	}

	trending, err := dbaser.TrendingPosts(db, 3)
	if err != nil {
		mainData.Metadata.Error = err.Error()
		return mainData, err
	}

	var trendData []models.PostData
	for _, p := range trending {
		data, err := GetPostData(db, p, userId)
		if err != nil {
			mainData.Metadata.Error = err.Error()
			return mainData, err
		}
		trendData = append(trendData, data)
	}
	categories, err := dbaser.Categories(db)
	if err != nil {
		mainData.Metadata.Error = err.Error()
		return mainData, err
	}
	loggedIn, err := dbaser.ValidSession(db, userId)
	if err != nil {
		mainData.Metadata.Error = err.Error()
		return mainData, err
	}
	pagination, err := NumberOfPages(db)
	if err != nil {
		mainData.Metadata.Error = err.Error()
		return mainData, err
	}
	metadata := models.Metadata{LoggedIn: loggedIn}
	mainData = models.MainPage{Categories: categories, Posts: postData, Trending: trendData, Metadata: metadata, Pagination: pagination}
	return mainData, nil
}

func PostPageData(db *sql.DB, postId, sessionUser int) (models.PostPage, error) {
	post, err := dbaser.PostById(db, postId)
	if err != nil {
		return models.PostPage{}, err
	}
	data, err := GetPostData(db, post, sessionUser)
	if err != nil {
		return models.PostPage{}, err
	}
	var comments []models.CommentData
	for _, comment := range data.Comments {
		commData, err := GetCommentData(db, comment, sessionUser)
		if err != nil {
			return models.PostPage{}, err
		}
		comments = append(comments, commData)
	}
	loggedIn, err := dbaser.ValidSession(db, sessionUser)
	if err != nil {
		return models.PostPage{}, err
	}
	metadata := models.Metadata{LoggedIn: loggedIn}
	postData := models.PostPage{Post: data, Comments: comments, Metadata: metadata}
	return postData, nil
}

func NumberOfPages(db *sql.DB) ([]int, error) {
	nPosts, err := dbaser.NumberOfPosts(db)
	if err != nil {
		return []int{}, err
	}
	var quot, rest int
	quot = nPosts / 5
	rest = nPosts % 5
	if quot == 0 || (quot == 1 && rest == 0) {
		return []int{1}, nil
	}
	var pagination []int
	if rest != 0 {
		quot++
	}
	for i := 1; i < quot+1; i++ {
		pagination = append(pagination, i)
	}
	return pagination, nil
}

func CreatePostData(db *sql.DB, id int) (models.MainPage, error) {
	categories, err := dbaser.Categories(db)
	if err != nil {
		return models.MainPage{}, err
	}

	loggedIn, err := dbaser.ValidSession(db, id)
	if err != nil {
		return models.MainPage{}, err
	}
	metadata := models.Metadata{LoggedIn: loggedIn}

	postData := models.MainPage{Categories: categories, Metadata: metadata}
	return postData, nil
}
