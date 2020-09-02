package api

import (
	"time"
  "github.com/qor/qor"
  "github.com/qor/admin"
  "github.com/jinzhu/gorm"
	"github.com/l-yc/puzzled/config"
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

func Configure(app config.Application) {
	// Mount API
	API := admin.New(&qor.Config{DB: app.DB})
	//user := API.AddResource(&User{})
	API.AddResource(&User{})
	API.AddResource(&Puzzle{})
	app.Router.Mount("/api", API.NewServeMux("/api"))
}
