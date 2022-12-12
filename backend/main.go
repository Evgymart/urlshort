package main

import (
	"log"
	"net/http"
	"shorturl/backend/api"
	"shorturl/backend/db"
	"shorturl/backend/web"
)

var (
	ListenAddr = ":80"
	RedisAddr  = "redis:6379"
)

func handleRequests() {
	log.Fatal(http.ListenAndServe(ListenAddr, nil))
}

func main() {
	db.NewDatabase(RedisAddr)
	api.InitApi()
	web.InitWeb()
	handleRequests()
}
