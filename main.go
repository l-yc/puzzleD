package main

import (
  "fmt"
  "net/http"
	"github.com/qor/qor/utils"
	adminapp "github.com/l-yc/puzzled/app/admin"
	apiapp "github.com/l-yc/puzzled/app/api"
	homeapp "github.com/l-yc/puzzled/app/home"
	config "github.com/l-yc/puzzled/config"
)

func main() {
	var app config.Application
	config.Setup(&app)

	adminapp.Configure(app)
	apiapp.Configure(app)
	homeapp.Configure(app)

	// Public Assets
	for _, path := range []string{"system", "javascripts", "stylesheets", "images"} {
		app.Router.Mount(fmt.Sprintf("/%s/", path), utils.FileServer(http.Dir("public")))
	}

  fmt.Println("Listening on: 8080")
	http.ListenAndServe(":8080", app.Router)
}

