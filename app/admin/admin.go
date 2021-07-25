package admin

import (
  "github.com/qor/admin"
	"github.com/qor/media"
	"github.com/qor/media/asset_manager"
	"github.com/l-yc/puzzled/config"
	models "github.com/l-yc/puzzled/app/models"

	"github.com/qor/qor"
	"github.com/qor/auth"
	"github.com/qor/auth_themes/clean"

	"fmt"
	"net/http"
)

var Auth *auth.Auth;

type AdminAuth struct {}

func (AdminAuth) LoginURL(c *admin.Context) string {
    return "/auth/login"
}

func (AdminAuth) LogoutURL(c *admin.Context) string {
    return "/auth/logout"
}

func (AdminAuth) GetCurrentUser(c *admin.Context) qor.CurrentUser {
    currentUser := Auth.GetCurrentUser(c.Request)
    if currentUser != nil {
      qorCurrentUser, ok := currentUser.(qor.CurrentUser)
      if !ok {
        fmt.Printf("User %#v haven't implement qor.CurrentUser interface\n", currentUser)
      }
      return qorCurrentUser
    }
    return nil
}





//// Controller home controller
//type Controller struct {
//	View *render.Render
//}
//
//// Index home index page
//func (ctrl Controller) Puzzles(w http.ResponseWriter, req *http.Request) {
//	ctrl.View.Execute("login", map[string]interface{}{}, req, w)
//}






func Configure(app config.Application) {
	Auth = clean.New(&auth.Config{
		DB:         app.DB,
		ViewPaths:	[]string{"github.com/qor/auth_themes/clean/views"},
		// User model needs to implement qor.CurrentUser interface (https://godoc.org/github.com/qor/qor#CurrentUser) to use it in QOR Admin
		UserModel:  models.User{},
	})
	//app.Router.Route("/auth", Auth.NewServeMux().(http.HandlerFunc))

	//controller := &Controller{
	//	View: render.New(&render.Config{}, "app/home/views"),
	//}

	//app.Router.Route("/auth", func(r chi.Router) {
	//	r.Get("/login", controller.Puzzles)
	//	r.Post("/login", controller.Puzzles)
	//	r.Get("/logout", controller.Puzzles)
	//})

  // Initalize
  Admin := admin.New(&admin.AdminConfig{
		SiteName: "PuzzleD",
		DB: app.DB,
		Auth: &AdminAuth{},
	})
	Admin.AssetFS.RegisterPath("app/admin/views")

	// Add Asset Manager, for rich editor
	assetManager := Admin.AddResource(&asset_manager.AssetManager{}, &admin.Config{Invisible: true})
	media.RegisterCallbacks(app.DB)

  // Create resources from GORM-backend model
	//user := Admin.AddResource(&models.User{})
	Admin.AddResource(&models.User{})
	//user.IndexAttrs("Name")
	//user.EditAttrs("Name", "AuthoredPuzzles", "SolvedPuzzles")

	puzzle := Admin.AddResource(&models.Puzzle{})
	puzzle.Meta(&admin.Meta{Name: "Description", Config: &admin.RichEditorConfig{
		AssetManager: assetManager,
	}})
	//puzzleSolvedUserMeta := puzzle.Meta(&admin.Meta{Name: "AuthorID"})
	//puzzleSolvedUserResource := puzzleSolvedUserMeta.Resource
	//puzzleSolvedUserResource.IndexAttrs("Name")
	puzzle.Meta(&admin.Meta{Name: "Authors", Type: "select_many"})

	//puzzle.IndexAttrs("Name", "Description", "Authors", "Solution")
	//puzzle.EditAttrs("Name", "Description", "Authors", "Solution")

	app.Router.Mount("/admin", Admin.NewServeMux("/admin"))
}
