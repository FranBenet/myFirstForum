package handlers

import (
	"log"
	"net/http"
	"os"

	"gitea.koodsisu.fi/josepfrancescbenetmorella/literary-lions/dbaser"
	"gitea.koodsisu.fi/josepfrancescbenetmorella/literary-lions/helpers"
	"gitea.koodsisu.fi/josepfrancescbenetmorella/literary-lions/models"
)

// To handle "/register"
func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	log.Println("Requested: Register Handler")

	if r.URL.Path != "/register" {
		log.Println("Error. Path Not Allowed.")
		http.Error(w, "Page Not Found", http.StatusNotFound)
		return
	}

	switch r.Method {

	case http.MethodPost:
		referer := r.Referer()

		// Parse form values
		r.ParseForm()

		user := models.User{
			Email:    r.FormValue("email"),
			Name:     r.FormValue("username"),
			Password: r.FormValue("password"),
		}

		log.Printf("Registration Details Submitted:\n Email: %s,\nUsername: %s,\nPassword: %s\n", r.FormValue("email"), r.FormValue("username"), r.FormValue("password"))

		//	Register user in the db
		_, err := dbaser.AddUser(h.db, user)
		if err != nil {
			log.Println(err)

			cleanURL := helpers.CleanQueryMessages(referer)
			finalURL := helpers.AddQueryMessage(cleanURL, "error", err.Error())

			log.Printf("Redirecting to: %s", finalURL)

			// Redirect to the referer with the error included in the query.
			http.Redirect(w, r, finalURL+"#registerModal", http.StatusFound)
			return

		} else {
			log.Println("Registration Succesful")

			cleanURL := helpers.CleanQueryMessages(referer)

			finalURL := helpers.AddQueryMessage(cleanURL, "success", "Registration Succesful!")

			log.Printf("Redirecting to: %s", finalURL)

			// Redirect to the referer with the error included in the query.
			http.Redirect(w, r, finalURL, http.StatusFound)
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
	log.Println("Requested: Login Handler")

	if r.URL.Path != "/login" {
		log.Println("Error. Path Not Allowed.")
		http.Error(w, "Page Not Found", http.StatusNotFound)
		return
	}

	switch r.Method {

	case http.MethodPost:
		// Parse form values
		r.ParseForm()

		email := r.FormValue("email")
		password := r.FormValue("password")

		log.Printf("Login Details Submitted:\n Email: %s,\nPassword: %s", r.FormValue("email"), r.FormValue("password"))

		//	Get the page where user requested to log in from.
		referer := r.Referer()

		//	Check the username and password are correct.
		valid, err := dbaser.CheckPassword(h.db, email, password)
		if !valid {
			log.Println("Email and Password Incorrect: ", err)

			//	Clean any error/success query from the URL
			cleanURL := helpers.CleanQueryMessages(referer)

			//	Adds a new query error message
			finalURL := helpers.AddQueryMessage(cleanURL, "error", err.Error())

			log.Printf("Redirecting to: %s", finalURL)

			// Redirect to the referer with the error included in the query.
			http.Redirect(w, r, finalURL, http.StatusFound)
			return
		}

		log.Println("Correct Email and Password.")

		//	Get UserData from email.
		user, err := dbaser.UserByEmail(h.db, r.FormValue("email"))
		if err != nil {
			log.Println("Error getting User Data", err)

			//	Clean any error/success query from the URL
			cleanURL := helpers.CleanQueryMessages(referer)

			//	Adds a new query error message
			finalURL := helpers.AddQueryMessage(cleanURL, "error", "Log In failed. Please, try again later.")

			log.Printf("Redirecting to: %s", finalURL)

			// Redirect to the referer with the error included in the query.
			http.Redirect(w, r, finalURL, http.StatusFound)
			return
		}

		log.Println("User ID fetched correctly from the data base.")

		//	Create Session for the user.
		sessionUUID, err := dbaser.AddSession(h.db, user)
		if err != nil {
			log.Println(err)
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

		cleanURL := helpers.CleanQueryMessages(referer)

		log.Printf("Redirecting to: %s", referer)

		http.Redirect(w, r, cleanURL, http.StatusFound)

	default:
		w.Header().Set("Allow", "GET, POST")
		http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
		return
	}

}

// To handle "/login"
func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	log.Println("Requested: Logout Handler")

	if r.URL.Path != "/logout" {
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
