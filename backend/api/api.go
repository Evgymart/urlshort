package api

import (
	"encoding/json"
	"net/http"
	"shorturl/backend/core"
	"shorturl/backend/validation"
)

func InitApi() {
	http.HandleFunc("/api/store", storeUrl)
}

func storeUrl(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	fullUrl, expiresAfter, err := validation.ValidateStore(request)
	if err != nil {
		http.Error(writer, "Bad Request", http.StatusBadRequest)
		return
	}

	data, err := core.BuildUrlData(fullUrl, expiresAfter)
	if err != nil {
		http.Error(writer, "Bad Request", http.StatusBadRequest)
		return
	}

	urlData, err := json.Marshal(data)
	if err != nil {
		http.Error(writer, "Bad Request", http.StatusBadRequest)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.Write([]byte(urlData))
}
