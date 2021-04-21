package server

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	Engine *gin.Engine
	// Database *database.Database
	// Config   appCfg.AppConfig
}

func New() (*Server, error) {
	engine, err := initEngine()
	if err != nil {
		return nil, err
	}

	return &Server{
		Engine: engine,
	}, nil
}

func (s *Server) Run(addr string) error {
	return s.Engine.Run(addr)
}

func initEngine() (*gin.Engine, error) {

	engine := gin.Default()
	// engine.Use(gin.Logger())
	// engine.Use(gin.Recovery())

	return engine, nil
}
