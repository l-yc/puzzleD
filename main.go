package main

import (
  "fmt"
	"time"
  "net/http"
  _ "github.com/qor/qor"
  "github.com/qor/admin"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/qor/qor/utils"
	"github.com/qor/media"
	"github.com/qor/media/asset_manager"
	_ "github.com/qor/media/oss"
	_ "github.com/qor/media/media_library"
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

func main() {
  // Set up the database
  DB, _ := gorm.Open("sqlite3", "demo.db")
  DB.AutoMigrate(&User{}, &Puzzle{}, &asset_manager.AssetManager{},)

  // Initalize
  Admin := admin.New(&admin.AdminConfig{
		SiteName: "PuzzleD",
		DB: DB,
	})

	// Add Asset Manager, for rich editor
	assetManager := Admin.AddResource(&asset_manager.AssetManager{}, &admin.Config{Invisible: true})

	media.RegisterCallbacks(DB)
	// Add Media Library
	//Admin.AddResource(&media_library.MediaLibrary{}, &admin.Config{Menu: []string{"Site Management"}})

  // Create resources from GORM-backend model
  Admin.AddResource(&User{})
	puzzle := Admin.AddResource(&Puzzle{})
	//puzzle.Meta(&admin.Meta{Name: "Description", Type: "rich_editor"})
	puzzle.Meta(&admin.Meta{Name: "Description", Config: &admin.RichEditorConfig{
		AssetManager: assetManager,
		//Plugins:      []admin.RedactorPlugin{
		//	{Name: "medialibrary", Source: "/admin/assets/javascripts/qor_redactor_medialibrary.js"},
		//	{Name: "table", Source: "/admin/assets/javascripts/qor_kindeditor.js"},
		//},
		//Settings: map[string]interface{}{
		//	"medialibraryUrl": "/public/system/",
		//},
	}})

  // Initalize an HTTP request multiplexer
  mux := http.NewServeMux()

	for _, path := range []string{"system", "javascripts", "stylesheets", "images"} {
		mux.Handle(fmt.Sprintf("/%s/", path), utils.FileServer(http.Dir("public")))
	}

  // Mount admin to the mux
  Admin.MountTo("/admin", mux)

  fmt.Println("Listening on: 8080")
  http.ListenAndServe(":8080", mux)
}

