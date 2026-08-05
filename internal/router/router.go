package router

import (
	"carbon/internal/modules/auth"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func SetupRouter() *chi.Mux {

	r := chi.NewRouter()

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server is running"))
	})

	r.Route("/api/v1/auth", func(r chi.Router) {
		auth.RegisterRoutes(r)
	})

	return r
}
