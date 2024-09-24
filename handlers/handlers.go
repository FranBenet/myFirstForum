package handlers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"net/url"
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
		log.Println("You are in the Homepage Handler")
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
		log.Println("There is no cookie available:", err)
	} else {
		//	Get the value of the session from the cookie
		sessionUUID := sessionToken.Value
		log.Println("Session UUID is:", sessionUUID)
		userID, err = middleware.IsUserLoggedIn(h.db, sessionUUID)
		if err != nil {
			userID = 0
			log.Println("Error. No valid session available:", err)
		}
	}

	// ---------------------------------------------------PROVISIONAL CODE FOR TEST----------------------------------------------------------------------------------------

	data, err := helpers.MainPageData(h.db, userID)
	if err != nil {
		log.Println("Error getting data", err)
	}

	log.Println("Is user logged in?:", data.LoggedIn)

	// Parse the query parameters from the URL
	query := r.URL.Query()

	// Get the value of the "error" parameter
	errorMessage := query.Get("error")
	fmt.Println("Query Parameter Error:", errorMessage)
	successMessage := query.Get("success")
	fmt.Println("Query Parameter Success:", successMessage)

	unescapedError, err := url.QueryUnescape(errorMessage)
	if err != nil || unescapedError == "" {
		log.Println("Error unsecaping Error:", err)
	}
	data.Metadata.Error = unescapedError

	unescapedSuccess, err := url.QueryUnescape(successMessage)
	if err != nil || unescapedSuccess == "" {
		log.Println("Error unsecaping Success:", err)
	}
	data.Metadata.Success = unescapedSuccess
	data.Metadata.Error = unescapedError

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

		data, err := helpers.PostPageData(h.db, postId, userID)
		if err != nil {
			log.Println(err)
		}

		// Parse the query parameters from the URL
		fmt.Println("URL:", r.URL.Path)
		query := r.URL.Query()

		// Get the value of the "error" parameter
		errorMessage := query.Get("errorLogIn")
		fmt.Println("ErrorMessage1:", errorMessage)

		unescapedError, err := url.QueryUnescape(errorMessage)

		if err != nil || unescapedError == "" {
			log.Println(err)
		}
		data.Metadata.Error = unescapedError

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
	// userID := r.Context().Value(models.UserIDKey).(int)

	// ---------------------------------------------------PROVISIONAL CODE FOR TEST----------------------------------------------------------------------------------------
	//	Get cookie from request
	var userID int
	sessionToken, err := r.Cookie("session_token")

	if err != nil {
		userID = 0
		log.Println("Error Getting cookie:", err)
	} else {
		//	Get the value of the session from the cookie
		sessionUUID := sessionToken.Value
		fmt.Println("session:", sessionUUID)
		userID, err = middleware.IsUserLoggedIn(h.db, sessionUUID)
		if err != nil {
			userID = 0
			log.Println("Error, validating session:", err)

		}
	}

	// ---------------------------------------------------PROVISIONAL CODE FOR TEST----------------------------------------------------------------------------------------

	//	Check the request comes from a logged-in user or not and act in consequence
	if userID == 0 {
		http.Redirect(w, r, "/#registerModal", http.StatusForbidden)
	} else {
		switch r.Method {
		case http.MethodGet:

			//	Call a function that returns categories and loggedIn status:
			data, err := helpers.CreatePostData(h.db, userID)
			if err != nil {
				log.Println(err)
			}
			helpers.RenderTemplate(w, "post-create", data)

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
	// userID := r.Context().Value(models.UserIDKey).(int)

	// ---------------------------------------------------PROVISIONAL CODE FOR TEST----------------------------------------------------------------------------------------
	//	Get cookie from request
	var userID int
	sessionToken, err := r.Cookie("session_token")

	if err != nil {
		userID = 0
		log.Println("Error Getting cookie:", err)
	} else {
		//	Get the value of the session from the cookie
		sessionUUID := sessionToken.Value
		fmt.Println("session:", sessionUUID)
		userID, err = middleware.IsUserLoggedIn(h.db, sessionUUID)
		if err != nil {
			userID = 0
			log.Println("Error, validating session:", err)

		}
	}
	// ---------------------------------------------------PROVISIONAL CODE FOR TEST----------------------------------------------------------------------------------------

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
			id := r.FormValue("post_Id")
			postID, err := strconv.Atoi(id)
			if err != nil {
				log.Println(err)
			}

			// Get the reactionType based on which button was clicked
			form_reaction := r.FormValue("state")

			var reaction bool
			switch form_reaction {
			case "like":
				reaction = true
			case "dislike":
				reaction = false
			}

			newReaction := models.PostReaction{PostId: postID, UserId: userID, Liked: reaction}
			dbaser.AddPostReaction(h.db, newReaction)

			//	Get the page where user requested to log in.
			referer := r.Referer()

			// Redirect to the referer with the error included in the query.
			http.Redirect(w, r, referer, http.StatusFound)
		default:
			w.Header().Set("Allow", "GET, POST")
			http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
			return
		}

	}
}
