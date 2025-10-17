package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/pc-configurator/components/internal/usecase"
)

func ComponentsRouter(r *chi.Mux, uc *usecase.UseCase) {
	r.Get("/components", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world"))
	})
}
