package session

import (
	"crypto/rand"
	"encoding/hex"
	"literary-lions-fran/models"
	"log"
	"net/http"
	"time"
)

// CREATES A NEW SESSION, STORES IT IN THE DATABASE AND ADDS THEM IN THE RESPONSEWRITER
func CreateSession(w http.ResponseWriter, userID int) {
	sessionID, err := generateSessionID()
	if err != nil {
		log.Printf("Error Generating Sesssion ID: %v", err)
		return
	}

	expiration := time.Now().Add(1 * time.Hour)

	session := &models.Session{
		Id:        sessionID,
		UserID:    userID,
		ExpiresAt: expiration,
	}

	//	Save session data in the db
	saveSession(session)

	//	Create a cookie
	cookie := &http.Cookie{
		Name:     "session_token",
		Value:    sessionID,
		Expires:  expiration,
		HttpOnly: true,
	}

	//	Set cookie in the response
	http.SetCookie(w, cookie)
}

// GETS SESSION ID FROM THE REQUEST, CHECKS IF EXISTS IN THE DATABASE,
// CHECKS IF IS STILL VALID AND RETURNS A BOOLEAN INDICATED IF THERE IS A VALID SESSION TOGETHER WITH THE DATA
func ValidateSession(r *http.Request) (*models.Session, bool) {
	//	Get user's session from cookie
	cookie, err := r.Cookie("session_token")
	if err != nil {
		//	No Cookie found
		return nil, false
	}

	sessionID := cookie.Value

	//	Get Session from db
	session, exists := getSession(sessionID)
	if !exists || session.ExpiresAt.Before(time.Now()) {
		return nil, false
	}

	return &session, true
}

// CLOSE A SESSION BY MAKING INVALID THE CURRENT SESSION
func EndSession() {
	//	Get user's session from cookie
	cookie, err := r.Cookie("session_token")
	if err != nil {
		//	No Cookie found
		return nil, false
	}

	sessionID := cookie.Value

	err = deleteSessionID(sessionID)
	if err != nil{
		log.Printf("Error Deleting SessionID from database: %v", err)
		return
	}

	//	Expire the session for the cookie
	expiredCookie := &http.Cookie{
		Name: "session_token",
		Value: "",
		Expires: time.Now().Add(-1*time.Day())
		HttpOnly: true,
	}

	http.SetCookie(w,expiredCookie)
}

// CREATES A SESSION USING CRYPTO/RAND. RETURNS A TOKEN STRING
func generateSessionID() (string, error) {
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


// CLEANS ALL EXPIRED SESSIONS IN THE DATABASE
func CleanExpiredSession() {}

// SAVE SESSION DATA IN THE DATABASE
func saveSession(session *models.Session) {}

// GETS DATA SESSION FROM THE DATABASE
func getSession(sessionID string) (*Session, bool) {}

//	DELETES THE SESSION ID FOR A USER IN THE DATABASE
//	NOT SURE IF IT SHOULD DELETE THE SESSION ID ONLY OR THE WHOLE SESSION DATA
func deleteSessionID (sessionId string) error {}
