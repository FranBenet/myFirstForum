package main

import (
	"fmt"
	"log"
	"net/http"

	"gitea.koodsisu.fi/josepfrancescbenetmorella/literary-lions/dbaser"
	"gitea.koodsisu.fi/josepfrancescbenetmorella/literary-lions/helpers"
)

func routes() *http.ServeMux {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./web/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// mux.HandleFunc("/", handlers.Homepage)
	// mux.HandleFunc("/post/{id}", handlers.GetPost)
	// mux.HandleFunc("/login", handlers.Login)
	// mux.HandleFunc("/register", handlers.Register)
	// mux.HandleFunc("/search", handlers.Search)

	// mux.Handle("/post/create", middleware.MidlewareSession(http.HandlerFunc(handlers.NewPost)))
	// mux.Handle("/profile", middleware.MidlewareSession(http.HandlerFunc(handlers.Profile)))
	// mux.Handle("/user/{username}/profile", middleware.MidlewareSession(http.HandlerFunc(handlers.Profile)))
	// mux.Handle("/posts/liked", middleware.MidlewareSession(http.HandlerFunc(handlers.LikedPosts)))
	// mux.Handle("/posts/myposts", middleware.MidlewareSession(http.HandlerFunc(handlers.MyPosts)))

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
	// session := models.Session{UserId: 2, Uuid: "ssdlkd;-.29384FBERF098234", ExpiresAt: time.Now().Add(1 * time.Hour)}
	md, _ := helpers.MainPageData(db, "")
	fmt.Println(md)
	// fmt.Println(dbaser.ValidSession(db, "ssdlkd;-.29384FBERF098234"))
	// fmt.Println(dbaser.TrendingPosts(db, 3))

	//	PROVISIONAL STARTING WEB SERVER CODE
	// mux := routes()
	// Creating a server
	// server := &http.Server{
	// 	Addr:    ":8080",
	// 	Handler: mux,
	// 	// Errorlog: ,
	// }
	// fmt.Println("Server Running in port 8080...")
	// Listen and Serve the server. If error, Fatal error.
	// if err := server.ListenAndServe(); err != nil {
	// 	log.Fatal(err)
	// }
}
