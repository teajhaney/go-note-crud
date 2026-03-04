package notes

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)


type Note struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Title string `bson:"title" json:"title"`
	Content string `bson:"content" json:"content"`
	Pinned bool `bson:"pinned" json:"pinned"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}


type CreateNoteRequest struct {
	Title string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	Pinned bool `json:"pinned"`
}


type UpdateNoteRequest struct {
	Title *string `json:"title,omitempty"`
	Content *string `json:"content,omitempty"`
	Pinned *bool `json:"pinned,omitempty"`
}
