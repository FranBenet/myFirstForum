package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"gitea.koodsisu.fi/josepfrancescbenetmorella/literary-lions/dbaser"
	"gitea.koodsisu.fi/josepfrancescbenetmorella/literary-lions/helpers"
	"gitea.koodsisu.fi/josepfrancescbenetmorella/literary-lions/models"
)

// To handle "/register"
func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	fmt.Println("You have reached Register function")
	if r.URL.Path != "/register" {
		log.Println("Register")
		log.Println("Error. Path Not Allowed.")
		http.Error(w, "Page Not Found", http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		// helpers.RenderTemplate(w, "register.html", nil)
		fmt.Println("DO WE NEED A GET METHOD FOR REGISTER???")

	case http.MethodPost:
		user := models.User{
			Email:    r.FormValue("email"),
			Name:     r.FormValue("username"),
			Password: r.FormValue("password"),
		}
		fmt.Println(user)
		//	Register user in the db
		_, err := dbaser.AddUser(h.db, user)
		if err != nil {
			log.Println(err)

		}

		//	RESPONSE TAKES USER TO THE SAME PAGE IT WAS AND PRINT SUCCESFUL MESSAGE
		http.Redirect(w, r, "/#registerModal", http.StatusFound)

	default:
		w.Header().Set("Allow", "GET, POST")
		http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
		os.Exit(0)
		return
	}
}

// To handle "/login"
func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("You have reached Login function")
	if r.URL.Path != "/login" {
		log.Println("Login")
		log.Println("Error. Path Not Allowed.")
		http.Error(w, "Page Not Found", http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		// helpers.RenderTemplate(w, "login.html", nil)
		fmt.Println("DO WE NEED A GET METHOD FOR LOGIN???")

	case http.MethodPost:

		//	Check the username and password are correct.
		valid, err := dbaser.CheckPassword(h.db, r.FormValue("email"), r.FormValue("password"))
		if !valid {
			log.Printf("Incorrect password: %v", err)
			message := "Failed to log in. Try again!"
			helpers.RenderTemplate(w, "/#loginModal", message)
		}

		//	Create Session by the
		//	Get User ID from email.
		user, err := dbaser.UserByEmail(h.db, r.FormValue("email"))
		if err != nil {
			log.Println(err)
		}
		userID := user.Id

		data, err := helpers.MainPageData(h.db, userID)
		if err != nil {
			fmt.Println("Error Getting MainPageData")
			log.Println(err)
		}
		data.LoggedIn = true

		helpers.RenderTemplate(w, "home", data)

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
	var userID int
	sessionToken, err := r.Cookie("session_token")
	if err != nil {
		userID = 0
		log.Println(err)
	} else {
		//	Delete the session
		sessionUUID := sessionToken.Value
		_, err = dbaser.DeleteSession(h.db, sessionUUID)
		if err != nil {
			log.Println(err)
		}
	}

	data, err := helpers.MainPageData(h.db, userID)
	if err != nil {
		fmt.Println("Error Getting MainPageData")
		log.Println(err)
	}
	helpers.RenderTemplate(w, "home", data)
}
