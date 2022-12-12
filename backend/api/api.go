package api

import (
	"encoding/json"
	"net/http"
	"shorturl/backend/core"
	"strconv"
)

func InitApi() {
	http.HandleFunc("/api/store", storeUrl)
}

func storeUrl(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	writer.Header().Set("Content-Type", "application/json")

	fullUrl := request.FormValue("FullUrl")
	expiresAfter := request.FormValue("ExpiresAfter")
	expiry, _ := strconv.ParseInt(expiresAfter, 0, 64)
	data := core.UrlData{fullUrl, "http://localhost:8080/adfgttbrwetfwef", expiry}
	response, _ := json.Marshal(data)

	writer.Write([]byte(response))
}
