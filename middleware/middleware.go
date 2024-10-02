package middleware

import (
	"context"
	"database/sql"
	"log"
	"net/http"

	"gitea.koodsisu.fi/josepfrancescbenetmorella/literary-lions/dbaser"
	"gitea.koodsisu.fi/josepfrancescbenetmorella/literary-lions/models"
)

type Middleware struct {
	db *sql.DB
}

func NewMiddleware(db *sql.DB) *Middleware {
	return &Middleware{db: db}
}

//	Middleware checks if there is a cookie with a valid session UUID.
// 	If there is not, userID is set to 0.
//	If there is, userID is set to its corresponding value.
//	Then embedds the userID in the context of the Request and passes the request to the relevant handler.

func (mw *Middleware) MiddlewareSession(requestedHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var userID int

		//	Get cookie from request
		sessionToken, err := r.Cookie("session_token")
		if err != nil {
			log.Println("Middleware: No cookie available:", err)
			userID = 0

		} else {

			//	Get the value of the session from the cookie
			sessionUUID := sessionToken.Value

			//	Check if userID has a valid session
			exists, err := dbaser.ValidSession(mw.db, sessionUUID)
			if err != nil {
				log.Printf("No valid session: %v", err)
				userID = 0

			} else if !exists {
				log.Printf("Session does not exist: %v", err)
				userID = 0

			} else {
				//	Get UserID from database
				userID, err = dbaser.SessionUser(mw.db, sessionUUID)
				if err != nil {
					log.Printf("Error getting user ID: %v", err)
					userID = 0
				}
			}
		}

		//	Create a new context with a key-value pair containing userID.
		ctx := context.WithValue(r.Context(), models.UserIDKey, userID)

		//	Send the request to the correct handler, using .WithContext() to include the context
		requestedHandler.ServeHTTP(w, r.WithContext(ctx))
	})

}

// FUNCTION PROVISIONAL TO SKIP MIDDLEWARE AND CHECK THAT ITS NOT GENERATING PROBLEMS
// func IsUserLoggedIn(db *sql.DB, sessionUUID string) (int, error) {

// 	// Check if userID has a valid session
// 	exists, err := dbaser.ValidSession(db, sessionUUID)
// 	fmt.Println(exists, err)
// 	if err != nil {
// 		return 0, err

// 	} else if !exists {
// 		err := errors.New("no valid session")
// 		return 0, err
// 	}

// 	// Get UserID from database
// 	userID, err := dbaser.SessionUser(db, sessionUUID)
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	log.Println("UserID:", userID, "has a valid session?")
// 	return userID, nil
// }
