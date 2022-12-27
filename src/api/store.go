package api

import (
	"encoding/json"
	"net/http"
	"shorturl/backend/core"
	"shorturl/backend/validation"
)

func storeUrl(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	fullUrl, expiresAfter, err := validation.ValidateStore(request)
	var responseMessage []byte
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
