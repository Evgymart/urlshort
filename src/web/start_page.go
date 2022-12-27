package web

import (
	"html/template"
	"net/http"
	"path/filepath"
)

func startPage(writer http.ResponseWriter, request *http.Request) {
	// fmt.Fprintf(writer, "Hello, world!")

	filePrefix, _ := filepath.Abs("./web/templates")
	template := template.Must(template.ParseFiles(filePrefix + "/start_page.html"))
	template.Execute(writer, nil)
}
