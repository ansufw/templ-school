package handlers

import (
	"net/http"

	"templ-school/web/templates/pages"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	component := pages.Login("")
	component.Render(r.Context(), w)
}

func LoginPostHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		component := pages.Login("Invalid form data")
		component.Render(r.Context(), w)
		return
	}

	email := r.FormValue("email")
	password := r.FormValue("password")

	// Simple validation for demo
	if email == "" || password == "" {
		component := pages.Login("Email and password are required")
		component.Render(r.Context(), w)
		return
	}

	// Mock authentication
	if email == "admin@example.com" && password == "password" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	component := pages.Login("Invalid credentials")
	component.Render(r.Context(), w)
}
