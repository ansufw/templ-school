package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"templ-school/internal/handlers"
)

func main() {
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)

	// Static files
	fileServer := http.FileServer(http.Dir("web/static/"))
	r.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	// Routes
	r.Get("/", handlers.HomeHandler)
	r.Get("/login", handlers.LoginHandler)
	r.Post("/login", handlers.LoginPostHandler)
	r.Get("/admin", handlers.AdminHome)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server starting on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
