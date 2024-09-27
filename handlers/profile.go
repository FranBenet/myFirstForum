package handlers

import (
	"log"
	"net/http"

	"gitea.koodsisu.fi/josepfrancescbenetmorella/literary-lions/helpers"
	"gitea.koodsisu.fi/josepfrancescbenetmorella/literary-lions/middleware"
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

	//	Get userID that is making the request
	// userID := r.Context().Value(models.UserIDKey).(int)

	// ---------------------------------------------------PROVISIONAL CODE FOR TEST----------------------------------------------------------------------------------------
	//	Get cookie from request
	var userID int
	sessionToken, err := r.Cookie("session_token")
	if err != nil {
		userID = 0
		log.Println("Error getting cookie:", err)
	} else {
		//	Get session UUID from the cookie
		sessionUUID := sessionToken.Value
		log.Println("Session UUID is:", sessionUUID)
		userID, err = middleware.IsUserLoggedIn(h.db, sessionUUID)
		if err != nil {
			log.Println(err)
		}
	}

	// ---------------------------------------------------PROVISIONAL CODE FOR TEST----------------------------------------------------------------------------------------

	switch r.Method {
	case http.MethodGet:

		if userID == 0 {
			referer := "http://localhost:8080/"

			finalURL := helpers.AddQueryMessage(referer, "error", "Please, log in to access this page.")

			log.Printf("Redirecting to: %s", finalURL)

			http.Redirect(w, r, finalURL, http.StatusFound)

		} else {
			data, err := helpers.ProfilePageData(h.db, userID)
			if err != nil {
				log.Println("Profile Page. Error getting user data:", err)

				referer := "http://localhost:8080/"

				finalURL := helpers.AddQueryMessage(referer, "error", "Sorry, could not access your profile.")

				log.Printf("Redirecting to: %s", finalURL)

				http.Redirect(w, r, finalURL, http.StatusFound)
			}

			helpers.RenderTemplate(w, "profile", data)
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
