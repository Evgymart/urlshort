package main

import (
	"shorturl/backend/app"
	"shorturl/backend/config"
)

var (
	HttpAddr  = ":80"
	RedisAddr = "redis:6379"
)

func main() {
	settings := config.Settings{
		HttpAddr:  HttpAddr,
		RedisAddr: RedisAddr,
	}

	app, err := app.NewApp(settings)
	if err != nil {
		panic(err.Error())
		return
	}

	app.Start()
}
