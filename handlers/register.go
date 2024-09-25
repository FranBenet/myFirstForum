package handlers

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"gitea.koodsisu.fi/josepfrancescbenetmorella/literary-lions/dbaser"
	"gitea.koodsisu.fi/josepfrancescbenetmorella/literary-lions/helpers"
	"gitea.koodsisu.fi/josepfrancescbenetmorella/literary-lions/models"
)

// To handle "/register"
func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	log.Println("You are in the Register Handler")

	if r.URL.Path != "/register" {
		log.Println("Register")
		log.Println("Error. Path Not Allowed.")
		http.Error(w, "Page Not Found", http.StatusNotFound)
		return
	}

	switch r.Method {

	case http.MethodPost:
		userID := 0
		currentPage, err := helpers.GetRefererPage(r)
		if err != nil {
			log.Println(err)
		}

		data, err := helpers.MainPageData(h.db, userID, currentPage)
		if err != nil {
			log.Println(err)
		}

		user := models.User{
			Email:    r.FormValue("email"),
			Name:     r.FormValue("username"),
			Password: r.FormValue("password"),
		}
		log.Printf("Registration Details. Email: %s, Username: %s, Password: %s", r.FormValue("email"), r.FormValue("username"), r.FormValue("password"))

		//	Register user in the db
		_, err = dbaser.AddUser(h.db, user)
		if err != nil {
			log.Println(err)
			//	Include the error in the data to be printed on screen.
			data.Metadata.Error = fmt.Sprintf("%s", err)

			//	Get the page where user requested to log in.
			referer := r.Referer()

			//	Convert URL to url.url format
			refererURL, err2 := url.Parse(referer)
			if err2 != nil {
				log.Println("Failed to parse referer:", err2)
				refererURL = &url.URL{Path: "/"}
			}

			//	Delete any queries that the url may have
			refererURL.RawQuery = ""

			// Get all query values
			query := refererURL.Query()

			// Add/Update the error to the query.
			query.Set("error", err.Error())

			// Set the updated query back to the referer URL
			refererURL.RawQuery = query.Encode()

			log.Printf("Redirecting to: %s", refererURL.String())

			// Redirect to the referer with the error included in the query.
			http.Redirect(w, r, refererURL.String(), http.StatusFound)
			return

		} else {
			log.Println("User Registered in the data base succesfully!")

			data.Metadata.Success = "Registration Succesful!"
			//	Get the page where user requested to log in.
			referer := r.Referer()

			//	Convert URL to url.url format
			refererURL, err2 := url.Parse(referer)
			if err2 != nil {
				log.Println("Failed to parse referer:", err2)
				refererURL = &url.URL{Path: "/"}
			}

			//	Delete any queries that the url may have
			refererURL.RawQuery = ""

			// Get all query values
			query := refererURL.Query()

			// Add/Update the error to the query.
			query.Set("success", "Registration Succesful!")

			// Set the updated query back to the referer URL
			refererURL.RawQuery = query.Encode()

			log.Printf("Redirecting to: %s", refererURL.String())

			// Redirect to the referer with the error included in the query.
			http.Redirect(w, r, refererURL.String(), http.StatusFound)
			return
		}

	default:
		w.Header().Set("Allow", "GET, POST")
		http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
		os.Exit(0)
		return
	}
}

// To handle "/login"
func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	log.Println("You are in the Login Handler")
	if r.URL.Path != "/login" {
		log.Println("Login")
		log.Println("Error. Path Not Allowed.")
		http.Error(w, "Page Not Found", http.StatusNotFound)
		return
	}

	switch r.Method {

	case http.MethodPost:
		email := r.FormValue("email")
		password := r.FormValue("password")

		log.Printf("Login Details. Email: %s, Password: %s", r.FormValue("email"), r.FormValue("password"))

		//	Check the username and password are correct.
		// 	If not correct, return to the page with error message.
		valid, err := dbaser.CheckPassword(h.db, email, password)
		if !valid {
			log.Println("Email and Password Incorrect:", err)

			//	Get the page where user requested to log in.
			referer := r.Referer()

			//	Convert URL to url.url format
			refererURL, err2 := url.Parse(referer)
			if err2 != nil {
				log.Println("Failed to parse referer:", err2)
				refererURL = &url.URL{Path: "/"}
			}

			//	Delete any queries that the url may have
			refererURL.RawQuery = ""

			// Get all query values
			query := refererURL.Query()

			// Add/Update the error to the query.
			query.Set("error", err.Error())
			fmt.Println(err)

			// Set the updated query back to the referer URL
			refererURL.RawQuery = query.Encode()

			log.Printf("Redirecting to: %s", refererURL.String())

			// Redirect to the referer with the error included in the query.
			http.Redirect(w, r, refererURL.String(), http.StatusFound)
			return
		}

		log.Println("Email and Password Correct")

		//	Get UserData from email.
		user, err := dbaser.UserByEmail(h.db, r.FormValue("email"))
		if err != nil {
			log.Println("UserByEmail", err)
		}

		log.Println("User ID got correctly from data base.")

		//	Create Session for the user.
		sessionUUID, err := dbaser.AddSession(h.db, user)
		if err != nil {
			log.Println("AddSession", err)
		}

		log.Println("New Session created succesfully!")

		cookie := &http.Cookie{
			Name:     "session_token",
			Value:    sessionUUID,
			Path:     "/",
			HttpOnly: true,
			Secure:   true,
		}

		// Send the cookie to the client
		http.SetCookie(w, cookie)

		//	Get the page where user request to log in.
		referer := r.Referer()

		//	Convert the string url into a url.url format.
		refererURL, err := url.Parse(referer)
		if err != nil {
			log.Println("Failed to parse referer:", err)
			refererURL = &url.URL{Path: "/"}
		}

		// Get the query parameters
		queryParams := refererURL.Query()

		// Remove the "error" and "success" query parameter
		queryParams.Del("error")
		queryParams.Del("success")

		// Set the modified query parameters back to the URL
		refererURL.RawQuery = queryParams.Encode()

		//	Convert url.url format to string format
		referer = refererURL.String()

		log.Printf("Redirecting to: %s", referer)
		http.Redirect(w, r, referer, http.StatusFound)

	default:
		w.Header().Set("Allow", "GET, POST")
		http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
		return
	}

}

// To handle "/login"
func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	fmt.Println("You have reached Logout function")
	if r.URL.Path != "/logout" {
		log.Println("Logout")
		log.Println("Error. Path Not Allowed.")
		http.Error(w, "Page Not Found", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
		return
	}

	//	Get cookie from request
	sessionToken, err := r.Cookie("session_token")
	if err != nil {
		log.Println(err)
	} else {
		//	Delete the session
		sessionUUID := sessionToken.Value
		_, err = dbaser.DeleteSession(h.db, sessionUUID)
		if err != nil {
			log.Println(err)
		}
	}

	http.Redirect(w, r, "/", http.StatusFound)

}
