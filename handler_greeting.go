package main

import (
	"net/http"

	"github.com/SamuelAboelkhir/new_portfolio/components"
	"github.com/SamuelAboelkhir/new_portfolio/internal"
)

func (cfg *apiConfig) handleGreeting(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	component := components.Greeting()
	err := internal.GenerateHtml(component, cfg.filePathRoot+"greeting")
	if err != nil {
		http.Error(w, "Failed to generate html", http.StatusInternalServerError)
	}

	if err := component.Render(r.Context(), w); err != nil {
		http.Error(w, "could not render greeting", http.StatusInternalServerError)
	}
}
