package handlers

import (
	"log"
	"net/http"
	"strconv"

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

	log.Println("Searched Posts data succesfully collected")

	helpers.RenderTemplate(w, "search", data)

	log.Println("Searched Posts succesfully served")
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

	filterCategory, err := helpers.GetQueryFilter(r, "category")
	if err != nil {
		log.Println(err)
	}

	filterSort, err := helpers.GetQueryFilter(r, "sort")
	if err != nil {
		log.Println(err)
	}

	log.Println("UserID:", userID, "Requested Filter: ", filterCategory, filterSort, "Page number: ", requestedPage)

	if filterSort == "" && filterCategory == "" {
		log.Println("No Filter Applied")

		referer := r.Referer()

		finalURL := helpers.AddQueryMessage(referer, "error", "No filters applied")

		log.Printf("Redirecting to: %s", finalURL)

		http.Redirect(w, r, "finalURL", http.StatusSeeOther)
		return

	} else if filterCategory != "" && filterSort == "" {
		log.Println("UserID:", userID, "Requested Category: ", filterCategory, "Page number: ", requestedPage)

		categoryId, err := strconv.Atoi(filterCategory)
		if err != nil {
			log.Println("Category requested does not exist")

			referer := r.Referer()

			finalURL := helpers.AddQueryMessage(referer, "error", "This category does not exist")

			log.Printf("Redirecting to: %s", finalURL)

			http.Redirect(w, r, finalURL, http.StatusSeeOther)

			return
		}

		// Get data according to the page requested.
		data, err := helpers.CategoryFilterPageData(h.db, categoryId, userID, requestedPage)
		if err != nil {
			log.Println("Error getting Category posts: ", err)

			referer := r.Referer()

			finalURL := helpers.AddQueryMessage(referer, "error", "Something happend and  we couldn't get posts for that category. Try again later!")

			log.Printf("Redirecting to: %s", finalURL)

			http.Redirect(w, r, finalURL, http.StatusSeeOther)
		}
		log.Println("Category Posts data succesfully collected")

		// Get error/successful messages from the query parameters
		errorMessage, successMessage, err := helpers.GetQueryMessages(r)
		if err != nil {
			log.Println("Error getting Messages: ", err)
		}

		// Add Error/success messages to the data.
		data.Metadata.Error = errorMessage
		data.Metadata.Success = successMessage
		data.Pagination.IdRequested = filterCategory
		// log.Println(data.Trending)
		helpers.RenderTemplate(w, "category_page", data)

		log.Println("Category Posts succesfully served")

	} else if filterSort != "" && filterCategory == "" {
		log.Println("UserID:", userID, "Requested Filter by: ", filterSort, "Page number: ", requestedPage)

		switch filterSort {
		case "likes":
			data, err := helpers.TrendingPageData(h.db, userID, requestedPage)
			if err != nil {
				log.Println("Error getting trend posts: ", err)

				referer := r.Referer()

				finalURL := helpers.AddQueryMessage(referer, "error", "Something happend and  we couldn't get posts for that filter. Try again later!")

				log.Printf("Redirecting to: %s", finalURL)

				http.Redirect(w, r, finalURL, http.StatusSeeOther)
			}
			log.Println("Trend Posts succesfully collected")

			// Get error/successful messages from the query parameters
			errorMessage, successMessage, err := helpers.GetQueryMessages(r)
			if err != nil {
				log.Println("Error getting Messages: ", err)
			}

			// Add Error/success messages to the data.
			data.Metadata.Error = errorMessage
			data.Metadata.Success = successMessage

			helpers.RenderTemplate(w, "filter_liked_page", data)

			log.Println("Trend Posts succesfully served")

		case "dislikes":
			data, err := helpers.UntrendingPageData(h.db, userID, requestedPage)
			if err != nil {
				log.Println("Error getting untrend posts: ", err)

				referer := r.Referer()

				finalURL := helpers.AddQueryMessage(referer, "error", "Something happend and  we couldn't get posts for that filter. Try again later!")

				log.Printf("Redirecting to: %s", finalURL)

				http.Redirect(w, r, finalURL, http.StatusSeeOther)
			}

			log.Println("Disliked Posts succesfully collected")

			// Get error/successful messages from the query parameters
			errorMessage, successMessage, err := helpers.GetQueryMessages(r)
			if err != nil {
				log.Println("Error getting Messages: ", err)
			}

			// Add Error/success messages to the data.
			data.Metadata.Error = errorMessage
			data.Metadata.Success = successMessage
			helpers.RenderTemplate(w, "filter_dislikes_page", data)
			log.Println("Disliked Posts succesfully served")

		default:
			http.Redirect(w, r, "/", http.StatusSeeOther)
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

	if r.URL.Path != "/user" {
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
	userIdRequested, err := helpers.GetQueryFilter(r, "user")
	if err != nil {
		log.Println("No user Id required:", err)
	}

	log.Println("UserID:", userID, "Requested page number: ", requestedPage, "for post from userID: ", userIdRequested)

	//	Convert userIdRequested to string
	userIdReqInt, err := strconv.Atoi(userIdRequested)
	if err != nil {
		log.Println("ID for the user is not a number. ", err)

		//	This function includes a query parameter in the URL with an error/success to be printed on screen
		finalURL := helpers.AddQueryMessage("http://localhost:8080/", "error", "ID for the user is not a number")

		log.Printf("Redirecting to: %s", finalURL)

		http.Redirect(w, r, finalURL, http.StatusSeeOther)
		return

	}

	data, err := helpers.UsersPageData(h.db, userID, userIdReqInt, requestedPage)
	if err != nil {
		log.Println("Error getting user's posts: ", err)

		referer := r.Referer()

		finalURL := helpers.AddQueryMessage(referer, "error", "Sorry, we couldn't get your posts. Try again later!")

		log.Printf("Redirecting to: %s", finalURL)

		http.Redirect(w, r, finalURL, http.StatusSeeOther)
	}
	log.Println("Users Posts succesfully collected")

	data.Pagination.IdRequested = userIdRequested

	log.Println(data.Pagination.IdRequested)

	helpers.RenderTemplate(w, "user", data)
	log.Println("Users Posts succesfully served")
}
