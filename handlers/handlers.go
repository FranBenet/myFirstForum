package handlers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"gitea.koodsisu.fi/josepfrancescbenetmorella/literary-lions/dbaser"
	"gitea.koodsisu.fi/josepfrancescbenetmorella/literary-lions/helpers"
	"gitea.koodsisu.fi/josepfrancescbenetmorella/literary-lions/middleware"
	"gitea.koodsisu.fi/josepfrancescbenetmorella/literary-lions/models"
)

type Handler struct {
	db *sql.DB
}

func NewHandler(db *sql.DB) *Handler {
	return &Handler{db: db}
}

// To handle "/".
func (h *Handler) Homepage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		log.Println("Homepage")
		log.Println("Error. Path Not Allowed.")
		http.Error(w, "Page Not Found", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
		return
	}

	//	Get userID that is making the request
	// userID := r.Context().Value(models.UserIDKey).(int)

	// ---------------------------------------------------PROVISIONAL CODE FOR TEST----------------------------------------------------------------------------------------
	//	Get cookie from request
	var userID int
	sessionToken, err := r.Cookie("session_token")
	if err != nil {
		userID = 0
		log.Println(err)
	} else {
		//	Get the value of the session from the cookie
		sessionUUID := sessionToken.Value
		userID, err = middleware.IsUserLoggedIn(h.db, sessionUUID)
		if err != nil {
			userID = 0
			log.Println(err)

		}
	}

	// ---------------------------------------------------PROVISIONAL CODE FOR TEST----------------------------------------------------------------------------------------

	data, err := helpers.MainPageData(h.db, userID)
	if err != nil {
		log.Println(err)
	}
	data.LoggedIn = false
	helpers.RenderTemplate(w, "home", data)
}

// To handle "/post/{id}"
func (h *Handler) GetPost(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
		return
	}

	//	Get userID that is making the request
	// userID := r.Context().Value(models.UserIDKey).(int)

	// ---------------------------------------------------PROVISIONAL CODE FOR TEST----------------------------------------------------------------------------------------
	//	Get cookie from request
	var userID int

	sessionToken, err := r.Cookie("session_token")
	if err != nil {
		userID = 0
		log.Println(err)
	} else {
		//	Get the value of the session from the cookie
		sessionUUID := sessionToken.Value
		userID, err = middleware.IsUserLoggedIn(h.db, sessionUUID)
		if err != nil {
			userID = 0
			log.Println(err)
		}
	}

	// ---------------------------------------------------PROVISIONAL CODE FOR TEST END----------------------------------------------------------------------------------------

	path := r.URL.Path
	pathDivide := strings.Split(path, "/")

	if len(pathDivide) == 3 && pathDivide[1] == "post" && pathDivide[0] == "" {

		postId, err := strconv.Atoi(pathDivide[2])
		if err != nil {
			http.Error(w, "ID for the post is not a number", http.StatusBadRequest)
		}
		fmt.Println(postId, userID)
		data, err := helpers.PostPageData(h.db, postId, userID)
		if err != nil {
			log.Println(err)
		}
		fmt.Println(data)
		helpers.RenderTemplate(w, "post-id", data)

	} else {
		log.Println("Post ID")
		log.Println("Error. Path Not Allowed.")
		http.Error(w, "Page Not Found", http.StatusNotFound)
		return
	}

}

// To handle "/post/create"

func (h *Handler) NewPost(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/post/create" {
		log.Println("Post Create")
		log.Println("Error. Path Not Allowed.")
		http.Error(w, "Page Not Found", http.StatusNotFound)
		return
	}

	//	Get userID that is making the request
	userID := r.Context().Value(models.UserIDKey).(int)

	//	Check the request comes from a logged-in user or not and act in consequence
	if userID == 0 {
		http.Redirect(w, r, "/#registerModal", http.StatusForbidden)
	} else {
		switch r.Method {
		case http.MethodGet:
			//	Call a function that returns all existent categories:
			categories, err := dbaser.Categories(h.db)
			if err != nil {
				log.Println(err)
			}
			helpers.RenderTemplate(w, "post-create", categories)

		case http.MethodPost:

			post := models.Post{
				UserId:  userID,
				Title:   r.FormValue("title"),
				Content: r.FormValue("content"),
			}
			//	Save the post into the database
			dbaser.AddPost(h.db, post)
			//	Save the categories associated to the post into the database
			//	Print a Succesful message
			referer := r.Referer()
			msg := "Post was created succesfully!"
			helpers.RenderTemplate(w, referer, msg)

		default:
			w.Header().Set("Allow", "GET, POST")
			http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
			return
		}

	}
}

func (h *Handler) Reaction(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/reaction" {
		log.Println("Post Create")
		log.Println("Error. Path Not Allowed.")
		http.Error(w, "Page Not Found", http.StatusNotFound)
		return
	}

	//	Get userID that is making the request
	userID := r.Context().Value(models.UserIDKey).(int)

	//	Check the request comes from a logged-in user or not and act in consequence
	if userID == 0 {
		referer := r.Referer()
		http.Redirect(w, r, referer+"#registerModal", http.StatusForbidden)

	} else {
		switch r.Method {
		case http.MethodPost:
			// Parse form values
			r.ParseForm()

			// Get the postID from the hidden input
			// postID := r.FormValue("post_Id")

			// Get the reactionType based on which button was clicked
			// reaction := r.FormValue("state")

		default:
			w.Header().Set("Allow", "GET, POST")
			http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
			return
		}

	}
}
