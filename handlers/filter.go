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
	fmt.Println(userID)
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
	// referer := r.Referer()

	// Parse form values
	r.ParseForm()
	// content := r.FormValue("new-comment")

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
	fmt.Println(userID)
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

	//	Get the page number requested if not set the page number to 1.
	requestedPage, err := helpers.GetQueryPage(r)
	if err != nil {
		log.Println("No Page Required:", err)
		requestedPage = 1
	}
	fmt.Println(requestedPage)

	filterCategory, err := helpers.GetQueryFilter(r, "category")
	if err != nil {
		log.Println(err)
	}

	filterSort, err := helpers.GetQueryFilter(r, "sort")
	if err != nil {
		log.Println(err)
	}

	if filterSort == "" && filterCategory == "" {
		log.Println("No Filters on the URL")
		referer := r.Referer()

		finalURL := helpers.AddQueryMessage(referer, "error", "Filters not available")

		log.Printf("Redirecting to: %s", finalURL)
		http.Redirect(w, r, "finalURL", http.StatusFound)
		return

	} else if filterCategory != "" && filterSort == "" {
		log.Println("Category Filter in the URL")

		// categoryId, err := strconv.Atoi(filterCategory)
		// if err != nil {
		// 	log.Println("Category does not exist")
		// 	referer := r.Referer()

		// 	finalURL := helpers.AddQueryMessage(referer, "error", "This category does not exist")

		// 	log.Printf("Redirecting to: %s", finalURL)
		// 	http.Redirect(w, r, finalURL, http.StatusFound)
		// 	return
		// }

		//	Get data according to the page requested.
		// data, err := helpers.CollectCategoryData(h.db, userID, requestedPage, categoryId)
		// if err != nil {
		// 	log.Println("Error getting category data", err)
		// }

		// fmt.Println("Logged In status: ", data.Metadata.LoggedIn)

		// helpers.RenderTemplate(w, "home", data)

	} else if filterSort != "" && filterCategory == "" {
		log.Println("Sort Filter in the URL")

		switch filterSort {
		case "likes":
			// data, err := helpers.CollectLikesData(h.db, userID, requestedPage, filterSort)
		case "dislikes":
			// data, err := helpers.CollectDislikesData(h.db, userID, requestedPage, filterSort)
		case "mostrecent":
			// data, err := helpers.CollectMostRecentData(h.db, userID, requestedPage, filterSort)
		default:
		}

	} else {
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
