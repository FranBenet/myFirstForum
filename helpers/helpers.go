package helpers

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"gitea.koodsisu.fi/josepfrancescbenetmorella/literary-lions/dbaser"
	"gitea.koodsisu.fi/josepfrancescbenetmorella/literary-lions/models"
)

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

	tmpl := template.Must(template.ParseFiles(htmlTemplates...))

	err := tmpl.ExecuteTemplate(w, "base.html", data)
	if err != nil {
		fmt.Println("ERROR EXECUTING TEMPLATES")
		log.Printf("Error Executing Template: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

/* Mindmap of post workflow
   To create PostData I'll need a session ID to identify the user requesting, and a post ID.
   With the post ID I have to fetch the post itself, the number of likes and dislikes, all the
   comments as well as reactions to all the comments. For the post and each of the comments I
   have to check if the user requesting has liked or disliked them.

   TODO
   - Add pagination function, 5 posts per page.
   - Add message field to structs in order to pass error messages (login/resgistration, for example).
*/

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

func MainPageData(db *sql.DB, id int) (models.MainPage, error) {
	posts, err := dbaser.Posts(db)
	if err != nil {
		log.Print(err)
		return models.MainPage{}, err
	}

	var postData []models.PostData
	for _, p := range posts {
		data, err := GetPostData(db, p, id)
		if err != nil {
			return models.MainPage{}, err
		}
		postData = append(postData, data)
	}

	trending, err := dbaser.TrendingPosts(db, 3)
	if err != nil {
		return models.MainPage{}, err
	}

	var trendData []models.PostData
	for _, p := range trending {
		data, err := GetPostData(db, p, id)
		if err != nil {
			return models.MainPage{}, err
		}
		trendData = append(trendData, data)
	}

	categories, err := dbaser.Categories(db)
	if err != nil {
		return models.MainPage{}, err
	}

	loggedIn, err := dbaser.ValidSession(db, id)
	if err != nil {
		return models.MainPage{}, err
	}

	mainData := models.MainPage{Categories: categories, Posts: postData, Trending: trendData, LoggedIn: loggedIn}
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
	postData := models.PostPage{Post: data, Comments: comments, LoggedIn: loggedIn}
	return postData, nil
}
