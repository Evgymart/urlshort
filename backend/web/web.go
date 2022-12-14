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
	fmt.Printf("Req: %s %s\n", request.Host, request.URL.Path)
	if request.URL.Path == "/" {
		homePage(writer)
		return
	}

	code := strings.TrimPrefix(request.URL.Path, "/")
	redirectUrl, err := core.GetRedirectUrl(code)
	fmt.Printf("Redirect URL: %s\n", redirectUrl)
	if err != nil {
		http.Error(writer, "Bad Request", http.StatusBadRequest)
		return
	}

	http.Redirect(writer, request, redirectUrl, http.StatusFound)
}

func homePage(writer http.ResponseWriter) {
	fmt.Fprintf(writer, "Welcome to the HomePage!")
}
