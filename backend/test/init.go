package test

import (
	"shorturl/backend/app"
	"shorturl/backend/config"
	"testing"

	"github.com/alicebob/miniredis/v2"
)

var (
	HttpAddrTest = "127.0.0.1:8000"
)

func initTest(t *testing.T) *app.App {
	redisServer := miniredis.RunT(t)
	settings := config.Settings{
		HttpAddr:  HttpAddrTest,
		RedisAddr: redisServer.Addr(),
	}

	app, err := app.NewApp(settings)
	if err != nil {
		t.Errorf("Error during init: %s", err.Error())
		panic(err)
	}

	return app
}
