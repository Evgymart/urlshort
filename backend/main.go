package main

import (
	"log"
	"net/http"
	"shorturl/backend/api"
	"shorturl/backend/core"
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
	redis, err := db.NewDatabase(RedisAddr)
	if err != nil {
		panic("Database error")
		return
	}

	api.InitApi()
	web.InitWeb()
	core.InitCore(redis)
	handleRequests()
}
