package server

import (
	"github.com/gin-gonic/gin"
	"github/linfengOu/write-backend/config"
)

type Server struct {
	Router *gin.Engine
}

func New() *Server {
	var s Server
	s.Router = gin.Default()

	// set router mapping
	mapRouter(s.Router)

	return &s
}

func (s *Server) Start() {
	// Start engine
	s.Router.Run(config.Get(config.ServicePort))
}
