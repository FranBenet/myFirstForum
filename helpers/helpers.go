package helpers

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	"gitea.koodsisu.fi/josepfrancescbenetmorella/literary-lions/dbaser"
	"gitea.koodsisu.fi/josepfrancescbenetmorella/literary-lions/models"
)

// Set of functions used within the html templates.
var funcMap = template.FuncMap{
	"sub": func(a, b int) int {
		return a - b
	},
	"add": func(a, b int) int {
		return a + b
	},
	"seq": func(start, end int) []int {
		s := make([]int, end-start+1)
		for i := start; i <= end; i++ {
			s[i-start] = i
		}
		return s
	},
	"comp": func(a string) string {
		switch a {
		case "/myposts":
			return "My Posts"
		case "/liked":
			return "My liked Posts"
		case "/profile":
			return "Profile"
		case "/post/create":
			return "Create Post"
		case "/post/":
			return "Post"
		default:
			return "Current Page"
		}
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

	//	funcMap is a set of functions that are embedd in the htmls to be used within html.
	tmpl := template.Must(template.New("base.html").Funcs(funcMap).ParseFiles(htmlTemplates...))

	err := tmpl.ExecuteTemplate(w, "base.html", data)
	if err != nil {
		log.Printf("Error Executing Template: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// GetPostData retrieves additional data related to a post, such as author data (models.User), comments, categories,
// likes, dislikes, and if the user requesting has reacted to the post.
func GetPostData(db *sql.DB, post models.Post, sessionUser int) (models.PostData, error) {
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
	if sessionUser == 0 {
		likeStatus = 0
	} else {
		likeStatus, err = dbaser.PostLikeStatus(db, post.Id, sessionUser)
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

// GetCommentData retrieves additional data related to a comment, such as author, likes, dislikes, and if
// the user requesting has reacted to the comment.
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

func NumberOfPages(nPosts int) int {
	var quot, rest int
	quot = nPosts / 5
	rest = nPosts % 5
	// if quot == 0 || (quot == 1 && rest == 0) {
	// 	return 0
	// }
	if rest != 0 {
		quot++
	}
	return quot
}

func PostSlice(total, currentPage int) (int, int) {
	end := currentPage * 5
	start := (currentPage - 1) * 5
	if end > total {
		end = total
	}
	return start, end
}

func CreatePostData(db *sql.DB, userId int) (models.MainPage, error) {
	var createPostData models.MainPage
	categories, err := dbaser.Categories(db)
	if err != nil {
		return models.MainPage{}, err
	}
	if userId > 0 {
		createPostData.Metadata.LoggedIn = true
		userData, err := dbaser.UserById(db, userId)
		if err != nil {
			createPostData.Metadata.Error = err.Error()
			return createPostData, err
		}
		user := models.User{Avatar: userData.Avatar}
		createPostData.User = user
	}
	createPostData.Categories = categories
	return createPostData, nil
}
