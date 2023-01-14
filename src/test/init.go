package test

import (
	"net"
	"path/filepath"
	"shorturl/backend/app"
	"shorturl/backend/config"
	"strconv"
	"testing"

	"github.com/alicebob/miniredis/v2"
)

func setConfig(t *testing.T) {
	redisServer := miniredis.RunT(t)
	dbHost, dbPort, _ := net.SplitHostPort(redisServer.Addr())

	conf := config.Config{}
	conf.Server.Host = "127.0.0.1"
	conf.Server.Port = 8000
	conf.Database.Host = dbHost
	conf.Database.Port, _ = strconv.Atoi(dbPort)
	config.SetConfig(&conf)
}

func initTest(t *testing.T) *app.App {
	setConfig(t)
	root, _ := filepath.Abs("..")
	path := config.Path{
		AppRoot: root,
	}

	config.InitPath(&path)
	app, err := app.NewApp()
	if err != nil {
		t.Errorf("Error during init: %s", err.Error())
		panic(err)
	}

	return app
}

func buildUrl(addr string) string {
	return "http://" + addr + "/"
}

func buildUrlPath(addr string, path string) string {
	return "http://" + addr + "/" + path
}
