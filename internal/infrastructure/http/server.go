package http

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Server struct {
	engine *gin.Engine
	cfg    ServerConfig
}

type ServerConfig struct {
	Port           string
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	HandlerTimeout time.Duration
}

func NewServer(cfg ServerConfig, handler http.Handler) *Server {
	server := gin.Default()
	// server.Use(gin.Recovery())
	// server.Use(gin.Logger())
	server.Any("/*any", gin.WrapH(handler))
	return &Server{
		engine: server,
		cfg:    cfg,
	}
}

func (s *Server) Start() error {
	return s.engine.Run(":" + s.cfg.Port)
}
