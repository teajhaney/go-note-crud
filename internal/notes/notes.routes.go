package notes

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)



func RegisterRoutes(router *gin.Engine, db *mongo.Database	) {

	//create a new note repository and handler
	repo := NewRepo(db)
	handler := NewNoteHandler(repo)

	//group routes under /notes
	notesGroup := router.Group("/notes")
	{
		notesGroup.POST("/", handler.CreateNote)
		// Additional routes for GET, PUT, DELETE can be added here
	}
}
