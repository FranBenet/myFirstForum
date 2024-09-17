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
	userID := r.Context().Value(models.UserIDKey).(int)

	data, err := helpers.MainPageData(h.db, userID)
	if err != nil {
		//	HANDLE ERROR
	}

	helpers.RenderTemplate(w, "home", data)
}

// To handle "/post/{id}"
func (h *Handler) GetPost(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
		return
	}

	path := r.URL.Path
	pathDivide := strings.Split(path, "/")

	if len(pathDivide) == 3 && pathDivide[1] == "post" && pathDivide[0] == "" {

		postId, err := strconv.Atoi(pathDivide[2])
		if err != nil {
			http.Error(w, "Invalid post ID", http.StatusBadRequest)
		}

		//	Get userID that is making the request
		userID := r.Context().Value(models.UserIDKey).(int)

		data, err := helpers.PostPageData(h.db, postId, userID)
		if err != nil {
			//	HANDLE ERROR
		}
		helpers.RenderTemplate(w, "post-id", data)

	} else {
		log.Println("Post ID")
		log.Println("Error. Path Not Allowed.")
		http.Error(w, "Page Not Found", http.StatusNotFound)
		return
	}

}

// To handle "/search"
func (h *Handler) Search(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/search" {
		log.Println("Seach")
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

	// data, err := helpers.MainPageDataFilter(h.db, userID)
	// if err != nil {
	// 	//	HANDLE ERROR
	// }
	// helpers.RenderTemplate(w, "home.html", data)
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
		http.Redirect(w, r, "/login", http.StatusFound)
	} else {

		switch r.Method {
		case http.MethodGet:

			//	Call a function that returns all existent categories:
			//	data := getDataNewPost() -> this should include all existent categories and a field Message type string that can be used to print error/success when posting.

			// helpers.RenderTemplate(w, "post-create", data)

		case http.MethodPost:
			//	Save the post into the database
			//	Print a Succesful message

		default:
			w.Header().Set("Allow", "GET, POST")
			http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
			return
		}

	}
}

// To handle "/profile"
func (h *Handler) Profile(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/profile" {
		log.Println("Profile")
		log.Println("Error. Path Not Allowed.")
		http.Error(w, "Page Not Found", http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		//	Get userID that is making the request
		userID := r.Context().Value(models.UserIDKey).(int)

		if userID == 0 {
			http.Redirect(w, r, "/login", http.StatusFound)

		} else {
			// data, err := helpers.ProfilePageData(h.db, userID)
			// if err != nil {
			// 	//	HANDLE ERROR
			// }
			// helpers.RenderTemplate(w, "profile.html", data)
		}

	case http.MethodPost:
		//	To edit password and name
	default:
		w.Header().Set("Allow", "GET, POST")
		http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
		return
	}
}

// To handle "/posts/liked"
func (h *Handler) LikedPosts(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/posts/liked" {
		log.Println("Liked Posts")
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
	userID := r.Context().Value(models.UserIDKey).(int)

	if userID == 0 {
		http.Redirect(w, r, "/login", http.StatusFound)
	} else {

		//	Call function to get data
		// 	data := getLikedPosts(userID int)

		// helpers.RenderTemplate(w, "likedPosts.html", data)
	}
}

// To handle "/posts/mined"
func (h *Handler) MyPosts(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/posts/myPosts" {
		log.Println("posts/myPosts")
		log.Println("Error. Path Not Allowed.")
		http.Error(w, "Page Not Found", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
		return
	}
	// Get userID that is making the request
	userID := r.Context().Value(models.UserIDKey).(int)

	if userID == 0 {
		http.Redirect(w, r, "/login", http.StatusFound)
	} else {

		//	Call function to get data
		//	data := getUserPosts(userID int)

		// helpers.RenderTemplate(w, "myPosts.html", data)
	}

}

// To handle "/login"
func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/login" {
		log.Println("Login")
		log.Println("Error. Path Not Allowed.")
		http.Error(w, "Page Not Found", http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		// helpers.RenderTemplate(w, "login.html", nil)
		fmt.Println("DO WE NEED A GET METHOD FOR LOGIN???")

	case http.MethodPost:

		//	Check the username and password are correct.
		valid, err := dbaser.CheckPassword(h.db, r.FormValue("email"), r.FormValue("password"))
		if !valid {
			log.Printf("Incorrect password: %v", err)
			message := "Failed to log in. Try again!"
			helpers.RenderTemplate(w, "login.html", message)
		}

		//	Get User ID from email.
		user, err := dbaser.UserByEmail(h.db, r.FormValue("email"))
		if err != nil {
			//	SOMETHING
		}
		userID := user.Id
		fmt.Println(userID)

		//	Create User Session for the User.

		// session.CreateSession(w, userID)

		//	data := getDataHome (loggedIn, userID)
		// helpers.RenderTemplate(w, "home", data)

	default:
		w.Header().Set("Allow", "GET, POST")
		http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
		return
	}

}

// To handle "/register"
func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/register" {
		log.Println("Register")
		log.Println("Error. Path Not Allowed.")
		http.Error(w, "Page Not Found", http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		// helpers.RenderTemplate(w, "register.html", nil)
		fmt.Println("DO WE NEED A GET METHOD FOR REGISTER???")

	case http.MethodPost:
		user := models.User{
			Email:    r.FormValue("email"),
			Name:     r.FormValue("username"),
			Password: r.FormValue("password"),
		}

		//	Check if email exists in the db
		exist := dbaser.UserEmailExists(h.db, user.Email)
		if !exist {
			//	RESPONSE TAKES USER TO THE SAME PAGE IT WAS AND PRINT ERROR MESSAGE
		}

		//	Check if username exists in the db
		exist = dbaser.UsernameExists(h.db, user.Name)
		if !exist {
			//	RESPONSE TAKES USER TO THE SAME PAGE IT WAS AND PRINT ERROR MESSAGE
		}

		//	Register user in the db
		_, err := dbaser.AddUser(h.db, user)
		if err != nil {
			log.Printf("Error registering new user: %v", err)
		}
		//	RESPONSE TAKES USER TO THE SAME PAGE IT WAS AND PRIN SUCCESFUL MESSAGE

	default:
		w.Header().Set("Allow", "GET, POST")
		http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
		return
	}
}

// To handle "/users/{username}/profile"
// func (h *Handler)UsersPost(w http.ResponseWriter, r *http.Request) {
// 	if r.URL.Path != "/users/{username}/profile" {
// 		log.Println("Users Post")
// 		log.Println("Error. Path Not Allowed.")
// 		http.Error(w, "Page Not Found", http.StatusNotFound)
// 		return
// 	}

// 	if r.Method != http.MethodGet {
// 		w.Header().Set("Allow", http.MethodGet)
// 		http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
// 		return
// 	}

// 	data := helpers.GetData()
// 	helpers.RenderTemplate(w, "", data)
// }
