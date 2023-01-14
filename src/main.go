package main

import (
	"path/filepath"
	"shorturl/backend/app"
	"shorturl/backend/config"
)

func main() {
	root, _ := filepath.Abs(".")
	path := config.Path{
		AppRoot: root,
	}
	config.InitPath(&path)

	configPath, _ := filepath.Abs("../config")
	config.Initialize(configPath)

	app, err := app.NewApp()
	if err != nil {
		panic(err.Error())
		return
	}

	app.Start()
}
