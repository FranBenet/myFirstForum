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
	mux.Handle("/post/", mw.MiddlewareSession(http.HandlerFunc(handler.GetPost)))
	mux.Handle("/post/create", mw.MiddlewareSession(http.HandlerFunc(handler.NewPost)))
	mux.Handle("/post/comment", mw.MiddlewareSession(http.HandlerFunc(handler.NewComment)))
	mux.Handle("/reaction", mw.MiddlewareSession(http.HandlerFunc(handler.Reaction)))

	mux.HandleFunc("/login", handler.Login)       //	No need to go through the middleware because never will be a cookie
	mux.HandleFunc("/register", handler.Register) //	No need to go through the middleware because never will be a cookie
	mux.HandleFunc("/logout", handler.Logout)     //	No need to go through the middleware because never will be a cookie

	// mux.Handle("/search", mw.MiddlewareSession(http.HandlerFunc(handler.Search)))
	// mux.Handle("/filter", mw.MiddlewareSession(http.HandlerFunc(handler.Filter)))
	// mux.Handle("/user/{username}/profile", mw.MiddlewareSession(http.HandlerFunc(handler.Profile)))

	mux.Handle("/profile", mw.MiddlewareSession(http.HandlerFunc(handler.Profile)))
	// mux.Handle("/profile/edit", mw.MiddlewareSession(http.HandlerFunc(handler.Profile)))
	// mux.Handle("/liked", mw.MiddlewareSession(http.HandlerFunc(handler.LikedPosts)))
	// mux.Handle("/myposts", mw.MiddlewareSession(http.HandlerFunc(handler.MyPosts)))

	// mux.HandleFunc("/post/", handler.GetPost)
	// mux.HandleFunc("/post/create", handler.NewPost)
	// mux.HandleFunc("/post/comment", handler.NewComment)
	// mux.HandleFunc("/reaction", handler.Reaction)
	// mux.HandleFunc("/", handler.Homepage)

	// mux.HandleFunc("/login", handler.Login)
	// mux.HandleFunc("/register", handler.Register)
	// mux.HandleFunc("/logout", handler.Logout)

	mux.HandleFunc("/search", handler.Search)
	mux.HandleFunc("/filter", handler.Filter)
	mux.HandleFunc("/user/", handler.Profile)
	mux.HandleFunc("/404", handler.NotFound)

	// mux.HandleFunc("/profile", handler.Profile)
	mux.HandleFunc("/profile/edit", handler.Profile)
	mux.HandleFunc("/liked", handler.LikedPosts)
	mux.HandleFunc("/myposts", handler.MyPosts)

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
	// valid, err := dbaser.ValidateLogin(db, "PageTurner@example.com", "SbdfbWE345$")
	// uuid, err := dbaser.GenerateUuid(32)
	// fmt.Println(uuid, err)
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
