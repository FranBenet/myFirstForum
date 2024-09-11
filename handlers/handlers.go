package handlers

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"gitea.koodsisu.fi/josepfrancescbenetmorella/literary-lions/dbaser"
	"gitea.koodsisu.fi/josepfrancescbenetmorella/literary-lions/helpers"
	"gitea.koodsisu.fi/josepfrancescbenetmorella/literary-lions/models"
	"gitea.koodsisu.fi/josepfrancescbenetmorella/literary-lions/session"
)

// To handle "/".
func Homepage(w http.ResponseWriter, r *http.Request) {
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

	//	Get logIn status for the request
	loggedIn := r.Context().Value("loggedIn").(bool)

	//	Get userID that is making the request
	userID := r.Context().Value("userID").(int)

	//	Call a function passing two paramenters:
	//	- Boolean value indicating if user is loggedIn or not
	//	- Int value indicating the ID of the User.
	//	Example: func getDataHome(loggedIn bool, userID int)

	//	data := getDataHome (loggedIn, userID)

	helpers.RenderTemplate(w, "home", data)
}

// To handle "/post/{id}"
func GetPost(w http.ResponseWriter, r *http.Request) {

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

		//	Get logIn status for the request
		loggedIn := r.Context().Value("loggedIn").(bool)

		//	Get userID that is making the request
		userID := r.Context().Value("userID").(int)

		//	Call a function passing two paramenters:
		//	- Boolean value indicating if user is loggedIn or not
		//	- Int value indicating the ID of the User.
		//	Example: func getDataPostID(loggedIn bool, userID int, postId int)

		//	data := getDataPostID (loggedIn, userID, postId)

		helpers.RenderTemplate(w, "post-id", data)

	} else {
		log.Println("Post ID")
		log.Println("Error. Path Not Allowed.")
		http.Error(w, "Page Not Found", http.StatusNotFound)
		return
	}

}

// To handle "/search"
func Search(w http.ResponseWriter, r *http.Request) {
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

	//	Get logIn status for the request
	loggedIn := r.Context().Value("loggedIn").(bool)

	//	Get userID that is making the request
	userID := r.Context().Value("userID").(int)

	//	Call a function passing two paramenters:
	//	- Boolean value indicating if user is loggedIn or not
	//	- Int value indicating the ID of the User.
	//	Example: func getDataHomeFiltered(loggedIn bool, userID int)

	//	data := getDataHomeFiltered (loggedIn, userID)

	helpers.RenderTemplate(w, "home.html", data)
}

// To handle "/post/create"
func NewPost(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/post/create" {
		log.Println("Post Create")
		log.Println("Error. Path Not Allowed.")
		http.Error(w, "Page Not Found", http.StatusNotFound)
		return
	}

	loggedIn := r.Context().Value("loggedIn").(bool)

	//	Check the request comes from a logged-in user or not and act in consequence
	if !loggedIn {
		http.Redirect(w, r, "/login", http.StatusFound)
	} else {

		switch r.Method {
		case http.MethodGet:

			//	Call a function that returns all existent categories:
			//	data := getDataNewPost() -> this should include all existent categories and a field Message type string that can be used to print error/success when posting.

			helpers.RenderTemplate(w, "post-create", data)

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
func Profile(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/profile" {
		log.Println("Profile")
		log.Println("Error. Path Not Allowed.")
		http.Error(w, "Page Not Found", http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		//	Check if user is logged in
		loggedIn := r.Context().Value("loggedIn").(bool)

		if !loggedIn {
			http.Redirect(w, r, "/login", http.StatusFound)
		} else {
			//	Get userID that is making the request
			userID := r.Context().Value("userID").(int)

			//	Call funtion to collect data
			// data:= getUserDetails(userID int)

			helpers.RenderTemplate(w, "profile.html", data)
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
func LikedPosts(w http.ResponseWriter, r *http.Request) {
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

	//	Check if user is logged in
	loggedIn := r.Context().Value("loggedIn").(bool)

	if !loggedIn {
		http.Redirect(w, r, "/login", http.StatusFound)
	} else {
		//	Get userID that is making the request
		userID := r.Context().Value("userID").(int)

		//	Call function to get data
		// 	data := getLikedPosts(userID int)

		helpers.RenderTemplate(w, "likedPosts.html", data)
	}
}

// To handle "/posts/mined"
func MyPosts(w http.ResponseWriter, r *http.Request) {
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

	//	Check if user is logged in
	loggedIn := r.Context().Value("loggedIn").(bool)

	if !loggedIn {
		http.Redirect(w, r, "/login", http.StatusFound)
	} else {
		// Get userID that is making the request
		userID := r.Context().Value("userID").(int)

		//	Call function to get data
		//	data := getUserPosts(userID int)

		helpers.RenderTemplate(w, "myPosts.html", data)
	}

}

// To handle "/login"
func Login(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/login" {
		log.Println("Login")
		log.Println("Error. Path Not Allowed.")
		http.Error(w, "Page Not Found", http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		helpers.RenderTemplate(w, "login.html", nil)

	case http.MethodPost:
		user := models.User{
			Email:    r.FormValue("email"),
			Password: r.FormValue("password"),
		}
		//	Check the username and password are correct.
		valid, err := dbaser.CheckPassword(user)
		if !valid {
			log.Printf("Incorrect password: %v", err)
			message := "Failed to log in. Try again!"
			helpers.RenderTemplate(w, "login.html", message)
		}

		//	Get UserID from email
		//	userID := userIDbyEmail(r.FormValue("email"))

		//	Create session for the userID
		session.CreateSession(w, userID)

		//	Call a function passing two paramenters:
		//	- Boolean value indicating if user is loggedIn or not
		//	- Int value indicating the ID of the User.
		//	Example: func getDataHome(loggedIn bool, userID int)

		//	data := getDataHome (loggedIn, userID)
		helpers.RenderTemplate(w, "home", data)

	default:
		w.Header().Set("Allow", "GET, POST")
		http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
		return
	}

}

// To handle "/register"
func Register(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/register" {
		log.Println("Register")
		log.Println("Error. Path Not Allowed.")
		http.Error(w, "Page Not Found", http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		helpers.RenderTemplate(w, "register.html", nil)

	case http.MethodPost:
		//	Check if email exists in the db
		//	Create a new user in the DB. Returns an error if failed.

		if err != nil {
			log.Printf("Error registering new user: %v", err)
		}
		message := "Registration Succesful"
		helpers.RenderTemplate(w, "register.html", message)

	default:
		w.Header().Set("Allow", "GET, POST")
		http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
		return
	}
}

// To handle "/users/{username}/profile"
// func UsersPost(w http.ResponseWriter, r *http.Request) {
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
