package main

import (
  "fmt"
	"time"
  "net/http"
  "github.com/qor/qor"
  "github.com/qor/admin"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/qor/qor/utils"
	"github.com/qor/media"
	"github.com/qor/media/asset_manager"
	_ "github.com/qor/media/oss"
	_ "github.com/qor/media/media_library"
	"github.com/qor/render"
	"github.com/go-chi/chi"
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

// Controller home controller
type Controller struct {
	View *render.Render
}

// Index home index page
func (ctrl Controller) Index(w http.ResponseWriter, req *http.Request) {
	ctrl.View.Execute("index", map[string]interface{}{}, req, w)
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
	Admin.AssetFS.RegisterPath("app/admin/views")

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

	r := chi.NewRouter()
	r.Mount("/admin", Admin.NewServeMux("/admin"))

	for _, path := range []string{"system", "javascripts", "stylesheets", "images"} {
		r.Mount(fmt.Sprintf("/%s/", path), utils.FileServer(http.Dir("public")))
	}

	// Mount API
	API := admin.New(&qor.Config{DB: DB})
  //user := API.AddResource(&User{})
  API.AddResource(&User{})
  API.AddResource(&Puzzle{})
	r.Mount("/api", API.NewServeMux("/api"))

	// Homepage
	controller := &Controller{
		View: render.New(&render.Config{}, "app/home/views"),
	}
	//r.Get("/app", controller.Index)

	r2 := chi.NewRouter()
	r2.Get("/", controller.Index)
	r2.Mount("/assets/", utils.FileServer(http.Dir("app/home/views/assets")))
	r.Mount("/app", r2)

	//application.Router.Get("/", controller.Index)

  fmt.Println("Listening on: 8080")
	http.ListenAndServe(":8080", r)

	//API.MountTo("/api", mux)

  //// Initalize an HTTP request multiplexer
  //mux := http.NewServeMux()

	//for _, path := range []string{"system", "javascripts", "stylesheets", "images"} {
	//	mux.Handle(fmt.Sprintf("/%s/", path), utils.FileServer(http.Dir("public")))
	//}

  //// Mount admin to the mux
  //Admin.MountTo("/admin", mux)

	//// Mount API
	//API := admin.New(&qor.Config{DB: DB})
  ////user := API.AddResource(&User{})
  //API.AddResource(&User{})
  //API.AddResource(&Puzzle{})

	//API.MountTo("/api", mux)

	//// home page
	//controller := &Controller{
	//	View: render.New(&render.Config{}, "app/home/views"),
	//}
	//mux.Handle("/app", http.Handler{
	//	ServeHTTP: controller.Index,
	//})
	////application.Router.Get("/", controller.Index)
	////application.Router.Get("/switch_locale", controller.SwitchLocale)

  //fmt.Println("Listening on: 8080")
  //http.ListenAndServe(":8080", mux)
}

