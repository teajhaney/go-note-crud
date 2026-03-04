package main

import (
	"fmt"
	"log"
	"notes-api/internal/config"
	"notes-api/internal/db"
	"notes-api/internal/server"
)


func main (){

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Config Error: %v", err)
	}

	client, database, err := db.Connect(*cfg)
	if err != nil {
		log.Fatal("Database Connection Error",err)
	}
defer func() {
	if err := db.Disconnect(client); err != nil {
		log.Printf("Database Disconnection Error: %v", err)
	}
}()

	router := server.NewRouter(database)

	port := fmt.Sprintf(":%s", cfg.Port)
	if err := router.Run(port); err != nil {
		log.Fatal("Server Error",err)
	}

}
