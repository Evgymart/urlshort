package web

import (
	"net/http"
	"shorturl/backend/core"
	"strings"
)

func InitWeb(mux *http.ServeMux) {
	mux.HandleFunc("/", webRoutine)
}

func webRoutine(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path == "/" {
		startPage(writer, request)
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
