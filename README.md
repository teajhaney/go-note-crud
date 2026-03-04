# go-note-crud

Simple notes API written in Go.

## Setup

1. Copy `.env.example` to `.env` in the repository root and update values.
   ```sh
   cp .env.example .env
   ```

2. Ensure you run commands from the project root (not from `cmd/api`). The loader
   uses [godotenv](https://github.com/joho/godotenv) which searches the working
   directory, so running from subfolders can prevent your `.env` from being
   picked up.

3. Build or run the server:
   ```sh
   cd cmd/api
   go run main.go
   # or from root
   go run ./cmd/api
   ```

## Environment variables

The application expects the following variables to be set:

* `MONGODB_URI` - MongoDB connection string
* `MONGO_DB_NAME` - name of the database
* `PORT` - port the HTTP server listens on

These are automatically loaded from `.env` by `internal/config`.
