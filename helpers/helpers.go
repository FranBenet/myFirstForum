package helpers

import (
	"html/template"
	"literary-lions/model"
	"log"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, name string, data *model.Data) {
	htmlTemplates := []string{
		"web/templates/base.html",
		"web/templates/header.html",
		"web/templates/home.html",
		"web/templates/posts-section.html",
		"web/templates/post-templates.html",
		"web/templates/post-view.html",
		"web/templates/breadcrumbs.html",
		"web/templates/sidebar.html",
		"web/templates/login.html",
		"web/templates/register.html",
		"web/templates/user-profile.html",
		"web/templates/notifications.html",
		"web/templates/post-create.html",
	}

	tmpl, err := template.ParseFiles(htmlTemplates...)
	if err != nil {
		log.Println("Error Parsing Template: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, name, data)
	if err != nil {
		log.Println("Error Executing Template: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// THIS FUNCTIONS HAS TO BE DELETED. IT'S JUST TO CREATE A SIMPLE DATA FOR DEVELOPING
func GetData() *model.Data {

	post01 := &model.Post{
		Id:            1,
		UserId:        1,
		Username:      "HistoricalFictionBuff",
		Category:      "Fantasy",
		Title:         "Lord Of the rings, everything about the trilogy that changed the world",
		Content:       "flksjadogaj ogsalj goaigj aog jaogj oajsgj aoligjao jgiodsja ogja jgoajo jgdsa ogjaoj",
		PostCreated:   "16 March 2024",
		LikesCount:    5,
		CommentsCount: 3,
		LikeState:     false,
		DislikeState:  false,
	}

	post02 := &model.Post{
		Id:            2,
		UserId:        2,
		Username:      "BookEater",
		Category:      "War",
		Title:         "The Impact of Setting on 'All the Light We Cannot See",
		Content:       "The settings in All the Light We Cannot See—particularly the war-torn cities of France—play such a crucial role in shaping the narrative. The way Anthony Doerr describes Saint-Malo almost makes it feel like a character itself.",
		PostCreated:   "21 November 2022",
		LikesCount:    9,
		CommentsCount: 9,
		LikeState:     true,
		DislikeState:  false,
	}

	minipost01 := &model.MiniPost{
		Id:         1,
		Username:   "BookLover90",
		Title:      "Exploring the Depths of Character Development in 'To Kill a Mockingbird",
		LikesCount: 10,
	}

	minipost02 := &model.MiniPost{
		Id:         2,
		Username:   "FantasyFantastic",
		Title:      "Comparing Magic Systems: Sanderson vs. Rowling",
		LikesCount: 7,
	}

	var postsCollection []model.Post
	postsCollection = append(postsCollection, *post01)
	postsCollection = append(postsCollection, *post02)

	var minipostsCollection []model.MiniPost
	minipostsCollection = append(minipostsCollection, *minipost01)
	minipostsCollection = append(minipostsCollection, *minipost02)

	categories := []string{"Fantasy", "War", "Fiction", "Non-fiction", "Romance", "Crime"}

	data := &model.Data{
		Categories: categories,
		Posts:      postsCollection,
		MiniPosts:  minipostsCollection,
	}

	return data
}

// THIS FUNCTIONS HAS TO BE DELETED. IT'S JUST TO CREATE A SIMPLE DATA FOR DEVELOPING
func getPostId() *model.Data {
	post01 := &model.Post{
		Id:            1,
		UserId:        1,
		Username:      "HistoricalFictionBuff",
		Category:      "Fantasy",
		Title:         "Lord Of the rings, everything about the trilogy that changed the world",
		Content:       "flksjadogaj ogsalj goaigj aog jaogj oajsgj aoligjao jgiodsja ogja jgoajo jgdsa ogjaoj",
		PostCreated:   "16 March 2024",
		LikesCount:    5,
		CommentsCount: 3,
		LikeState:     false,
		DislikeState:  false,
		Avatar:        "meerkat.png",
	}

	var postsCollection []model.Post
	postsCollection = append(postsCollection, *post01)

	data := &model.Data{
		Posts: postsCollection,
	}
	return data
}
