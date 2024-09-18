package handlers

import (
	"log"
	"net/http"
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

	// data, err := helpers.MainPageDataFilter(h.db, userID)
	// if err != nil {
	// 	//	HANDLE ERROR
	// }
	// helpers.RenderTemplate(w, "home.html", data)
}

// To handle "/filter"
func (h *Handler) Filter(w http.ResponseWriter, r *http.Request) {
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
