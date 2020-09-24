package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func mapRouter(r *gin.Engine) {
	r.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "Hello world",
		})
	})
}
