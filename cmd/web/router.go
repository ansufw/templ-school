package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"templ-school/internal/handlers"
)

func newRouter(app *application) http.Handler {
	r := chi.NewRouter()

	// Middleware
	r.Use(Logger(app.log))

	// Static files
	fileServer := http.FileServer(http.Dir("web/static/"))
	r.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	// Routes
	r.Get("/", handlers.HomeHandler)
	r.Get("/login", handlers.LoginHandler)
	r.Post("/login", handlers.LoginPostHandler)
	r.Get("/admin", handlers.AdminHome)

	return r
}
