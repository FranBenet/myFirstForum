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

// To handle "/login"
func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
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
			helpers.RenderTemplate(w, "login.html", message)
		}

		//	Get User ID from email.
		user, err := dbaser.UserByEmail(h.db, r.FormValue("email"))
		if err != nil {
			log.Println(err)
		}
		userID := user.Id
		fmt.Println(userID)

		//	Create User Session for the User.

		data, err := helpers.MainPageData(h.db, userID)
		if err != nil {
			fmt.Println("Error Getting MainPageData")
			log.Println(err)
		}
		// data.LoggedIn = true

		helpers.RenderTemplate(w, "home", data)

	default:
		w.Header().Set("Allow", "GET, POST")
		http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
		return
	}

}

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

		//	Check if email exists in the db
		exist := dbaser.UserEmailExists(h.db, user.Email)
		if !exist {
			//	RESPONSE TAKES USER TO THE SAME PAGE IT WAS AND PRINT ERROR MESSAGE
			log.Println("Error. Email is already registered. Try logging in.")
			os.Exit(0)
		}

		//	Check if username exists in the db
		exist = dbaser.UsernameExists(h.db, user.Name)
		if !exist {
			//	RESPONSE TAKES USER TO THE SAME PAGE IT WAS AND PRINT ERROR MESSAGE
			log.Println("Error. Username is already registered. Try logging in.")
			os.Exit(0)
		}

		//	Register user in the db
		_, err := dbaser.AddUser(h.db, user)
		if err != nil {
			log.Println(err)
			os.Exit(0)
		}

		//	RESPONSE TAKES USER TO THE SAME PAGE IT WAS AND PRINT SUCCESFUL MESSAGE
		http.Redirect(w, r, "/login", http.StatusFound)

	default:
		w.Header().Set("Allow", "GET, POST")
		http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
		os.Exit(0)
		return
	}
}
