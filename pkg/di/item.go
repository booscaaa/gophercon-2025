package di

import (
	"github.com/booscaaa/hamburguer-go/internal/core/domain"
	"github.com/booscaaa/hamburguer-go/internal/core/usecase"
	"github.com/booscaaa/hamburguer-go/internal/infra/gateway"
	"github.com/booscaaa/hamburguer-go/internal/infra/repository"
	"github.com/jmoiron/sqlx"
)

func NewItemUseCase(database *sqlx.DB) domain.ItemUseCase {
	return usecase.NewItemUseCase(
		gateway.NewItemLLM(),
		gateway.NewItemScraper(),
		repository.NewItemDatabase(database),
	)
}
