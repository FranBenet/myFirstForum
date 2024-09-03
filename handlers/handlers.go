package handlers

import (
	"log"
	"net/http"

	"gitea.koodsisu.fi/josepfrancescbenetmorella/literary-lions/helpers"
)

// To handle "/"
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

	data := helpers.GetData()
	helpers.RenderTemplate(w, "home.html", data)
}

// To handle "/post/{id}"
func GetPost(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/post/{id}" {
		log.Println("Post ID")
		log.Println("Error. Path Not Allowed.")
		http.Error(w, "Page Not Found", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
		return
	}

	data := helpers.GetData()
	helpers.RenderTemplate(w, "post-view.html", data)
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

	data := helpers.GetData()
	helpers.RenderTemplate(w, "home.html", data)
}

// To handle "/login"
func Login(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/login" {
		log.Println("Login")
		log.Println("Error. Path Not Allowed.")
		http.Error(w, "Page Not Found", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
		return
	}

	data := helpers.GetData()
	helpers.RenderTemplate(w, "login.html", data)
}

// To handle "/register"
func Register(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/register" {
		log.Println("Register")
		log.Println("Error. Path Not Allowed.")
		http.Error(w, "Page Not Found", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
		return
	} else if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.Method == http.MethodGet {
		data := helpers.GetData()
		helpers.RenderTemplate(w, "register.html", data)
	} else if r.Method == http.MethodPost {
		//	func CreateNewUser()
		// data := helpers.GetData()
		helpers.RenderTemplate(w, "registerSuccesful.html", nil)
	}
}

// To handle "/post/create"
func NewPost(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/post/create" {
		log.Println("Post Create")
		log.Println("Error. Path Not Allowed.")
		http.Error(w, "Page Not Found", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
		return
	}

	data := helpers.GetData()
	helpers.RenderTemplate(w, "create.html", data)
}

// To handle "/profile"
func Profile(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/profile" {
		log.Println("Profile")
		log.Println("Error. Path Not Allowed.")
		http.Error(w, "Page Not Found", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
		return
	}

	data := helpers.GetData()
	helpers.RenderTemplate(w, "profile.html", data)
}

// To handle "/users/{username}/profile"
func UsersPost(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/users/{username}/profile" {
		log.Println("Users Post")
		log.Println("Error. Path Not Allowed.")
		http.Error(w, "Page Not Found", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
		return
	}

	data := helpers.GetData()
	helpers.RenderTemplate(w, "", data)
}

// To handle "/notifications"
func Notifications(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/notifications" {
		log.Println("Notifications")
		log.Println("Error. Path Not Allowed.")
		http.Error(w, "Page Not Found", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
		return
	}

	data := helpers.GetData()
	helpers.RenderTemplate(w, "notifications.html", data)
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

	data := helpers.GetData()
	helpers.RenderTemplate(w, "", data)
}

// To handle "/posts/mined"
func MyPosts(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/posts/mined" {
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

	data := helpers.GetData()
	helpers.RenderTemplate(w, "", data)
}
