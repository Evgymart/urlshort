package web

import (
	"fmt"
	"net/http"
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
}

func homePage(writer http.ResponseWriter) {
	fmt.Fprintf(writer, "Welcome to the HomePage!")
}
