package app

import (
	"net/http"
	"shorturl/backend/api"
	"shorturl/backend/config"
	"shorturl/backend/core"
	"shorturl/backend/db"
	"shorturl/backend/web"
)

type App struct {
	DB       *db.Database
	Mux      *http.ServeMux
	Settings *config.Settings
}

func NewApp(settings config.Settings) (*App, error) {
	redis, err := db.NewDatabase(settings.RedisAddr)
	if err != nil {
		return nil, err
	}

	mux := http.NewServeMux()
	api.InitApi(mux)
	web.InitWeb(mux)
	core.InitCore(redis)

	app := App{
		DB:       redis,
		Mux:      mux,
		Settings: &settings,
	}

	return &app, nil
}

func (app *App) Start() {
	http.ListenAndServe(app.Settings.HttpAddr, app.Mux)
}
