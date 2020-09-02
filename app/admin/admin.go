package admin

import (
	"time"
  "github.com/qor/admin"
  "github.com/jinzhu/gorm"
	"github.com/qor/media"
	"github.com/qor/media/asset_manager"
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
  // Initalize
  Admin := admin.New(&admin.AdminConfig{
		SiteName: "PuzzleD",
		DB: app.DB,
	})
	Admin.AssetFS.RegisterPath("app/admin/views")

	// Add Asset Manager, for rich editor
	assetManager := Admin.AddResource(&asset_manager.AssetManager{}, &admin.Config{Invisible: true})

	media.RegisterCallbacks(app.DB)
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

	app.Router.Mount("/admin", Admin.NewServeMux("/admin"))
}
