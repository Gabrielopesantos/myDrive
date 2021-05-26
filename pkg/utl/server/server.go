package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type Server struct {
	echo *echo.Echo
	db   *sqlx.DB
}

func NewServer(db *sqlx.DB) *Server {
	return &Server{echo: echo.New(), db: db}
}

func (s *Server) Run() error {

	server := &http.Server{
		Addr: "8888",
	}

	go func() {
		s.echo.Logger.Infof("Server is listening on port: 8888")
		if err := s.echo.StartServer(server); err != nil {
			s.echo.Logger.Fatalf("Error starting server", err)
		}
	}()

	if err := s.MapHandlers(s.echo); err != nil {
		return err
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return s.echo.Server.Shutdown(ctx)
}
