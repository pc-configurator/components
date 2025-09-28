package http_router

import (
	"github.com/go-chi/chi/v5"
	"github.com/pc-configurator/components/internal/controller/http_router/v1"
	"github.com/pc-configurator/components/internal/usecase"
)

func ComponentRouter(r *chi.Mux, uc *usecase.UseCase) {
	ver1 := v1.New(uc)

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/component/{id}", ver1.GetComponent)
	})
}
