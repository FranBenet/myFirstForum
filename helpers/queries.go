package helpers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"gitea.koodsisu.fi/josepfrancescbenetmorella/literary-lions/dbaser"
	"gitea.koodsisu.fi/josepfrancescbenetmorella/literary-lions/models"
)

func ProfilePageData(db *sql.DB, userId int) (models.ProfilePageData, error) {
	var profilePageData models.ProfilePageData
	user, err := dbaser.UserById(db, userId)
	if err != nil {
		return models.ProfilePageData{}, err
	}

	profilePageData.User = models.User{Avatar: user.Avatar, Name: user.Name, Email: user.Email}
	profilePageData.Metadata.LoggedIn = true

	return profilePageData, nil
}

// Adds a new "error" or "success" message to a given url (referer). You need to specify the key for the query parameter (error or success).
func AddQueryMessage(referer string, key string, message string) string {

	//	Convert URL to url.url format
	refererURL, err := url.Parse(referer)
	if err != nil {
		log.Println("Failed to parse referer:", err)
		refererURL = &url.URL{Path: "/"}
	}

	//	Delete any queries that the url may have
	refererURL.RawQuery = ""

	// Get the current query values from the referer URL
	query := refererURL.Query()

	// Add the error to the query.
	query.Set(key, message)

	// Set the updated query back to the referer URL
	refererURL.RawQuery = query.Encode()

	// Convert the modified referer URL back to a string
	finalURL := refererURL.String()

	return finalURL
}

// Deletes the "error" and "success" messages from the query parameters
func CleanQueryMessages(referer string) string {

	//	Convert URL to url.url format
	refererURL, err2 := url.Parse(referer)
	if err2 != nil {
		log.Println("Failed to parse referer:", err2)
		refererURL = &url.URL{Path: "/"}
	}

	// Get all query values
	query := refererURL.Query()

	// Remove the "error" and "success" query parameter
	query.Del("error")
	query.Del("success")

	// Set the updated query back to the referer URL
	refererURL.RawQuery = query.Encode()

	// Convert the modified referer URL back to a string
	cleanURL := refererURL.String()

	return cleanURL
}

// Gets the "error" and "success" messages from a query parameter
func GetQueryMessages(r *http.Request) (string, string, error) {

	// Parse the query parameters from the URL
	query := r.URL.Query()

	errorMessage := query.Get("error")
	successMessage := query.Get("success")

	if errorMessage != "" {
		unescapedError, err := url.QueryUnescape(errorMessage)
		if err != nil || unescapedError == "" {
			return "", "", err
		}
		return unescapedError, "", nil

	} else if successMessage != "" {
		unescapedSuccess, err := url.QueryUnescape(successMessage)
		if err != nil || unescapedSuccess == "" {
			return "", "", err
		}
		return "", unescapedSuccess, nil
	}
	return "", "", nil
}

// Gets the "page" from a query parameter
func GetQueryPage(r *http.Request) (int, error) {
	// Parse the query parameters from the URL
	pageNumber := r.URL.Query().Get("page")

	page, err := strconv.Atoi(pageNumber)
	if err != nil {
		return 0, err
	}
	return page, nil
}

// Gets the page number from a request
func GetRefererPage(r *http.Request) (int, error) {
	// Parse the query parameters from the URL
	referer := r.Referer()
	refererURL, err := url.Parse(referer)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	pageNumber := refererURL.Query().Get("page")
	if pageNumber == "" {
		return 1, nil
	}

	page, err := strconv.Atoi(pageNumber)
	if err != nil {
		log.Println("Error converting from string to int:", err)
		return 0, err
	}

	return page, nil
}

func GetQueryFilter(r *http.Request, key string) (string, error) {
	// Parse the query parameters from the URL
	queryParams := r.URL.Query()

	// Example: Get a specific key's value if it exists
	if value, ok := queryParams[key]; ok {
		return value[0], nil

	} else {
		err := fmt.Errorf("no filter available:%s", key)
		return "", err
	}

}
