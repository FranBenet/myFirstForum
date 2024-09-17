package middleware

import (
	"context"
	"database/sql"
	"log"
	"net/http"

	"gitea.koodsisu.fi/josepfrancescbenetmorella/literary-lions/dbaser"
	"gitea.koodsisu.fi/josepfrancescbenetmorella/literary-lions/models"
)

//	THE MIDDLEWARE CHECKS IF THE REQUEST HAS AN OPEN SESSION.
//	THEN ADDS A KEY-VALUE INFORMATION IN THE CONTEXT() OF THE REQUEST SPECIFYING
//	TRUE -> SESSION EXISTS AND IS VALID
//	FALSE -> SESSION DOES NOT EXISTS OR IS NOT VALID
//	THE REQUEST WITH THE NEW CONTEXT INFORMATION IS PASSED TO THE NEXT HANDLER

type Middleware struct {
	db *sql.DB
}

func NewMiddleware(db *sql.DB) *Middleware {
	return &Middleware{db: db}
}

func (mw *Middleware) MiddlewareSession(requestedHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		//	Get cookie from request
		sessionToken, err := r.Cookie("session_token")
		if err != nil {
			//	If no cookie available, create a new context with a key-value pair: userID = 0
			ctx := context.WithValue(r.Context(), models.UserIDKey, 0)

			//	Send the request to the correct handler, using .WithContext() to include the context
			requestedHandler.ServeHTTP(w, r.WithContext(ctx))
		}

		//	Get the value of the session from the cookie
		sessionUUID := sessionToken.Value

		//	Get UserID from database
		userID, err := dbaser.SessionUser(mw.db, sessionUUID)
		if err != nil {
			// LOG ERROR
			log.Printf("Error: %v", err)

			//	If fail to get the user of the session create a new context with a key-value pair: userID = 0
			ctx := context.WithValue(r.Context(), models.UserIDKey, 0)

			//	Send the request to the correct handler, using .WithContext() to include the context
			requestedHandler.ServeHTTP(w, r.WithContext(ctx))
		}

		//	Check if userID has a valid session
		exists, err := dbaser.ValidSession(mw.db, userID)
		if err != nil {
			// LOG ERROR
			log.Printf("Error: %v", err)

			//	If fail to get the user of the session create a new context with a key-value pair: userID = 0
			ctx := context.WithValue(r.Context(), models.UserIDKey, 0)

			//	Send the request to the correct handler, using .WithContext() to include the context
			requestedHandler.ServeHTTP(w, r.WithContext(ctx))

		} else if !exists {
			// LOG ERROR
			log.Printf("Error: No session valid for this user. %v", err)

			//	Create a new context with a key-value pair: loggedIn = false
			ctx := context.WithValue(r.Context(), models.UserIDKey, 0)

			//	Send the request to the correct handler, using .WithContext() to include the context
			requestedHandler.ServeHTTP(w, r.WithContext(ctx))
		}

		//	Create a new context with a key-value pair containing userID.
		ctx := context.WithValue(r.Context(), models.UserIDKey, userID)

		//	Send the request to the correct handler, using .WithContext() to include the context
		requestedHandler.ServeHTTP(w, r.WithContext(ctx))
	})

}
