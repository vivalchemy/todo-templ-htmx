package handlers

import (
	"net/http"

	"github.com/vivalchemy/ecom/routes"
)

func Info(w http.ResponseWriter, r *http.Request) {
	routes.Info().Render(r.Context(), w)
}
