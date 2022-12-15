package web

import (
	"fmt"
	"net/http"
	"shorturl/backend/core"
	"strings"
)

func InitWeb() {
	http.HandleFunc("/", webRoutine)
}

func webRoutine(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path == "/" {
		homePage(writer)
		return
	}

	code := strings.TrimPrefix(request.URL.Path, "/")
	redirectUrl, err := core.GetRedirectUrl(code)
	if err != nil {
		http.Error(writer, "Bad Request", http.StatusBadRequest)
		return
	}

	http.Redirect(writer, request, redirectUrl, http.StatusFound)
}

func homePage(writer http.ResponseWriter) {
	fmt.Fprintf(writer, "Welcome to the HomePage!")
}
