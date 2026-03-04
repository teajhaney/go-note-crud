package notes

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//data access layer for notes collection in MongoDB


type NoteRepository struct {
	// Fields for MongoDB connection and operations would go here
	collection *mongo.Collection
}

func  NewRepo(db *mongo.Database) *NoteRepository{
	return &NoteRepository{
		collection: db.Collection("notes"),
	}
}


//CREATE NOTE
func (r *NoteRepository) CreateNote(ctx context.Context, note Note) (Note,error) {

	//ctx-> parent contxt for managing the lifecycle of the operation, allowing for cancellation and timeouts
	//childCtx, cancel := context.WithTimeout(ctx, 5*time.Second) creates a new context with a timeout of 5 seconds. This ensures that if the note creation operation takes longer than 5 seconds, it will be automatically cancelled to prevent hanging operations.

	childCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	// Insert the note into MongoDB
	_, err := r.collection.InsertOne(childCtx, note)


	// Implementation for creating a note in MongoDB
	if err != nil {
		return Note{}, fmt.Errorf("failed to create note: %v", err)
	}
	// Return the created note with its ID set
	return note, nil
}
