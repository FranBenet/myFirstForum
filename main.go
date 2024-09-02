package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"gitea.koodsisu.fi/josepfrancescbenetmorella/literary-lions/dbaser"
	"gitea.koodsisu.fi/josepfrancescbenetmorella/literary-lions/models"
)

func routes() *http.ServeMux {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./web/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", homepage)
	// mux.HandleFunc("/search", search)
	// mux.HandleFunc("/login", logIn)
	// mux.HandleFunc("/register", register)
	// mux.HandleFunc("/create-post", createPost)
	// mux.HandleFunc("/profile", profile)
	// mux.HandleFunc("/notifications", notifications)
	// mux.HandleFunc("/post-id", postId)
	// mux.HandleFunc("/liked-posts", likedPosts)
	// mux.HandleFunc("/my-posts", myPosts)
	return mux
}

func main() {
	// dbaser.InitDb()
	// dbaser.PopulateDb()
	// user := models.User{"madrabbit@matrix.com", "whiterabbit", "Rz_;*$78)"}
	db, err := dbaser.DbHandle("forum.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	fmt.Println(dbaser.PostCategories(db, models.Post{Id: 1}))

	//	PROVISIONAL STARTING WEB SERVER CODE
	//mux := routes()
	//Creating a server
	// server := &http.Server{
	// 	Addr:    ":8080",
	// 	Handler: mux,
	// 	// Errorlog: ,
	// }
	// fmt.Println("Server Running in port 8080...")
	//Listen and Serve the server. If error, Fatal error.
	// if err := server.ListenAndServe(); err != nil {
	// 	log.Fatal(err)
	// }
}

func homepage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		log.Println("Homepage")
		log.Println("Error. Path Not Allowed.")
		http.Error(w, "Page Not Found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
		return
	}

	htmlTemplates := []string{
		"web/templates/main.html",
		"web/templates/main-bar.html",
		"web/templates/breadcrumb-nav.html",
		"web/templates/post-section.html",
		"web/templates/post-id.html",
		"web/templates/post.html",
		"web/templates/side-section.html",
		// "web/templates/registration.html",
		// "web/templates/login.html",
		// "web/templates/createPost.html",
	}

	tmpl, err := template.ParseFiles(htmlTemplates...)
	if err != nil {
		log.Printf("Error Parsing Template: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = tmpl.ExecuteTemplate(w, "main.html", "")
	if err != nil {
		log.Printf("Error Executing Template: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

}
