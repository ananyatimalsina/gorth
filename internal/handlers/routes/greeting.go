package routes

import (
	"gortth/web/templates"
	"net/http"
	"time"
)

func GreetingHandler(w http.ResponseWriter, r *http.Request) {
	greetings := []string{
		"Hello from the Gorth Stack! 🚀",
		"Greetings from your Go server! 👋",
		"Welcome to the future of web development!",
		"HTMX + Go = Pure Magic ✨",
		"Server-side rendering at its finest!",
	}

	currentTime := time.Now()
	greeting := greetings[currentTime.Second()%len(greetings)]

	templates.Greeting(greeting, currentTime.Format("15:04:05")).Render(r.Context(), w)
}
