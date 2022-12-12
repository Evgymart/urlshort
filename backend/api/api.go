package api

import "net/http"

func InitApi() {
	http.HandleFunc("/api/json", handleJson)
}

func handleJson(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	response := `{"status": "OK"}`
	writer.Write([]byte(response))
}
