package http

import (
	"github.com/go-chi/chi/v5"
	http_server "github.com/pc-configurator/components/gen/http/components_v1/server"
	v1 "github.com/pc-configurator/components/internal/controller/http/v1"
	"github.com/pc-configurator/components/internal/usecase"
	"github.com/pc-configurator/components/pkg/logger"
)

func ComponentRouter(r *chi.Mux, uc *usecase.UseCase) {
	ver1 := v1.New(uc)

	r.Route("/api", func(r chi.Router) {
		r.Use(logger.Middleware)

		r.Route("/v1", func(r chi.Router) {
			mux := http_server.NewStrictHandler(ver1, []http_server.StrictMiddlewareFunc{})
			http_server.HandlerFromMux(mux, r)
		})
	})
}
