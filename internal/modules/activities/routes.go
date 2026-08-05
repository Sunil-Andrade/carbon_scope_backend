package activities

import "github.com/go-chi/chi/v5"

func RegisterRoutes(r chi.Router) {

	r.Post("/", CreateActivity)

	r.Get("/", GetActivities)

	r.Get("/{id}", GetActivityByID)

	r.Put("/{id}", UpdateActivity)
	r.Patch("/{id}/verify", VerifyActivity)
	r.Delete("/{id}", DeleteActivity)
	r.Post("/upload", UploadProof)
}
