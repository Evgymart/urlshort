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

	fullUrl := request.FormValue("FullUrl")
	expiresAfter := request.FormValue("ExpiresAfter")
	expiresAfterInt, _ := strconv.ParseInt(expiresAfter, 0, 64)

	urlData, _ := json.Marshal(core.BuildUrlData(fullUrl, expiresAfterInt))
	writer.Header().Set("Content-Type", "application/json")
	writer.Write([]byte(urlData))
}
