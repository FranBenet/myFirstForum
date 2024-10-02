package handlers

import (
	"database/sql"
	"fmt"
	"html/template"
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

func (h *Handler) Homepage(w http.ResponseWriter, r *http.Request) {
	log.Println("Requested: Homepage Handler")

	if r.URL.Path != "/" {
		log.Printf("Error. Path %v Not Allowed.", r.URL.Path)
		http.Redirect(w, r, "/404", http.StatusSeeOther)

		return
	}

	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
		return
	}

	//	Get userID from the context request. If 0 > user is not logged in.
	userID := r.Context().Value(models.UserIDKey).(int)

	//	Get the number of page requested from the query parameters of the URL.
	requestedPage, err := helpers.GetQueryPage(r)
	if err != nil {
		requestedPage = 1
	}

	//	Get data according to the page requested.
	data, err := helpers.MainPageData(h.db, userID, requestedPage)
	if err != nil {
		log.Println("Error getting data", err)

		http.Redirect(w, r, "/500", http.StatusSeeOther)

		return
	}

	log.Println("Data for Homepage succesfully collected")

	//	Get error/successful messages from the query parameters
	errorMessage, successMessage, err := helpers.GetQueryMessages(r)
	if err != nil {
		log.Println("Error getting Messages: ", err)
	}

	//	Add Error/success messages to the data.
	data.Metadata.Error = errorMessage
	data.Metadata.Success = successMessage

	//	Rendering tempaltes and sending Response.
	helpers.RenderTemplate(w, "home", data)
	log.Println("Homepage succesfully served")
}

func (h *Handler) GetPost(w http.ResponseWriter, r *http.Request) {
	log.Println("Requested: GetPost Handler")

	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
		return
	}

	//	Get userID from the context request. If 0 > user is not logged in.
	userID := r.Context().Value(models.UserIDKey).(int)

	path := r.URL.Path
	pathDivide := strings.Split(path, "/")

	if len(pathDivide) == 3 && pathDivide[1] == "post" && pathDivide[0] == "" {

		postId, err := strconv.Atoi(pathDivide[2])
		if err != nil {
			log.Println("ID for the post is not a number. ", err)

			//	This function includes a query parameter in the URL with an error/success to be printed on screen
			finalURL := helpers.AddQueryMessage("http://localhost:8080/", "error", "ID for the post is not a number")

			log.Printf("Redirecting to: %s", finalURL)

			http.Redirect(w, r, finalURL, http.StatusSeeOther)

			return
		}

		data, err := helpers.PostPageData(h.db, postId, userID)
		if err != nil {
			log.Println(err)

			finalURL := helpers.AddQueryMessage("http://localhost:8080/", "error", "Post does not exist.")

			http.Redirect(w, r, finalURL, http.StatusSeeOther)

			return
		}

		log.Println("Data for GetPost succesfully collected")

		//	Get error/successful messages from the query parameters
		errorMessage, successMessage, err := helpers.GetQueryMessages(r)
		if err != nil {
			log.Println("Error getting Messages: ", err)
		}

		//	Add Error/success messages to the data.
		data.Metadata.Error = errorMessage
		data.Metadata.Success = successMessage
		data.Metadata.CurrentPage = "/post/"
		helpers.RenderTemplate(w, "post-id", data)

		log.Println("GetPost succesfully served")

		return
	} else {
		log.Println("Path Not Allowed.")

		//	This function includes a query parameter in the URL with an error/success to be printed on screen
		finalURL := helpers.AddQueryMessage("http://localhost:8080/", "error", "Path does not exist.")

		log.Printf("Redirecting to: %s", finalURL)

		http.Redirect(w, r, finalURL, http.StatusSeeOther)

		return
	}
}

