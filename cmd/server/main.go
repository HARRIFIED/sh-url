package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	"url-shortner-go/api/v1/handler"
	"url-shortner-go/config"
	"url-shortner-go/internal/repository/postgresql"
	"url-shortner-go/internal/service"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	cfg := config.Load()
	db, err := sql.Open("postgres", cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	repo := postgresql.NewURLRepo(db)
	urlService := service.NewURLService(repo)
	urlHandler := handler.NewURLHandler(urlService)

	router := gin.Default()
	// Serve landing page
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	router.SetHTMLTemplate(tmpl)
	router.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", nil)
	})

	// API endpoints
	router.POST("/shorten", urlHandler.Shorten)
	router.GET("/:code", urlHandler.Resolve)

	log.Printf("starting server on %s", cfg.Port)
	router.Run(cfg.Port)
}
