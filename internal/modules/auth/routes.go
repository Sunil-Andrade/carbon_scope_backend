package auth

import "github.com/go-chi/chi/v5"

func RegisterRoutes(r chi.Router) {

	r.Post("/register", Register)

}
