package home

import (
	"net/http"
	_ "github.com/qor/qor/utils"
	"github.com/qor/render"
	"github.com/l-yc/puzzled/config"
	"github.com/go-chi/chi"
	"fmt"
	"os"
	"path"
	"strings"
)

// Controller home controller
type Controller struct {
	View *render.Render
}

// Index home index page
func (ctrl Controller) Puzzles(w http.ResponseWriter, req *http.Request) {
	ctrl.View.Execute("puzzles", map[string]interface{}{}, req, w)
}

func RouterFileServer(dir string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		/**
			First, we need to determine the paths that we have already routed through.
			These are stored in rctx.RoutePatterns as "<fragment>/*".
		*/
		rctx := chi.RouteContext(r.Context())
		prefix := ""
		for _, route := range rctx.RoutePatterns {
			prefix = path.Join(prefix, strings.TrimSuffix(route, "/*"))
		}

		/**
		 Now, we can determine the path relative to the current router.
		 The implementation below references the implementation used in qor/utils
		 which disables directory listings.
		*/
		p := path.Join(string(dir), strings.TrimPrefix(r.URL.Path, prefix))
		if f, err := os.Stat(p); err == nil && !f.IsDir() {
			fmt.Println(os.Stat(p))
			http.ServeFile(w, r, p)
			return
		}

		http.NotFound(w, r)
	})
}

func Configure(app config.Application) {
	controller := &Controller{
		View: render.New(&render.Config{}, "app/home/views"),
	}

	app.Router.Route("/app", func(r chi.Router) {
		r.Get("/puzzles", controller.Puzzles)
		r.Route("/assets/", func(r chi.Router) {
			r.Get("/*", RouterFileServer("app/home/assets"))
		})
	})
}
