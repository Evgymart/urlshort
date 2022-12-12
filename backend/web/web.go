package web

import (
	"fmt"
	"net/http"
)

func InitWeb() {
	http.HandleFunc("/", homePage)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
}
