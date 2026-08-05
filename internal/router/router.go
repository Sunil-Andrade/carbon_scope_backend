package router

import (
	"carbon/internal/modules/activities"
	"carbon/internal/modules/auth"
	"carbon/internal/modules/wallet"
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

	r.Route("/api/v1/activities", func(r chi.Router) {
		activities.RegisterRoutes(r)
	})

	r.Route("/api/v1/wallet", func(r chi.Router) {

		wallet.RegisterRoutes(r)

	})

	fs := http.FileServer(http.Dir("./uploads"))
	r.Handle("/uploads/*", http.StripPrefix("/uploads/", fs))

	return r
}
