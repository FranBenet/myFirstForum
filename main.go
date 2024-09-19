package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"gitea.koodsisu.fi/josepfrancescbenetmorella/literary-lions/dbaser"
	"gitea.koodsisu.fi/josepfrancescbenetmorella/literary-lions/handlers"
	"gitea.koodsisu.fi/josepfrancescbenetmorella/literary-lions/middleware"
)

func routes(db *sql.DB) *http.ServeMux {
	mux := http.NewServeMux()

	// Create handler with the database connection
	handler := handlers.NewHandler(db)

	// Create middleware with the database connection
	mw := middleware.NewMiddleware(db)

	fileServer := http.FileServer(http.Dir("./web/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.Handle("/", mw.MiddlewareSession(http.HandlerFunc(handler.Homepage)))
	mux.Handle("/post/{id}", mw.MiddlewareSession(http.HandlerFunc(handler.GetPost)))
	mux.Handle("/post/create", mw.MiddlewareSession(http.HandlerFunc(handler.NewPost)))
	mux.Handle("/reaction", mw.MiddlewareSession(http.HandlerFunc(handler.Reaction)))

	mux.HandleFunc("/login", handler.Login)       //	No need to go through the middleware because never will be a cookie
	mux.HandleFunc("/register", handler.Register) //	No need to go through the middleware because never will be a cookie

	mux.Handle("/search", mw.MiddlewareSession(http.HandlerFunc(handler.Search)))
	mux.Handle("/filter", mw.MiddlewareSession(http.HandlerFunc(handler.Filter)))
	mux.Handle("/user/{username}/profile", mw.MiddlewareSession(http.HandlerFunc(handler.Profile)))

	mux.Handle("/profile", mw.MiddlewareSession(http.HandlerFunc(handler.Profile)))
	mux.Handle("/profile/edit", mw.MiddlewareSession(http.HandlerFunc(handler.Profile)))
	mux.Handle("/liked", mw.MiddlewareSession(http.HandlerFunc(handler.LikedPosts)))
	mux.Handle("/myposts", mw.MiddlewareSession(http.HandlerFunc(handler.MyPosts)))

	return mux
}

func main() {
	// dbaser.InitDb()
	// dbaser.PopulateDb()
	// user := models.User{"madrabbit@matrix.com", "whiterabbit", "Rz_;*$78)"}
	db, err := dbaser.DbHandle("./forum.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	valid, err := dbaser.ValidateLogin(db, "PageTurner@example.com", "SbdfbWE345$")
	fmt.Println(valid, err)
	// emailOk, err := dbaser.UsernameExists(db, "PageTurner")
	// fmt.Println(emailOk, err)
	//	PROVISIONAL STARTING WEB SERVER CODE
	mux := routes(db)
	// Creating a server
	server := &http.Server{
		Addr:     ":8080",
		Handler:  mux,
		ErrorLog: log.New(os.Stderr, "server: ", log.LstdFlags),
	}
	fmt.Println("Server Running in port 8080...")
	// Listen and Serve the server. If error, Fatal error.
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
