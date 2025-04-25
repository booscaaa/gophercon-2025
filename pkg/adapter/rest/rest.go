package rest

import (
	"fmt"
	"net/http"

	"github.com/booscaaa/hamburguer-go/pkg/di"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/jmoiron/sqlx"
)

func Initialize(database *sqlx.DB) {
	reviewController := di.NewReviewController(database)

	router := chi.NewRouter()

	router.Use(cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	}).Handler)

	router.Post("/review", reviewController.Save)
	router.Post("/review/alexa", reviewController.GetTop3Reviews)
	router.Get("/review/count", reviewController.Count)

	fmt.Println("Server running on port 3000")
	http.ListenAndServe(":3000", router)
}
