package helpers

import (
	"database/sql"
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
		log.Printf("Error Executing Template: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// THIS FUNCTIONS HAS TO BE DELETED. IT'S JUST TO CREATE A SIMPLE DATA FOR DEVELOPING
// func GetDataExample() *models.MainData {
// 	user01 := &models.User{
// 		Id:     1,
// 		Name:   "HistoricalFictionBuff",
// 		Avatar: "meerkat.png",
// 	}
// 	category01 := &models.Category{
// 		Name: "Fantasy",
// 	}
// 	category02 := &models.Category{
// 		Name: "Epic",
// 	}
// 	category03 := &models.Category{
// 		Name: "Trilogy",
// 	}
// 	var categories01 []models.Category
// 	categories01 = append(categories01, *category01)
// 	categories01 = append(categories01, *category02)
// 	categories01 = append(categories01, *category03)

// 	post01 := &models.Post{
// 		Id:            1,
// 		User:          *user01,
// 		Categories:    categories01,
// 		Title:         "Lord Of the rings, everything about the trilogy that changed the world",
// 		Content:       "flksjadogaj ogsalj goaigj aog jaogj oajsgj aoligjao jgiodsja ogja jgoajo jgdsa ogjaoj",
// 		Created:       "16 March 2024",
// 		LikesCount:    5,
// 		CommentsCount: 3,
// 		LikeState:     false,
// 		DislikeState:  false,
// 	}

// 	user02 := &models.User{
// 		Id:     2,
// 		Name:   "BookEater",
// 		Avatar: "bear.png",
// 	}
// 	category04 := &models.Category{
// 		Name: "War",
// 	}
// 	category05 := &models.Category{
// 		Name: "Shocking",
// 	}
// 	category06 := &models.Category{
// 		Name: "Romantic",
// 	}
// 	var categories02 []models.Category
// 	categories02 = append(categories02, *category04)
// 	categories02 = append(categories02, *category05)
// 	categories02 = append(categories02, *category06)

// 	post02 := &models.Post{
// 		Id:            2,
// 		User:          *user02,
// 		Categories:    categories02,
// 		Title:         "The Impact of Setting on 'All the Light We Cannot See",
// 		Content:       "The settings in All the Light We Cannot See—particularly the war-torn cities of France—play such a crucial role in shaping the narrative. The way Anthony Doerr describes Saint-Malo almost makes it feel like a character itself.Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.",
// 		Created:       "21 November 2022",
// 		LikesCount:    9,
// 		CommentsCount: 9,
// 		LikeState:     true,
// 		DislikeState:  false,
// 	}

// 	user03 := &models.User{
// 		Name: "BookLover90",
// 	}

// 	minipost01 := &models.MiniPost{
// 		Id:         1,
// 		User:       *user03,
// 		Title:      "Exploring the Depths of Character Development in 'To Kill a Mockingbird",
// 		LikesCount: 10,
// 	}

// 	user04 := &models.User{
// 		Name: "MarkKing",
// 	}
// 	minipost02 := &models.MiniPost{
// 		Id:         2,
// 		User:       *user04,
// 		Title:      "Comparing Magic Systems: Sanderson vs. Rowling",
// 		LikesCount: 7,
// 	}

// 	var postsCollection []models.Post
// 	postsCollection = append(postsCollection, *post01)
// 	postsCollection = append(postsCollection, *post02)

// 	var minipostsCollection []models.MiniPost
// 	minipostsCollection = append(minipostsCollection, *minipost01)
// 	minipostsCollection = append(minipostsCollection, *minipost02)

// 	categories := []string{"Fantasy", "War", "Fiction", "Non-fiction", "Romance", "Crime"}

// 	data := &models.MainData{
// 		Categories: categories,
// 		Posts:      postsCollection,
// 		MiniPosts:  minipostsCollection,
// 	}

// 	return data
// }

// // THIS FUNCTIONS HAS TO BE DELETED. IT'S JUST TO CREATE A SIMPLE DATA FOR DEVELOPING
// func GetPostIdExample() *models.Post {
// 	user03 := &models.User{
// 		Id:     1,
// 		Name:   "HistoricalFictionBuff",
// 		Avatar: "meerkat.png",
// 	}

// 	category01 := &models.Category{
// 		Name: "Fantasy",
// 	}

// 	category02 := &models.Category{
// 		Name: "Epic",
// 	}
// 	category03 := &models.Category{
// 		Name: "Trilogy",
// 	}

// 	var categories03 []models.Category
// 	categories03 = append(categories03, *category01)
// 	categories03 = append(categories03, *category02)
// 	categories03 = append(categories03, *category03)

// 	user01 := &models.User{
// 		Id:     34,
// 		Name:   "MarkKing",
// 		Avatar: "dog2.png",
// 	}

// 	comment01 := &models.Comment{
// 		Id:           1,
// 		Content:      "I've been re-reading To Kill a Mockingbird, and I'm amazed by how Harper Lee crafted such complex and relatable characters. Especially Atticus Finch—his moral compass and calm demeanor really resonate with me. How do you think Lee's portrayal of these characters reflects the social issues of the time? Also, who is your favorite character and why? Would love to hear your thoughts! /n I've been re-reading To Kill a Mockingbird, and I'm amazed by how Harper Lee crafted such complex and relatable characters. Especially Atticus Finch—his moral compass and calm demeanor really resonate with me. How do you think Lee's portrayal of these characters reflects the social issues of the time? Also, who is your favorite character and why? Would love to hear your thoughts!",
// 		Created:      "23 March 2026",
// 		User:         *user01,
// 		LikesCount:   21,
// 		LikeState:    true,
// 		DislikeState: false,
// 	}

// 	user02 := &models.User{
// 		Id:     38,
// 		Name:   "DeboraBooks",
// 		Avatar: "rabbit.png",
// 	}

// 	comment02 := &models.Comment{
// 		Id:           2,
// 		Content:      "SHIT SHIT SHIT shit shit shit what ever text may seem important here!",
// 		Created:      "2 April 2026",
// 		User:         *user02,
// 		LikesCount:   2,
// 		LikeState:    false,
// 		DislikeState: true,
// 	}

// 	var commentsCollection []models.Comment
// 	commentsCollection = append(commentsCollection, *comment01)
// 	commentsCollection = append(commentsCollection, *comment02)

// 	post01 := &models.Post{
// 		Id:            1,
// 		Categories:    categories03,
// 		Title:         "Lord Of the rings, everything about the trilogy that changed the world",
// 		Content:       "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.",
// 		Created:       "16 March 2024",
// 		User:          *user03,
// 		LikesCount:    5,
// 		CommentsCount: 3,
// 		LikeState:     false,
// 		DislikeState:  false,
// 		Comments:      commentsCollection,
// 	}

// 	return post01
// }

/* Mindmap of post workflow
   To create PostData I'll need a session ID to identify the user requesting, and a post ID.
   With the post ID I have to fetch the post itself, the number of likes and dislikes, all the
   comments as well as reactions to all the comments. For the post and each of the comments I
   have to check if the user requesting has liked or disliked them.
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
		Comment:      &comment,
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
