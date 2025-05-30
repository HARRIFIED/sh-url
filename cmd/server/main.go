package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	"url-shortner-go/api/v1/handler"
	"url-shortner-go/config"
	"url-shortner-go/internal/repository/postgresql"
	"url-shortner-go/internal/service"
)

func main() {
	cfg := config.Load()
	db, err := sql.Open("postgres", cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	repo := postgresql.NewURLRepo(db)
	svc := service.NewURLService(repo)
	urlHandler := handler.NewURLHandler(svc)

	r := gin.Default()
	// Add middleware, auth placeholder
	// r.Use(middleware.Auth())

	v1 := r.Group("/api/v1")
	urlHandler.Register(v1)

	log.Printf("starting server on %s", cfg.Port)
	r.Run(cfg.Port)
}
