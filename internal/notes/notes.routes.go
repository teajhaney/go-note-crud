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
		notesGroup.GET("/", handler.ListNotes)
		notesGroup.GET("/:id", handler.GetNoteByID)
		notesGroup.PATCH("/:id", handler.UpdateNoteByID)
		notesGroup.DELETE("/:id", handler.DeleteNoteByID)

	}

}
