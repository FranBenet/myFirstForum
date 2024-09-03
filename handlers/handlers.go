package handlers

import (
	"literary-lions/pkg/helpers"
	"log"
	"net/http"
)

// To handle "/"
func homepage(w http.ResponseWriter, r *http.Request) {
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
func getPost(w http.ResponseWriter, r *http.Request) {
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
func search(w http.ResponseWriter, r *http.Request) {
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
func login(w http.ResponseWriter, r *http.Request) {
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
func register(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/register" {
		log.Println("Register")
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
	helpers.RenderTemplate(w, "register.html", data)
}

// To handle "/post/create"
func newPost(w http.ResponseWriter, r *http.Request) {
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
func profile(w http.ResponseWriter, r *http.Request) {
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
func usersPost(w http.ResponseWriter, r *http.Request) {
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
func notifications(w http.ResponseWriter, r *http.Request) {
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
func likedPosts(w http.ResponseWriter, r *http.Request) {
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
func myPosts(w http.ResponseWriter, r *http.Request) {
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
