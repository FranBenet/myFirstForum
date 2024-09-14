package session

import (
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"log"
	"net/http"
	"time"

	"gitea.koodsisu.fi/josepfrancescbenetmorella/literary-lions/dbaser"
	"gitea.koodsisu.fi/josepfrancescbenetmorella/literary-lions/models"
)

// CREATES A NEW SESSION, STORES IT IN THE DATABASE AND ADDS THEM IN THE RESPONSEWRITER
func CreateSession(w http.ResponseWriter, userID int, db *sql.DB) {
	sessionUUID, err := generateSessionUUID()
	if err != nil {
		log.Printf("Error Generating Sesssion ID: %v", err)
		return
	}

	expiration := time.Now().Add(30 * time.Minute)

	session := models.Session{
		Uuid:      sessionUUID,
		UserId:    userID,
		ExpiresAt: expiration,
	}

	//	Save session data in the db.
	_, err = dbaser.AddSession(db, session)
	if err != nil {
		log.Printf("Error creating session for the %v: %v", userID, err)
		return
	}
	//	Create a cookie
	cookie := &http.Cookie{
		Name:     "session_token",
		Value:    sessionUUID,
		Expires:  expiration,
		HttpOnly: true,
	}

	//	Set cookie in the response
	http.SetCookie(w, cookie)
}

// CLOSE A SESSION BY MAKING INVALID THE CURRENT SESSION
func EndSession(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	//	Get user's session from cookie
	cookie, err := r.Cookie("session_token")
	if err != nil {
		//	No Cookie found
		return
	}
	sessionUUID := cookie.Value
	_, err = dbaser.DeleteSession(db, sessionUUID)
	if err != nil {
		log.Printf("Error Deleting SessionID: %v from database. %v", sessionUUID, err)
		return
	}
	//	Expire the session for the cookie
	expiredCookie := &http.Cookie{
		Name:     "session_token",
		Value:    "",
		Expires:  time.Now(),
		HttpOnly: true,
	}
	http.SetCookie(w, expiredCookie)
}

// CREATES A SESSION USING CRYPTO/RAND. RETURNS A TOKEN STRING
func generateSessionUUID() (string, error) {
	//	Create a byte slice with 16 bits
	byteSlice := make([]byte, 16)

	//	Fill the slice with randome bytes
	_, err := rand.Read(byteSlice)
	if err != nil {
		log.Println(err)
		return "", err
	}

	//	Convert the bytes to hexadecimal string
	sessionID := hex.EncodeToString(byteSlice)

	return sessionID, nil
}

// GETS SESSION ID FROM THE REQUEST, CHECKS IF EXISTS IN THE DATABASE,
// CHECKS IF IS STILL VALID AND RETURNS A BOOLEAN INDICATED IF THERE IS A VALID SESSION TOGETHER WITH THE DATA
// func ValidateSession(r *http.Request) (*models.Session, bool) {
// 	//	Get user's session from cookie
// 	cookie, err := r.Cookie("session_token")
// 	if err != nil {
// 		//	No Cookie found
// 		return nil, false
// 	}

// 	sessionID := cookie.Value

// 	//	Get Session from db
// 	session, exists := getSession(sessionID)
// 	if !exists || session.ExpiresAt.Before(time.Now()) {
// 		return nil, false
// 	}

// 	return &session, true
// }

// CLEANS ALL EXPIRED SESSIONS IN THE DATABASE
// func CleanExpiredSession() {}

// SAVE SESSION DATA IN THE DATABASE
// func saveSession(session *models.Session) {}

// GETS DATA SESSION FROM THE DATABASE
// func getSession(sessionID string) (*Session, bool) {}

// DELETES THE SESSION ID FOR A USER IN THE DATABASE
// NOT SURE IF IT SHOULD DELETE THE SESSION ID ONLY OR THE WHOLE SESSION DATA
// func deleteSessionID(sessionId string) error {}
