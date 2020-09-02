package home

import (
	"net/http"
	"github.com/qor/qor/utils"
	"github.com/go-chi/chi"
	"github.com/qor/render"
	"github.com/l-yc/puzzled/config"
)

// Controller home controller
type Controller struct {
	View *render.Render
}

// Index home index page
func (ctrl Controller) Index(w http.ResponseWriter, req *http.Request) {
	ctrl.View.Execute("index", map[string]interface{}{}, req, w)
}

func Configure(app config.Application) {
	controller := &Controller{
		View: render.New(&render.Config{}, "app/home/views"),
	}

	r := chi.NewRouter()
	r.Get("/", controller.Index)
	app.Router.Mount("/assets/", utils.FileServer(http.Dir("app/home/views/assets")))

	app.Router.Mount("/app", r)
}
