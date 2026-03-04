package notes

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


type NoteHandler struct {
	repo *NoteRepository
}

func NewNoteHandler(repo *NoteRepository) *NoteHandler {
	return &NoteHandler{
		repo: repo,
	}
}

func (h *NoteHandler) CreateNote(c *gin.Context) {
	// per req object from gin context
	//c.req, c.param, c.query, c.body, etc

	var req CreateNoteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	now := time.Now().UTC()

	note := Note{
		ID: primitive.NewObjectID(),
		Title: req.Title,
		Content: req.Content,
		Pinned: req.Pinned,
		CreatedAt: now,
		UpdatedAt: now,
	}

	createdNote, err := h.repo.CreateNote(c.Request.Context(), note)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create note"})
		return
	}

	c.JSON(http.StatusCreated, createdNote)
	}
