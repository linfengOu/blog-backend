package server

import (
	"github.com/gin-gonic/gin"
	"github/linfengOu/blog-backend/config"
	"net/http"
)

type Server struct {
	Router *gin.Engine
}

func New() *Server {
	var s Server
	s.Router = gin.Default()

	// set router mapping
	s.Router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "Hello world",
		})
	})

	return &s
}

func (s *Server) Start() {
	// Start engine
	s.Router.Run(config.Config.HTTPServer.Port)
}
