package main

import (
	"backend/internal/config"
	"backend/internal/database"
	"log"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	db, err := database.NewConnectionPool(cfg.Database)
	if err != nil {
		log.Fatal("Failed to connect to database");
	}
	defer db.Close()

	log.Println("Database connection established successfully")
}