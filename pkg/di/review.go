package di

import (
	"github.com/booscaaa/hamburguer-go/internal/core/domain"
	"github.com/booscaaa/hamburguer-go/internal/core/usecase"
	"github.com/booscaaa/hamburguer-go/internal/infra/controller"
	"github.com/booscaaa/hamburguer-go/internal/infra/gateway"
	"github.com/booscaaa/hamburguer-go/internal/infra/repository"
	"github.com/jmoiron/sqlx"
)

func NewReviewController(database *sqlx.DB) domain.ReviewController {
	return controller.NewReviewController(
		usecase.NewReviewUseCase(
			gateway.NewReviewLLM(),
			repository.NewReviewDatabaseRepository(database),
		),
	)
}
