package server

import (
	"net/http"
	"notes-api/internal/notes"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)


func NewRouter (database *mongo.Database) *gin.Engine{
	router := gin.Default()


	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"ok": true,
			"status": "Healthy",
		})
	})

	// Register notes routes
	notes.RegisterRoutes(router, database)
	return router
}
