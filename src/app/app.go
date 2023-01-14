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
	DB  *db.Database
	Mux *http.ServeMux
}

func NewApp() (*App, error) {
	redis, err := db.NewDatabase(config.GetDatabaseAddr())
	if err != nil {
		return nil, err
	}

	mux := http.NewServeMux()
	return &App{
		DB:  redis,
		Mux: mux,
	}, nil
}

func (app *App) Start() {
	api.InitApi(app.Mux)
	web.InitWeb(app.Mux)
	core.InitCore(app.DB)
	http.ListenAndServe(config.GetServerAddr(), app.Mux)
}