func (h *Handler) NewPost(w http.ResponseWriter, r *http.Request) {
	log.Println("Requested: NewPost Handler")

	if r.URL.Path != "/post/create" {
		log.Printf("Error. Path %v Not Allowed.", r.URL.Path)
		http.Redirect(w, r, "/404", http.StatusSeeOther)
		return
	}

	//	Get userID from the context request. If 0 > user is not logged in.
	userID := r.Context().Value(models.UserIDKey).(int)

	if userID == 0 {

		http.Redirect(w, r, "/#registerModal", http.StatusSeeOther)

		return
	} else {
		switch r.Method {
		case http.MethodGet:

			data, err := helpers.CreatePostData(h.db, userID)
			if err != nil {
				log.Println(err)

				//	This function includes a query parameter in the URL with an error/success to be printed on screen
				finalURL := helpers.AddQueryMessage("http://localhost:8080/", "error", "Ups! Something happened! \n Please, try again later.")

				log.Printf("Redirecting to: %s", finalURL)

				http.Redirect(w, r, finalURL, http.StatusSeeOther)
				return
			}

			log.Println("Data for NewPost succesfully collected")

			//	Get error/successful messages from the query parameters
			errorMessage, successMessage, err := helpers.GetQueryMessages(r)
			if err != nil {
				log.Println("Error getting Messages: ", err)
			}

			//	Add Error/success messages to the data.
			data.Metadata.Error = errorMessage
			data.Metadata.Success = successMessage
			data.Metadata.CurrentPage = "/post/create"

			helpers.RenderTemplate(w, "post-create", data)

			log.Println("NewPost succesfully served")

			return
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
				finalURL := helpers.AddQueryMessage(referer, "error", "Error saving post. Try again later.")

				log.Printf("Redirecting to: %s", finalURL)

				http.Redirect(w, r, finalURL, http.StatusSeeOther)

				return
			}

			//	Add Categories for the post ID
			var categories []string
			var postCategories []string
			categories = append(categories, r.FormValue("category1"), r.FormValue("category2"), r.FormValue("category3"))

			for _, category := range categories {
				if category != "" {
					postCategories = append(postCategories, strings.TrimSpace(category))
				}
			}
			log.Println("Categories for this post", postCategories)
			err = dbaser.AddPostCategories(h.db, postCategories, postID)
			if err != nil {
				//	Get the page where user send the request from.
				referer := r.Referer()

				//	This function includes a query parameter in the URL with an error/success to be printed on screen
				finalURL := helpers.AddQueryMessage(referer, "error", "Error saving categories for the post.")

				log.Printf("Redirecting to: %s", finalURL)

				http.Redirect(w, r, finalURL, http.StatusSeeOther)

				return
			}

			log.Println("Data for NewPost succesfully added in the database.")

			//	Get the page where user requested to log in.
			referer := r.Referer()

			finalURL := helpers.AddQueryMessage(referer, "success", "Post created succesfully!")

			log.Printf("Redirecting to: %s", finalURL)

			http.Redirect(w, r, finalURL, http.StatusFound)

			log.Println("NewPost succesfully served")

			return

		default:
			w.Header().Set("Allow", "GET, POST")
			http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
			return
		}

	}
}

