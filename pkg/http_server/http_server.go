package http_server

import (
	"context"
	"net"
	"net/http"
	"time"

	"github.com/pc-configurator/components/pkg/logger"
)

type Config struct {
	Port string `envconfig:"HTTP_PORT" default:"8080"`
}

type Server struct {
	server *http.Server
	notify chan error
}

func (s *Server) start() {
	s.notify <- s.server.ListenAndServe()
	close(s.notify)
}

func (s *Server) Notify() <-chan error {
	return s.notify
}

func (s *Server) Close() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := s.server.Shutdown(ctx)
	if err != nil {
		logger.Error(err, "s.server.Shutdown")
	}

	logger.Info("HTTP Server closed")
}

func New(handler http.Handler, port string) *Server {
	httpServer := &http.Server{
		Handler:      handler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		Addr:         net.JoinHostPort("", port),
	}

	s := &Server{
		server: httpServer,
		notify: make(chan error, 1),
	}

	go s.start()

	logger.Info("HTTP Server started on port: " + port)

	return s
}
