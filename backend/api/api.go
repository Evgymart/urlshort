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

	var responseMessage []byte
	fullUrl, expiresAfter, err := validation.ValidateStore(request)
	if err != nil {
		responseMessage, _ = json.Marshal(core.HandleError(err))
		writeResponse(writer, responseMessage)
		return
	}

	data, err := core.BuildUrlData(fullUrl, expiresAfter)
	if err != nil {
		responseMessage, _ = json.Marshal(core.HandleError(err))
		writeResponse(writer, responseMessage)
		return
	}

	responseMessage, _ = json.Marshal(data)
	writeResponse(writer, responseMessage)
}

func writeResponse(writer http.ResponseWriter, responseMessage []byte) {
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(responseMessage)
}
