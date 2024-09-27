package handlers

import (
	"fmt"
	"log"
	"net/http"

	"gitea.koodsisu.fi/josepfrancescbenetmorella/literary-lions/helpers"
	"gitea.koodsisu.fi/josepfrancescbenetmorella/literary-lions/middleware"
)

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

	// ---------------------------------------------------PROVISIONAL CODE FOR TEST----------------------------------------------------------------------------------------
	//	Get cookie from request
	var userID int
	sessionToken, err := r.Cookie("session_token")

	if err != nil {
		userID = 0
		log.Println("Error Getting cookie:", err)
	} else {
		//	Get session UUID from the cookie
		sessionUUID := sessionToken.Value
		fmt.Println("session:", sessionUUID)
		userID, err = middleware.IsUserLoggedIn(h.db, sessionUUID)
		if err != nil {
			userID = 0
			log.Println("Error, validating session:", err)

		}
	}
	// ---------------------------------------------------PROVISIONAL CODE FOR TEST----------------------------------------------------------------------------------------

	//	Get the page where user send the request from.
	referer := r.Referer()

	// Parse form values
	r.ParseForm()
	content := r.FormValue("new-comment")

	// data, err := helpers.MainPageDataFilter(h.db, userID)
	// if err != nil {
	// 	//	HANDLE ERROR
	// }
	// helpers.RenderTemplate(w, "home.html", data)
}

// To handle "/filter"
func (h *Handler) Filter(w http.ResponseWriter, r *http.Request) {
	log.Println("You are in the Filter Handler")
	if r.URL.Path != "/filter" {
		log.Println("Filter")
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
		log.Println("Error Getting cookie:", err)
	} else {
		//	Get session UUID from the cookie
		sessionUUID := sessionToken.Value
		fmt.Println("session:", sessionUUID)
		userID, err = middleware.IsUserLoggedIn(h.db, sessionUUID)
		if err != nil {
			userID = 0
			log.Println("Error, validating session:", err)

		}
	}
	// ---------------------------------------------------PROVISIONAL CODE FOR TEST----------------------------------------------------------------------------------------

	filter, err := helpers.GetQueryFilter(r)
	if err != nil {
		log.Println(err)

		referer := r.Referer()

		finalURL := helpers.AddQueryMessage(referer, "error", "Please, log in to access this page.")

		log.Printf("Redirecting to: %s", finalURL)

		http.Redirect(w, r, finalURL, http.StatusFound)
	}

	switch

	// data, err := helpers.MainPageDataFilter(h.db, userID)
	// if err != nil {
	// 	//	HANDLE ERROR
	// }
	// helpers.RenderTemplate(w, "home.html", data)
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
