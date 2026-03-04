package notes

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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


//List all notes
func (r *NoteRepository) ListNotes(ctx context.Context) ([]Note, error) {

	childCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	// Implementation for listing all notes from MongoDB
	filter := bson.M{} // Empty filter to retrieve all documents
	cursor, err := r.collection.Find(childCtx, filter)

	if err != nil {
		return []Note{}, fmt.Errorf("failed to list notes: %v", err)
	}

	// close the cursor after processing the results to free up resources and avoid potential memory leaks. The defer statement ensures that the cursor is closed when the function returns, regardless of whether it returns successfully or encounters an error.
	defer cursor.Close(childCtx)

	var notes []Note
	if err := cursor.All(childCtx, &notes); err != nil {
		return []Note{}, fmt.Errorf("failed to decode notes: %v", err)
	}
	return notes, nil
	
}


// GET NOTE BY ID retrieves a single note by its MongoDB _id value. Mongo stores the
// identifier as an ObjectID, so callers pass the hex representation and we
// convert it before querying. Failing to convert (or providing a bogus id)
// results in an error; the handler cantranslate that into a 404/400 response.
func (r *NoteRepository) GetNoteByID(ctx context.Context, id string) (Note, error) {

	childCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// convert the hex string to an ObjectID; this is what the collection uses for
	// the `_id` field. If the incoming id isn't valid hex the conversion fails
	// immediately rather than issuing a query that will always miss.
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return Note{}, fmt.Errorf("invalid id format: %v", err)
	}

	filter := bson.M{"_id": oid}
	var note Note
	err = r.collection.FindOne(childCtx, filter, options.FindOne()).Decode(&note)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return Note{}, fmt.Errorf("note not found")
		}
		return Note{}, fmt.Errorf("failed to get note: %v", err)
	}
	return note, nil
}


//UPDATE NOTE BY ID
func (r *NoteRepository) UpdateNoteByID(ctx context.Context, id primitive.ObjectID, note UpdateNoteRequest) (Note, error) {

	childCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{
		"title": note.Title,
		"content": note.Content,
		"pinned": note.Pinned,
		"updated_at": time.Now().UTC(),
	}}
	var updatedNote Note
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	err := r.collection.FindOneAndUpdate(childCtx, filter, update, opts).Decode(&updatedNote)
	if err != nil {
		return Note{}, fmt.Errorf("failed to update note: %v", err)
	}
	return updatedNote, nil

}


//DELETE NOTE BY ID
func (r *NoteRepository) DeleteNoteByID(ctx context.Context, id primitive.ObjectID) (bool,error) {

	childCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": id}
	response, err := r.collection.DeleteOne(childCtx, filter, options.Delete())
	if err != nil {
		return false, fmt.Errorf("failed to delete note: %v", err)
	}
	if response.DeletedCount == 0 {
		return false, fmt.Errorf("note not found")
	}
	return true, nil

}
