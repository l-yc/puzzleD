package config

import (
	"time"
  "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/go-chi/chi"
	"github.com/qor/media/asset_manager"
	"fmt"
)

// Define a GORM-backend model
type User struct {
  gorm.Model
}

// Define another GORM-backend model
type Puzzle struct {
  gorm.Model
  Name        string
  Description string
	//Attachment	oss.OSS `sql:"size:4294967295;"` //`sql:"size:4294967295;" media_library:"url:/system/{{class}}/{{primary_key}}/{{filename_with_hash}}"`
	//Attachment  media_library.MediaLibraryStorage `sql:"size:4294967295;" media_library:"url:/system/{{class}}/{{primary_key}}/{{column}}.{{extension}}"`
	ReleaseDate	time.Time
	AuthorId		int
	Solution		string
}

type Application struct {
	DB			*gorm.DB
	Router	chi.Router
}

func Setup(app *Application) {
	//app := new(Application)
  //app.DB, _ = gorm.Open("sqlite3", "demo.db")
  //app.DB.AutoMigrate(&User{}, &Puzzle{}, &asset_manager.AssetManager{},)

	//app.Router = new(chi.Router)

	//return app

  // Set up the database
  app.DB, _ = gorm.Open("sqlite3", "demo.db")
  app.DB.AutoMigrate(&User{}, &Puzzle{}, &asset_manager.AssetManager{},)

	app.Router = chi.NewRouter()

	fmt.Println("configuration complete")
	//return &app
}
