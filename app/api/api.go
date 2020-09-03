package api

import (
  "github.com/qor/qor"
  "github.com/qor/admin"
	"github.com/l-yc/puzzled/config"
	models "github.com/l-yc/puzzled/app/models"
)

func Configure(app config.Application) {
	// Mount API
	API := admin.New(&qor.Config{DB: app.DB})
	API.AssetFS.RegisterPath("app/admin/views")

	user := API.AddResource(&models.User{})
	user.IndexAttrs("ID", "Name", "SolvedPuzzles")

	puzzle := API.AddResource(&models.Puzzle{})
	puzzle.IndexAttrs("ID", "Name", "Description", "ReleaseDate", "Authors")

	app.Router.Mount("/api", API.NewServeMux("/api"))
}
