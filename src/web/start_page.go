package web

import (
	"html/template"
	"net/http"
	"shorturl/backend/config"
)

func startPage(writer http.ResponseWriter, request *http.Request) {
	root := config.GetPath().AppRoot
	template := template.Must(template.ParseFiles(root + "/web/templates/start_page.html"))
	template.Execute(writer, nil)
}
