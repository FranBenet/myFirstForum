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
	log.Println("You are in the Homepage Handler")

	if r.URL.Path != "/" {
		log.Println("Error. Path / Not Allowed.")
		http.Error(w, "Page Not Found", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
		return
	}

	//	Get userID that is making the request
	userID := r.Context().Value(models.UserIDKey).(int)

	// ---------------------------------------------------PROVISIONAL CODE FOR TEST----------------------------------------------------------------------------------------
	//	Get cookie from request
	// var userID int
	// sessionToken, err := r.Cookie("session_token")
	// if err != nil {
	// 	userID = 0
	// 	log.Println("Error getting cookie:", err)
	// } else {
	// 	//	Get session UUID from the cookie
	// 	sessionUUID := sessionToken.Value
	// 	log.Println("Session UUID is:", sessionUUID)
	// 	userID, err = middleware.IsUserLoggedIn(h.db, sessionUUID)
	// 	if err != nil {
	// 		log.Println(err)
	// 	}
	// }

	// ---------------------------------------------------PROVISIONAL CODE FOR TEST----------------------------------------------------------------------------------------
	//	Get the page number requested if not set the page number to 1.
	requestedPage, err := helpers.GetQueryPage(r)
	if err != nil {
		log.Println("No Page Required:", err)
		requestedPage = 1
	}

	log.Println("UserID:", userID, "Requested page number: ", requestedPage)

	//	Get data according to the page requested.
	data, err := helpers.MainPageData(h.db, userID, requestedPage)
	if err != nil {
		log.Println("Error getting data", err)
	}

	//	Get messages from the query parameters
	errorMessage, successMessage, err := helpers.GetQueryMessages(r)
	if err != nil {
		log.Println("Error getting Messages: ", err)
	}

	//	Add Error/success messages to the data.
	data.Metadata.Error = errorMessage
	data.Metadata.Success = successMessage

	fmt.Println("Logged In status: ", data.Metadata.LoggedIn)

	helpers.RenderTemplate(w, "home", data)

}

// To handle "/post/{id}"
func (h *Handler) GetPost(w http.ResponseWriter, r *http.Request) {
	log.Println("You are in the GetPost Handler")

	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
		return
	}

	//	Get userID that is making the request
	userID := r.Context().Value(models.UserIDKey).(int)

	// ---------------------------------------------------PROVISIONAL CODE FOR TEST----------------------------------------------------------------------------------------
	//	Get cookie from request
	// var userID int

	// sessionToken, err := r.Cookie("session_token")
	// if err != nil {
	// 	userID = 0
	// 	log.Println(err)
	// } else {
	// 	//	Get the value of the session from the cookie
	// 	sessionUUID := sessionToken.Value
	// 	userID, err = middleware.IsUserLoggedIn(h.db, sessionUUID)
	// 	if err != nil {
	// 		userID = 0
	// 		log.Println(err)
	// 	}
	// }

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
			finalURL := helpers.AddQueryMessage("http://localhost:8080/", "error", "Page not available.")
			http.Redirect(w, r, finalURL, http.StatusFound)
		}
		log.Println("COMMENTS:", data.Comments)
		//	Get messages from the query parameters
		errorMessage, successMessage, err := helpers.GetQueryMessages(r)
		if err != nil {
			log.Println("Error getting Messages: ", err)
		}

		//	Add Error/success messages to the data.
		data.Metadata.Error = errorMessage
		data.Metadata.Success = successMessage

		fmt.Println("Logged In status: ", data.Metadata.LoggedIn)

		helpers.RenderTemplate(w, "post-id", data)
		return
	} else {
		log.Println("Post ID")
		log.Println("Error. Path Not Allowed.")
		http.Error(w, "Page Not Found", http.StatusNotFound)
		return
	}

}

// To handle "/post/create"

