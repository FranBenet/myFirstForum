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
		log.Printf("Error. Path %v Not Allowed.", r.URL.Path)
		http.Redirect(w, r, "/404", http.StatusSeeOther)
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

	log.Println("UserID:", userID, "Requested page number: ", requestedPage)

	// Parse form values
	r.ParseForm()
	query := r.FormValue("search")

	log.Println("User requesting search posts for: ", query)

	data, err := helpers.SearchPageData(h.db, query, userID, requestedPage)
	if err != nil {
		log.Println("Error getting searched posts: ", err)

		referer := r.Referer()

		finalURL := helpers.AddQueryMessage(referer, "error", "Something happend and  we couldn't get posts for that search. Try again later!")

		log.Printf("Redirecting to: %s", finalURL)

		http.Redirect(w, r, finalURL, http.StatusSeeOther)
	}
	log.Println("Posts to be displayed: ", len(data.Posts))

	helpers.RenderTemplate(w, "home", data)
}

// To handle "/filter"
func (h *Handler) Filter(w http.ResponseWriter, r *http.Request) {
	log.Println("Requested: Filter Handler")

	if r.URL.Path != "/filter" {
		log.Printf("Error. Path %v Not Allowed.", r.URL.Path)
		http.Redirect(w, r, "/404", http.StatusSeeOther)
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

		http.Redirect(w, r, "finalURL", http.StatusSeeOther)
		return

	} else if filterCategory != "" && filterSort == "" {
		log.Println("UserID:", userID, "Requested Filter by: ", filterCategory, "Page number: ", requestedPage)

		categoryId, err := strconv.Atoi(filterCategory)
		if err != nil {
			log.Println("Category requested does not exist")

			referer := r.Referer()

			finalURL := helpers.AddQueryMessage(referer, "error", "This category does not exist")

			log.Printf("Redirecting to: %s", finalURL)

			http.Redirect(w, r, finalURL, http.StatusSeeOther)

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

		http.Redirect(w, r, "finalURL", http.StatusSeeOther)

		return
	}
}

// To handle "/users/{username}/profile"
func (h *Handler) UsersPost(w http.ResponseWriter, r *http.Request) {
	log.Println("Requested: UsersPost Handler")

	if r.URL.Path != "/users/" {
		log.Printf("Error. Path %v Not Allowed.", r.URL.Path)
		http.Redirect(w, r, "/404", http.StatusSeeOther)
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
