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
	log.Println("You are in the Profile Handler")
	if r.URL.Path != "/profile" {
		log.Println("Profile")
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
	log.Println("You are in the LikedPosts Handler")
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

	//	Get the page number requested if not set the page number to 1.
	requestedPage, err := helpers.GetQueryPage(r)
	if err != nil {
		log.Println("No Page Required:", err)
		requestedPage = 1
	}

	log.Println("UserID:", userID, "Requested page number: ", requestedPage)

	if userID == 0 {
		referer := "http://localhost:8080/"

		finalURL := helpers.AddQueryMessage(referer, "error", "Please, log in to access this page.")

		log.Printf("Redirecting to: %s", finalURL)

		http.Redirect(w, r, finalURL, http.StatusFound)

	} else {

		// data, err := helpers.SomeFunction(h.db, userID, requestedPage)
		// if err != nil {
		// 	log.Println("Error getting user's posts: ", err)

		// 	data.Metadata.Error = "Sorry, we couldn't get your posts. Try again later!"
		// 	referer := r.Referer()

		// 	finalURL := helpers.AddQueryMessage(referer, "error", "Sorry, we couldn't get your posts. Try again later!")

		// 	log.Printf("Redirecting to: %s", finalURL)

		// 	http.Redirect(w, r, finalURL, http.StatusFound)
		// }
		// log.Println(data)
		// helpers.RenderTemplate(w, "liked", data)
	}
}

// To handle "/posts/mined"
func (h *Handler) MyPosts(w http.ResponseWriter, r *http.Request) {
	log.Println("You are in the MyPosts Handler")
	if r.URL.Path != "/myposts" {
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
	//	Get the page number requested if not set the page number to 1.
	requestedPage, err := helpers.GetQueryPage(r)
	if err != nil {
		log.Println("No Page Required:", err)
		requestedPage = 1
	}

	log.Println("UserID:", userID, "Requested page number: ", requestedPage)

	if userID == 0 {
		referer := "http://localhost:8080/"

		finalURL := helpers.AddQueryMessage(referer, "error", "Please, log in to access this page.")

		log.Printf("Redirecting to: %s", finalURL)

		http.Redirect(w, r, finalURL, http.StatusFound)

	} else {

		// data, err := helpers.SomeFunction(h.db, userID, requestedPage)
		// if err != nil {
		// 	log.Println("Error getting user's posts: ", err)

		// 	data.Metadata.Error = "Sorry, we couldn't get your posts. Try again later!"
		// 	referer := r.Referer()

		// 	finalURL := helpers.AddQueryMessage(referer, "error", "Sorry, we couldn't get your posts. Try again later!")

		// 	log.Printf("Redirecting to: %s", finalURL)

		// 	http.Redirect(w, r, finalURL, http.StatusFound)
		// }
		// log.Println(data)
		// helpers.RenderTemplate(w, "myPosts", data)
	}
}
