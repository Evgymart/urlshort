package web

import (
	"net/http"
	"shorturl/backend/core"
	"strings"
)

func InitWeb(mux *http.ServeMux) {
	mux.HandleFunc("/", webRoutine)
	mux.Handle("/assets/css/", http.FileServer(http.Dir(".")))
	mux.HandleFunc("/favicon.ico", func(writer http.ResponseWriter, reader *http.Request) {
		http.ServeFile(writer, reader, "assets/favicon.ico")
	})
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