func (h *Handler) NewPost(w http.ResponseWriter, r *http.Request) {
	log.Println("You are in the NewPost Handler")
	if r.URL.Path != "/post/create" {
		log.Println("Post Create")
		log.Println("Error. Path Not Allowed.")
		http.Error(w, "Page Not Found", http.StatusNotFound)
		return
	}

	//	Get userID that is making the request
	userID := r.Context().Value(models.UserIDKey).(int)

	// ---------------------------------------------------PROVISIONAL CODE FOR TEST----------------------------------------------------------------------------------------
	//	Get cookie from request
	// var userID int
	// sessionToken, err := r.Cookie("session_token")

	// if err != nil {
	// 	userID = 0
	// 	log.Println("Error Getting cookie:", err)
	// } else {
	// 	//	Get session UUID from the cookie
	// 	sessionUUID := sessionToken.Value
	// 	fmt.Println("session:", sessionUUID)
	// 	userID, err = middleware.IsUserLoggedIn(h.db, sessionUUID)
	// 	if err != nil {
	// 		userID = 0
	// 		log.Println("Error, validating session:", err)

	// 	}
	// }
	// ---------------------------------------------------PROVISIONAL CODE FOR TEST----------------------------------------------------------------------------------------

	//	Check IS USER LOGGED IN?
	if userID == 0 {
		http.Redirect(w, r, "/#registerModal", http.StatusForbidden)
	} else {
		switch r.Method {
		case http.MethodGet:

			data, err := helpers.CreatePostData(h.db, userID)
			if err != nil {
				log.Println(err)
			}

			//	Get messages from the query parameters to embedd in the data and print on screen
			errorMessage, successMessage, err := helpers.GetQueryMessages(r)
			if err != nil {
				log.Println("Error getting Messages: ", err)
			}

			//	Add Error/success messages to the data.
			data.Metadata.Error = errorMessage
			data.Metadata.Success = successMessage

			helpers.RenderTemplate(w, "post-create", data)

		case http.MethodPost:
			// Parse form values
			r.ParseForm()

			post := models.Post{
				UserId:  userID,
				Title:   r.FormValue("title"),
				Content: r.FormValue("content"),
			}

			// //	Save the post into the database
			postID, err := dbaser.AddPost(h.db, post)
			if err != nil {
				log.Println(err)

				//	Get the page where user send the request from.
				referer := r.Referer()

				//	This function includes a query parameter in the URL with an error/success to be printed on screen
				finalURL := helpers.AddQueryMessage(referer, "error", "Error saving post. Try again later!")

				log.Printf("Redirecting to: %s", finalURL)

				http.Redirect(w, r, finalURL, http.StatusFound)

				return
			}

			//	Add Categories for the post ID
			var categories []string
			var postCategories []string
			categories = append(categories, r.FormValue("category1"), r.FormValue("category2"), r.FormValue("category3"))

			log.Println("Categories: ", categories)
			for _, category := range categories {
				if category != "" {
					postCategories = append(postCategories, strings.TrimSpace(category))
				}
			}

			log.Println("Final Categories: ", postCategories)

			err = dbaser.AddPostCategories(h.db, postCategories, postID)
			if err != nil {
				//	Get the page where user send the request from.
				referer := r.Referer()

				//	This function includes a query parameter in the URL with an error/success to be printed on screen
				finalURL := helpers.AddQueryMessage(referer, "error", "Error saving categories for the post.")

				log.Printf("Redirecting to: %s", finalURL)

				http.Redirect(w, r, finalURL, http.StatusFound)

				return
			}
			//	Get the page where user requested to log in.
			referer := r.Referer()

			finalURL := helpers.AddQueryMessage(referer, "success", "Post created succesfully!")

			log.Printf("Redirecting to: %s", finalURL)

			http.Redirect(w, r, finalURL, http.StatusFound)

		default:
			w.Header().Set("Allow", "GET, POST")
			http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
			return
		}

	}
}

