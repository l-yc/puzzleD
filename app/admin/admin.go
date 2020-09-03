package admin

import (
  "github.com/qor/admin"
	"github.com/qor/media"
	"github.com/qor/media/asset_manager"
	"github.com/l-yc/puzzled/config"
	models "github.com/l-yc/puzzled/app/models"
)

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

  // Create resources from GORM-backend model
	user := Admin.AddResource(&models.User{})
	user.IndexAttrs("Name")
	user.EditAttrs("Name", "AuthoredPuzzles", "SolvedPuzzles")

	puzzle := Admin.AddResource(&models.Puzzle{})
	puzzle.Meta(&admin.Meta{Name: "Description", Config: &admin.RichEditorConfig{
		AssetManager: assetManager,
	}})
	//puzzleSolvedUserMeta := puzzle.Meta(&admin.Meta{Name: "AuthorID"})
	//puzzleSolvedUserResource := puzzleSolvedUserMeta.Resource
	//puzzleSolvedUserResource.IndexAttrs("Name")

	puzzle.IndexAttrs("Name", "Description", "Authors", "Solution")
	puzzle.EditAttrs("Name", "Description", "Authors", "Solution")

	app.Router.Mount("/admin", Admin.NewServeMux("/admin"))
}
