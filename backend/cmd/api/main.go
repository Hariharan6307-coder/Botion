package main

import (
	"backend/internal/config"
	"backend/internal/database"
	"backend/internal/handler"
	"backend/internal/repository"
	"backend/internal/router"
	"backend/internal/service"
	"log"
	"net/http"
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

	blockRepo := repository.NewBlockRepository(db.Pool)

	blockService := service.NewBlockService(blockRepo)

	blockHandler := handler.NewBlockHandler(blockService)

	r := router.New()
	r.SetupRoutes(blockHandler)

	port := ":8080"
	log.Printf("Starting server on port %s", port)
	err = http.ListenAndServe(port, r.Handler())
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}

}