func (h *Handler) Reaction(w http.ResponseWriter, r *http.Request) {
	log.Println("You are in the Reaction Handler")
	if r.URL.Path != "/reaction" {
		log.Println("Post Create")
		log.Println("Error. Path Not Allowed.")
		http.Error(w, "Page Not Found", http.StatusNotFound)
		return
	}

	//	Get userID that is making the request
	userID := r.Context().Value(models.UserIDKey).(int)

	// ---------------------------------------------------PROVISIONAL CODE FOR TEST----------------------------------------------------------------------------------------
	//	Get cookie from request
	// var userID int
	// sessionToken, err := r.Cookie("session_token")
	// if err != nil {
	// 	userID = 0
	// 	log.Println("Error Getting cookie:", err)
	// } else {
	// 	//	Get the value of the session from the cookie
	// 	sessionUUID := sessionToken.Value
	// 	fmt.Println("session:", sessionUUID)
	// 	userID, err = middleware.IsUserLoggedIn(h.db, sessionUUID)
	// 	if err != nil {
	// 		userID = 0
	// 		log.Println("Error, validating session:", err)

	// 	}
	// }
	// ---------------------------------------------------PROVISIONAL CODE FOR TEST----------------------------------------------------------------------------------------

	//	Check the request comes from a logged-in user or not and act in consequence
	if userID == 0 {
		referer := r.Referer()

		finalURL := helpers.AddQueryMessage(referer, "error", "Need to be logged in for that action")

		http.Redirect(w, r, finalURL+"#loginModal", http.StatusFound)
		// http.Redirect(w, r, finalURL, http.StatusForbidden)
		return

	} else {
		switch r.Method {
		case http.MethodPost:
			// Parse form values
			r.ParseForm()

			// Get the Post ID
			id := r.FormValue("post_Id")

			if id != "" {
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
			}

			//	Get Comment ID
			id = r.FormValue("comment_Id")
			log.Println("User reacted to a comment.")
			log.Println("The comment ID is: ", id)
			if id != "" {
				log.Println("The comment ID is: ", id)
				commentID, err := strconv.Atoi(id)
				if err != nil {
					log.Println(err)
				}

				// Get the reactionType based on which button was clicked
				form_reaction := r.FormValue("state")

				var reaction bool
				switch form_reaction {
				case "like":
					log.Println("The comment was LIKED")
					reaction = true
				case "dislike":
					log.Println("The comment was DISLIKED")
					reaction = false
				}
				newReaction := models.CommentReaction{CommentId: commentID, UserId: userID, Liked: reaction}
				dbaser.AddCommentReaction(h.db, newReaction)
			}
			//	Get the page where user requested to log in.
			referer := r.Referer()

			finalURL := helpers.CleanQueryMessages(referer)

			// Redirect to the referer with the error included in the query.
			http.Redirect(w, r, finalURL, http.StatusFound)
		default:
			w.Header().Set("Allow", "GET, POST")
			http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
			return
		}
	}
}

func (h *Handler) NewComment(w http.ResponseWriter, r *http.Request) {
	log.Println("You are in the NewComment Handler")
	if r.URL.Path != "/post/comment" {
		log.Println("Post Create")
		log.Println("Error. Path Not Allowed.")
		http.Error(w, "Page Not Found", http.StatusNotFound)
		return
	}

	//	Get userID that is making the request
	userID := r.Context().Value(models.UserIDKey).(int)

	// ---------------------------------------------------PROVISIONAL CODE FOR TEST----------------------------------------------------------------------------------------
	//	Get cookie from request
	// var userID int
	// sessionToken, err := r.Cookie("session_token")

	// if err != nil {
	// 	userID = 0
	// 	log.Println("Error Getting cookie:", err)
	// } else {
	// 	//	Get session UUID from the cookie
	// 	sessionUUID := sessionToken.Value
	// 	fmt.Println("session:", sessionUUID)
	// 	userID, err = middleware.IsUserLoggedIn(h.db, sessionUUID)
	// 	if err != nil {
	// 		userID = 0
	// 		log.Println("Error, validating session:", err)

	// 	}
	// }
	// ---------------------------------------------------PROVISIONAL CODE FOR TEST----------------------------------------------------------------------------------------

	//	Check IS USER LOGGED IN?
	if userID == 0 {
		http.Redirect(w, r, "/#registerModal", http.StatusForbidden)
	} else {
		switch r.Method {
		case http.MethodPost:
			//	Get the page where user send the request from.
			referer := r.Referer()

			// Parse form values
			r.ParseForm()
			content := r.FormValue("new-comment")

			// Get the postID from the hidden input
			id := r.FormValue("post_Id")
			postID, err := strconv.Atoi(id)
			if err != nil {
				log.Println(err)
				//	This function includes a query parameter in the URL with an error/success to be printed on screen
				finalURL := helpers.AddQueryMessage(referer, "error", "Error saving comment. Try again later!")

				log.Printf("Redirecting to: %s", finalURL)

				http.Redirect(w, r, finalURL, http.StatusFound)

				return
			}

			newComment := models.Comment{PostId: postID, UserId: userID, Content: content}

			// //	Save the comment into the database
			_, err = dbaser.AddComment(h.db, newComment)
			if err != nil {
				log.Println(err)

				//	This function includes a query parameter in the URL with an error/success to be printed on screen
				finalURL := helpers.AddQueryMessage(referer, "error", "Error saving comment. Try again later!")

				log.Printf("Redirecting to: %s", finalURL)

				http.Redirect(w, r, finalURL, http.StatusFound)

				return
			}

			finalURL := helpers.AddQueryMessage(referer, "success", "Post created succesfully!")

			log.Printf("Redirecting to: %s", finalURL)

			http.Redirect(w, r, finalURL, http.StatusFound)

		default:
			w.Header().Set("Allow", "GET, POST")
			http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
			return
		}

	}
}
