package notes

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


type NoteHandler struct {
	repo *NoteRepository
}

func NewNoteHandler(noteRepository *NoteRepository) *NoteHandler {
	return &NoteHandler{
		repo: noteRepository,
	}
}


//CREATE NOTE
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

	c.JSON(http.StatusCreated, gin.H{"message": "Note created successfully", "note": createdNote})
	}



	//List all notes
func (h *NoteHandler) ListNotes(c *gin.Context) {
	notes, err := h.repo.ListNotes(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list notes"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Notes fetched successfully", "notes": notes, "length": len(notes)})
}


//Get a single note by ID
func (h *NoteHandler) GetNoteByID(c *gin.Context) {
	id := c.Param("id")
	note, err := h.repo.GetNoteByID(c.Request.Context(), id)
	if err != nil {
		// the repository returns a concrete message for not‑found vs other
		if err.Error() == "note not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "note not found"})
			return
		}
		// anything else is unexpected; include the text for easier debugging
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Note fetched successfully", "note": note})
}

//Update a note by ID
func (h *NoteHandler) UpdateNoteByID(c *gin.Context) {
	id:= c.Param("id")
	iod, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid note ID"})
		return
	}

	var req UpdateNoteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	note := UpdateNoteRequest{
		Title: req.Title,
		Content: req.Content,
		Pinned: req.Pinned,

	}
	//call the repo to update the note
	updatedNote, err := h.repo.UpdateNoteByID(c.Request.Context(), iod, note)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Note not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update note"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Note updated successfully", "note": updatedNote})
}



//Delete a note by ID
func (h *NoteHandler) DeleteNoteByID(c *gin.Context) {
	id := c.Param("id")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid note ID"})
		return
	}

	//call the repo to delete the note
	success, err := h.repo.DeleteNoteByID(c.Request.Context(), oid)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Note not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete note"})
		return
	}

	if !success {
		c.JSON(http.StatusNotFound, gin.H{"error": "Note not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Note deleted successfully"})
}


