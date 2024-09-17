package middleware

import (
	"context"
	"database/sql"
	"net/http"

	"gitea.koodsisu.fi/josepfrancescbenetmorella/literary-lions/dbaser"
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

		//	Check if there is a cookie in the request
		sessionToken, err := r.Cookie("session_token")
		if err != nil {
			//	Create a new context with a key-value pair: loggedIn = false
			ctx := context.WithValue(r.Context(), "loggedIn", false)

			//	Send the request to the correct handler, using .WithContext() to include the context
			requestedHandler.ServeHTTP(w, r.WithContext(ctx))
		}

		//	Get the value of the session from the cookie
		sessionUUID := sessionToken.Value

		//	Get UserID from database
		userID, err := dbaser.SessionUser(mw.db, sessionUUID)
		if err != nil {
			// HANDLE ERROR
		}

		//	Check if userID has a valid session
		exists, err := dbaser.ValidSession(mw.db, userID)
		if err != nil {
			//	HANDLE ERROR
		} else if !exists {
			//	Create a new context with a key-value pair: loggedIn = false
			ctx := context.WithValue(r.Context(), "loggedIn", false)

			//	Send the request to the correct handler, using .WithContext() to include the context
			requestedHandler.ServeHTTP(w, r.WithContext(ctx))
		}

		//	Create a new context with a key-value pair: loggedIn = true
		ctx := context.WithValue(r.Context(), "loggedIn", true)

		//	Create a new context with a key-value pair containing userID.
		ctx = context.WithValue(ctx, "userID", userID)

		//	Send the request to the correct handler, using .WithContext() to include the context
		requestedHandler.ServeHTTP(w, r.WithContext(ctx))
	})

}