func (h *Handler) Reaction(w http.ResponseWriter, r *http.Request) {
	log.Println("Requested: Reaction Handler")

	if r.URL.Path != "/reaction" {
		log.Printf("Error. Path %v Not Allowed.", r.URL.Path)
		http.Redirect(w, r, "/404", http.StatusSeeOther)
		return
	}

	//	Get userID from the context request. If 0 > user is not logged in.
	userID := r.Context().Value(models.UserIDKey).(int)

	//	Check the request comes from a logged-in user or not and act in consequence
	if userID == 0 {
		referer := r.Referer()

		finalURL := helpers.AddQueryMessage(referer, "error", "Need to be logged in for that action")

		log.Printf("Redirecting to: %s", finalURL)

		http.Redirect(w, r, finalURL+"#loginModal", http.StatusSeeOther)

		return

	} else {
		switch r.Method {
		case http.MethodPost:
			// Parse form values
			r.ParseForm()

			// Get the Post ID that was reacted
			id := r.FormValue("post_Id")

			if id != "" {
				log.Println("User reacted to the Post.Id = ", id)

				postID, err := strconv.Atoi(id)
				if err != nil {
					log.Println(err)

					referer := r.Referer()

					finalURL := helpers.AddQueryMessage(referer, "error", "Ups! Something happened. Try again later.")

					log.Printf("Redirecting to: %s", finalURL)

					http.Redirect(w, r, finalURL, http.StatusSeeOther)

					return
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

				_, err = dbaser.AddPostReaction(h.db, newReaction)
				if err != nil {
					log.Println("We could not save this reaction", err)

					referer := r.Referer()

					finalURL := helpers.AddQueryMessage(referer, "error", "Ups! Something happened. Try again later.")

					log.Printf("Redirecting to: %s", finalURL)

					http.Redirect(w, r, finalURL, http.StatusSeeOther)

					return
				}
				log.Println("Reaction succesfully included in the data base")
			}

			//	Get Comment ID
			id = r.FormValue("comment_Id")

			if id != "" {
				log.Println("User reacted to a Comment. Id = ", id)

				commentID, err := strconv.Atoi(id)
				if err != nil {
					log.Println(err)

					referer := r.Referer()

					finalURL := helpers.AddQueryMessage(referer, "error", "Ups! Something happened. Try again later.")

					log.Printf("Redirecting to: %s", finalURL)

					http.Redirect(w, r, finalURL, http.StatusSeeOther)

					return
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
				newReaction := models.CommentReaction{CommentId: commentID, UserId: userID, Liked: reaction}

				_, err = dbaser.AddCommentReaction(h.db, newReaction)
				if err != nil {
					log.Println("We could not save this reaction", err)

					referer := r.Referer()

					finalURL := helpers.AddQueryMessage(referer, "error", "Ups! Something happened. Try again later.")

					log.Printf("Redirecting to: %s", finalURL)

					http.Redirect(w, r, finalURL, http.StatusSeeOther)

					return
				}

				log.Println("Reaction succesfully included in the data base")
			}

			//	Get the page where user requested to log in.
			referer := r.Referer()

			finalURL := helpers.CleanQueryMessages(referer)

			log.Printf("Redirecting to: %s", finalURL)
			// Redirect to the referer with the error included in the query.
			http.Redirect(w, r, finalURL, http.StatusFound)

			return
		default:
			w.Header().Set("Allow", "POST")
			http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
			return
		}
	}
}

func (h *Handler) NewComment(w http.ResponseWriter, r *http.Request) {
	log.Println("Requested: NewComment Handler")

	if r.URL.Path != "/post/comment" {
		log.Printf("Error. Path %v Not Allowed.", r.URL.Path)
		http.Redirect(w, r, "/404", http.StatusSeeOther)
		return
	}

	// Get userID from the context request. If 0 > user is not logged in.
	userID := r.Context().Value(models.UserIDKey).(int)

	//	Check IS USER LOGGED IN?
	if userID == 0 {

		http.Redirect(w, r, "/#registerModal", http.StatusSeeOther)

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

				http.Redirect(w, r, finalURL, http.StatusSeeOther)

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

				http.Redirect(w, r, finalURL, http.StatusSeeOther)

				return
			}

			log.Println("Comment succesfully added to the database")

			finalURL := helpers.AddQueryMessage(referer, "success", "Comment created succesfully!")

			log.Printf("Redirecting to: %s", finalURL)

			http.Redirect(w, r, finalURL, http.StatusFound)

			return
		default:
			w.Header().Set("Allow", "POST")
			http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
			return
		}

	}
}

func (h *Handler) NotFound(w http.ResponseWriter, r *http.Request) {
	log.Println("Requested: NotFound Handler")

	if r.URL.Path != "/404" {
		log.Printf("Error. Path %v Not Allowed.", r.URL.Path)
		http.Error(w, "Page Not Found", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
		return
	}

	tmpl, err := template.ParseFiles("web/templates/404.html")
	if err != nil {
		fmt.Printf("Error Parsing Template: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNotFound)
	err = tmpl.ExecuteTemplate(w, "404.html", nil)
	if err != nil {
		fmt.Printf("Error Executing Template: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

}

func (h *Handler) InternalError(w http.ResponseWriter, r *http.Request) {
	log.Println("Requested: NotFound Handler")

	if r.URL.Path != "/500" {
		log.Printf("Error. Path %v Not Allowed.", r.URL.Path)
		http.Error(w, "Page Not Found", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
		return
	}

	tmpl, err := template.ParseFiles("web/templates/500.html")
	if err != nil {
		fmt.Printf("Error Parsing Template: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusInternalServerError)
	err = tmpl.ExecuteTemplate(w, "500.html", nil)
	if err != nil {
		fmt.Printf("Error Executing Template: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
