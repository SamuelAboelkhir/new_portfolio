package main

import (
	"net/http"

	"github.com/SamuelAboelkhir/new_portfolio/internal"
	"github.com/SamuelAboelkhir/new_portfolio/views"
)

func (cfg *apiConfig) handlerHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	component := views.Home()
	err := internal.GenerateHtml(component, cfg.filePathRoot+"index")
	if err != nil {
		http.Error(w, "Failed to generate html", http.StatusInternalServerError)
	}

	if err := component.Render(r.Context(), w); err != nil {
		http.Error(w, "Failed to render page", http.StatusInternalServerError)
	}
}
