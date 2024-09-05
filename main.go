package main

import (
	"fmt"
	"log"
	"net/http"

	"gitea.koodsisu.fi/josepfrancescbenetmorella/literary-lions/dbaser"
	"gitea.koodsisu.fi/josepfrancescbenetmorella/literary-lions/handlers"
	"gitea.koodsisu.fi/josepfrancescbenetmorella/literary-lions/models"
)

func routes() *http.ServeMux {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./web/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", handlers.Homepage)
	mux.HandleFunc("/post/{id}", handlers.GetPost)
	// mux.HandleFunc("/search", handlers.Search)
	// mux.HandleFunc("/login", handlers.Login)
	// mux.HandleFunc("/register", handlers.Register)
	// mux.HandleFunc("/post/create", handlers.NewPost)
	// mux.HandleFunc("/profile", handlers.Profile)
	// mux.HandleFunc("/users/{username}/profile", handlers.Profile)
	// mux.HandleFunc("/notifications", handlers.Notifications)
	// mux.HandleFunc("/posts/liked", handlers.LikedPosts)
	// mux.HandleFunc("/posts/mined", handlers.MyPosts)
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
