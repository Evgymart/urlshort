package main

import (
	"path/filepath"
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

	root, _ := filepath.Abs(".")
	path := config.Path{
		AppRoot: root,
	}

	configPath, _ := filepath.Abs("../config")
	config.Initialize(configPath)

	config.InitPath(&path)
	app, err := app.NewApp(settings)
	if err != nil {
		panic(err.Error())
		return
	}

	app.Start()
}
