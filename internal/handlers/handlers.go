package handlers

import (
	"gortth/internal/handlers/routes"
	"gortth/web/templates"
	"net/http"

	"github.com/a-h/templ"
)

func LoadHanders(router *http.ServeMux) {
	// Index page
	router.Handle("/", templ.Handler(templates.Index()))

	apiRouter := http.NewServeMux()

	apiRouter.HandleFunc("GET /greeting", routes.GreetingHandler)
	apiRouter.HandleFunc("GET /time", routes.TimeHandler)
	apiRouter.HandleFunc("GET /stats", routes.StatsHandler)

	router.Handle("/api/", http.StripPrefix("/api", apiRouter))
}
