package httpserver

import (
	"context"
	"errors"
	"net"
	"net/http"
	"time"

	"github.com/pc-configurator/components/pkg/logger"
)

type Config struct {
	Port string `default:"8080" envconfig:"HTTP_PORT"`
}

type Server struct {
	server *http.Server
}

func New(handler http.Handler, c Config) *Server {
	httpServer := &http.Server{
		Handler:      handler,
		ReadTimeout:  20 * time.Second,
		WriteTimeout: 20 * time.Second,
		Addr:         net.JoinHostPort("", c.Port),
	}

	s := &Server{
		server: httpServer,
	}

	go s.start()

	logger.Info("http server: started on port: " + c.Port)

	return s
}

func (s *Server) start() {
	err := s.server.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		logger.Error(err, "httpserver.start")
	}
}

func (s *Server) Close() {
	ctx, cancel := context.WithTimeout(context.Background(), 25*time.Second)
	defer cancel()

	err := s.server.Shutdown(ctx)
	if err != nil {
		logger.Error(err, "httpserver.Shutdown")
	}

	logger.Info("httpserver: closed")
}
