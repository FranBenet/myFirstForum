package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"gitea.koodsisu.fi/josepfrancescbenetmorella/literary-lions/helpers"
	"gitea.koodsisu.fi/josepfrancescbenetmorella/literary-lions/models"
)

// To handle "/search"
func (h *Handler) Search(w http.ResponseWriter, r *http.Request) {
	log.Println("Requested: Search Handler")

	if r.URL.Path != "/search" {
		log.Println("Error. Path Not Allowed.")
		http.Error(w, "Page Not Found", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get userID from the context request. If 0 > user is not logged in.
	userID := r.Context().Value(models.UserIDKey).(int)

	//	Get the page number requested if not set the page number to 1.
	requestedPage, err := helpers.GetQueryPage(r)
	if err != nil {
		log.Println("Page is not available or specified")
		requestedPage = 1
	}
	fmt.Println(requestedPage)
	fmt.Println(userID)
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
	log.Println("Requested: Filter Handler")

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

	// Get userID from the context request. If 0 > user is not logged in.
	userID := r.Context().Value(models.UserIDKey).(int)

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
		log.Println("No Filter Applied")

		referer := r.Referer()

		finalURL := helpers.AddQueryMessage(referer, "error", "No filters applied")

		log.Printf("Redirecting to: %s", finalURL)

		http.Redirect(w, r, "finalURL", http.StatusFound)
		return

	} else if filterCategory != "" && filterSort == "" {
		log.Println("UserID:", userID, "Requested Filter by: ", filterCategory, "Page number: ", requestedPage)

		categoryId, err := strconv.Atoi(filterCategory)
		if err != nil {
			log.Println("Category requested does not exist")

			referer := r.Referer()

			finalURL := helpers.AddQueryMessage(referer, "error", "This category does not exist")

			log.Printf("Redirecting to: %s", finalURL)

			http.Redirect(w, r, finalURL, http.StatusFound)

			return
		}

		fmt.Println(categoryId)

		// Get data according to the page requested.
		// data, err := helpers.CollectCategoryData(h.db, userID, requestedPage, categoryId)
		// if err != nil {
		// 	log.Println("Error getting category data", err)
		// }

		// helpers.RenderTemplate(w, "home", data)

	} else if filterSort != "" && filterCategory == "" {
		log.Println("UserID:", userID, "Requested Filter by: ", filterSort, "Page number: ", requestedPage)

		switch filterSort {
		case "likes":
			// data, err := helpers.CollectLikesData(h.db, userID, requestedPage, filterSort)
			// helpers.RenderTemplate(w, "home", data)
		case "dislikes":
			// data, err := helpers.CollectDislikesData(h.db, userID, requestedPage, filterSort)
			// helpers.RenderTemplate(w, "home", data)
		case "mostrecent":
			// data, err := helpers.CollectMostRecentData(h.db, userID, requestedPage, filterSort)
			// helpers.RenderTemplate(w, "home", data)
		default:
		}

	} else {
		log.Println("UserID:", userID, "Requested too many filters.")

		referer := r.Referer()

		finalURL := helpers.AddQueryMessage(referer, "error", "Too many filters applied")

		log.Printf("Redirecting to: %s", finalURL)

		http.Redirect(w, r, "finalURL", http.StatusFound)

		return
	}
}

// To handle "/users/{username}/profile"
func (h *Handler) UsersPost(w http.ResponseWriter, r *http.Request) {
	log.Println("Requested: UsersPost Handler")

	if r.URL.Path != "/users/" {
		log.Println("Error. Path Not Allowed.")
		http.Error(w, "Page Not Found", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
		return
	}

	path := r.URL.Path
	pathDivide := strings.Split(path, "/")
	user := pathDivide[3]

	fmt.Println(user)
	// data := helpers.GetData()
	// helpers.RenderTemplate(w, "", data)
}
