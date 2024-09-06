package middleware

import (
	"net/http"
	"time"
)

// THE MIDDLEWARE CHECKS IF THE REQUEST HAS AN OPEN SESSION.
// IF EXISTS -> REQUEST IS SEND TO HANDLERS.
// IF DOES NOT EXISTS -> REQUEST IS REDIRECTED TO /LOGIN PAGE.
func MidlewareSession(requestedHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session_token")
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		sessionID := cookie.Value

		session, exists := dbaser.getSession(sessionID)
		if !exists || session.ExpiresAt.Before(time.Now()) {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		requestedHandler.ServeHTTP(w, r)
	})

}
