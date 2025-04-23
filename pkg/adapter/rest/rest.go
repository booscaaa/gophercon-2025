package rest

import (
	"fmt"
	"net/http"

	"github.com/booscaaa/hamburguer-go/pkg/di"
	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
)

func Initialize(database *sqlx.DB) {
	reviewController := di.NewReviewController(database)

	router := chi.NewRouter()
	router.Post("/review", reviewController.Save)
	router.Post("/review-alexa", reviewController.GetTop3Reviews)

	fmt.Println("Server running on port 3000")
	http.ListenAndServe(":3000", router)
}
