package handlers

import (
	"log"
	"net/http"

	"gitea.koodsisu.fi/josepfrancescbenetmorella/literary-lions/dbaser"
	"gitea.koodsisu.fi/josepfrancescbenetmorella/literary-lions/helpers"
	"gitea.koodsisu.fi/josepfrancescbenetmorella/literary-lions/models"
)

// To handle "/profile"
func (h *Handler) Profile(w http.ResponseWriter, r *http.Request) {
	log.Println("Requested: Profile Handler")

	if r.URL.Path != "/profile" {
		log.Printf("Error. Path %v Not Allowed.", r.URL.Path)
		http.Redirect(w, r, "/404", http.StatusSeeOther)
		return
	}

	// Get userID from the context request. If 0 > user is not logged in.
	userID := r.Context().Value(models.UserIDKey).(int)

	switch r.Method {
	case http.MethodGet:

		if userID == 0 {

			finalURL := helpers.AddQueryMessage("http://localhost:8080/", "error", "Please, log in to access this page.")

			log.Printf("Redirecting to: %s", finalURL)

			http.Redirect(w, r, finalURL, http.StatusSeeOther)

		} else {
			data, err := helpers.ProfilePageData(h.db, userID)

			if err != nil {

				log.Println(err)

				finalURL := helpers.AddQueryMessage("http://localhost:8080/", "error", "Sorry, could not access your profile.")

				log.Printf("Redirecting to: %s", finalURL)

				http.Redirect(w, r, finalURL, http.StatusSeeOther)
			}
			log.Println("Profile data succesfully collected")

			data.Metadata.CurrentPage = "/profile"

			helpers.RenderTemplate(w, "profile", data)

			log.Println("Profile succesfully served")
		}

	case http.MethodPost:
		referer := r.Referer()
		finalURL := referer
		r.ParseForm()
		newAvatar := r.FormValue("avatar")

		_, err := dbaser.UpdateAvatar(h.db, userID, newAvatar)
		if err != nil {
			log.Println(err)

			finalURL = helpers.AddQueryMessage(referer, "error", "Sorry, we could not update your avatar picture. Try again later.")
		}

		log.Printf("User selected avatar: %v ", newAvatar)

		log.Println("Avatar succesfully updated")

		log.Printf("Redirecting to: %s", finalURL)

		http.Redirect(w, r, referer, http.StatusSeeOther)

	default:
		w.Header().Set("Allow", "GET, POST")
		http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
		return
	}
}

// To handle "/posts/liked"
func (h *Handler) LikedPosts(w http.ResponseWriter, r *http.Request) {
	log.Println("Requested: LikedPosts Handler")

	if r.URL.Path != "/liked" {
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

	log.Println("UserID: ", userID)

	//	Get the number of page requested from the query parameters of the URL.
	requestedPage, err := helpers.GetQueryPage(r)
	if err != nil {
		log.Println("No Page Required:", err)
		requestedPage = 1
	}

	log.Println("UserID:", userID, "Requested page number: ", requestedPage)

	if userID == 0 {

		finalURL := helpers.AddQueryMessage("http://localhost:8080/", "error", "Please, log in to access this page.")

		log.Printf("Redirecting to: %s", finalURL)

		http.Redirect(w, r, finalURL, http.StatusSeeOther)

	} else {
		log.Println("Let's get Data: ", err)
		data, err := helpers.MyLikedPostsPageData(h.db, userID, requestedPage)
		if err != nil {
			log.Println("Error getting user's posts: ", err)

			referer := r.Referer()

			finalURL := helpers.AddQueryMessage(referer, "error", "Sorry, we couldn't get your posts. Try again later!")

			log.Printf("Redirecting to: %s", finalURL)

			http.Redirect(w, r, finalURL, http.StatusSeeOther)
		}

		log.Println("Liked Posts collected succesfully")
		data.Metadata.CurrentPage = "/liked"
		helpers.RenderTemplate(w, "liked", data)
		log.Println("Liked Posts succesfully served")
	}
}

func (h *Handler) MyPosts(w http.ResponseWriter, r *http.Request) {
	log.Println("Requested: MyPosts Handler")

	if r.URL.Path != "/myposts" {
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

	//	Get the number of page requested from the query parameters of the URL.
	requestedPage, err := helpers.GetQueryPage(r)
	if err != nil {
		log.Println("No Page Required:", err)
		requestedPage = 1
	}

	log.Println("UserID:", userID, "Requested page number: ", requestedPage)

	if userID == 0 {

		finalURL := helpers.AddQueryMessage("http://localhost:8080/", "error", "Please, log in to access this page.")

		log.Printf("Redirecting to: %s", finalURL)

		http.Redirect(w, r, finalURL, http.StatusSeeOther)

	} else {

		data, err := helpers.MyPostsPageData(h.db, userID, requestedPage)
		log.Println("Posts to display", len(data.Posts))
		if err != nil {
			log.Println("Error getting user's posts: ", err)

			referer := r.Referer()

			finalURL := helpers.AddQueryMessage(referer, "error", "Sorry, we couldn't get your posts. Try again later!")

			log.Printf("Redirecting to: %s", finalURL)

			http.Redirect(w, r, finalURL, http.StatusSeeOther)
		}

		log.Printf("Posts for user: %v collected succesfully", userID)

		data.Metadata.CurrentPage = "/myposts"

		helpers.RenderTemplate(w, "myPosts", data)

		log.Println("My Posts collected served")
	}
}
