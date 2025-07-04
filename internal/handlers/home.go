package handlers

import (
	"net/http"

	"templ-school/web/templates/pages"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	component := pages.Home()
	component.Render(r.Context(), w)
}
