package api

import (
	"net/http"
)

func InitApi() {
	http.HandleFunc("/api/store", storeUrl)
}

func writeResponse(writer http.ResponseWriter, responseMessage []byte) {
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(responseMessage)
}
