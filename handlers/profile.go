package handlers

import (
	"log"
	"net/http"

	"gitea.koodsisu.fi/josepfrancescbenetmorella/literary-lions/models"
)

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
