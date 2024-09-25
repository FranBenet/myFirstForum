package helpers

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

func GetQueryMessages(r *http.Request) (string, string, error) {

	// Parse the query parameters from the URL
	query := r.URL.Query()

	errorMessage := query.Get("error")
	fmt.Println("Query Parameter Error:", errorMessage)

	successMessage := query.Get("success")
	fmt.Println("Query Parameter Success:", successMessage)

	if errorMessage != "" {
		unescapedError, err := url.QueryUnescape(errorMessage)
		if err != nil || unescapedError == "" {
			log.Println("Error unsecaping Error:", err)
			return "", "", err
		}
		return unescapedError, "", nil

	} else if successMessage != "" {
		unescapedSuccess, err := url.QueryUnescape(successMessage)
		if err != nil || unescapedSuccess == "" {
			log.Println("Error unsecaping Success:", err)
			return "", "", err
		}
		return "", unescapedSuccess, nil
	}
	return "", "", nil
}

func GetQueryPage(r *http.Request) (int, error) {
	// Parse the query parameters from the URL

	pageNumber := r.URL.Query().Get("page")
	fmt.Println("Query Parameter Page:", pageNumber)

	page, err := strconv.Atoi(pageNumber)
	if err != nil {
		log.Println("Error converting from string to int:", err)
		return 0, err
	}
	return page, nil
}

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
