package handlers

import (
	"gorth/internal/handlers/routes"
	"html/template"
	"net/http"
)

func LoadHanders(router *http.ServeMux) {
	// Index page
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("web/templates/index.html"))
		tmpl.Execute(w, nil)
	})

	api := http.NewServeMux()
	api.Handle("/api/", http.StripPrefix("/api", router))

	router.HandleFunc("GET /greeting", routes.GreetingHandler)
	router.HandleFunc("GET /time", routes.TimeHandler)
	router.HandleFunc("GET /api/stats", routes.StatsHandler)
}
