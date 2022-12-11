package main

import (
	"fmt"
	"log"
	"net/http"
	"shorturl/backend/db"
)

var (
	ListenAddr = ":80"
	RedisAddr  = "redis:6379"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(ListenAddr, nil))
}

func main() {
	db.NewDatabase(RedisAddr)
	handleRequests()
}
