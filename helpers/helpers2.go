package helpers

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func CheckQueryParameters(r *http.Request) {
	// Parse the query parameters from the URL
	query := r.URL.Query()

	// Get possible values for parameters

	errorMessage := query.Get("error")
	fmt.Println("Query Parameter Error:", errorMessage)

	successMessage := query.Get("success")
	fmt.Println("Query Parameter Success:", successMessage)

	page := r.URL.Query().Get("page")
	fmt.Println("Query Parameter Page:", page)

	unescapedError, err := url.QueryUnescape(errorMessage)
	if err != nil || unescapedError == "" {
		log.Println("Error unsecaping Error:", err)
	}
	data.Metadata.Error = unescapedError

	unescapedSuccess, err := url.QueryUnescape(successMessage)
	if err != nil || unescapedSuccess == "" {
		log.Println("Error unsecaping Success:", err)
	}
}
