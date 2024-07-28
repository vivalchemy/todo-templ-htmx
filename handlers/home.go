package handlers

import (
	"net/http"

	"github.com/vivalchemy/ecom/routes"
)

// Handler functions.
func Home(w http.ResponseWriter, r *http.Request) {
	routes.Home().Render(r.Context(), w)
}
