package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func NewRouter () *gin.Engine{
	router := gin.Default()


	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"ok": true,
			"status": "Healthy",
		})
	})

	return router
}
