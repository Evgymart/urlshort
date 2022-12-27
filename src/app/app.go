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
	return &App{
		DB:       redis,
		Mux:      mux,
		Settings: &settings,
	}, nil
}

func (app *App) Start() {
	api.InitApi(app.Mux)
	web.InitWeb(app.Mux)
	core.InitCore(app.DB)
	http.ListenAndServe(app.Settings.HttpAddr, app.Mux)
}
