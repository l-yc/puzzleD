package config

import (
  "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
  "github.com/go-chi/httplog"
	"github.com/qor/media/asset_manager"
	"fmt"
	models "github.com/l-yc/puzzled/app/models"
)

type Application struct {
	DB			*gorm.DB
	Router	chi.Router
}

func Setup(app *Application) {
  // Set up the database
  app.DB, _ = gorm.Open("sqlite3", "demo.db")
  app.DB.AutoMigrate(
		&models.User{},
		&models.Puzzle{},
		&asset_manager.AssetManager{},
	)

	// Set up the router
	app.Router = chi.NewRouter()
	logger := httplog.NewLogger("httplog-example", httplog.Options{
		JSON: true,
	})
	app.Router.Use(httplog.RequestLogger(logger))
	app.Router.Use(middleware.Heartbeat("/ping"))

	fmt.Println("configuration complete")
	//return &app
}
