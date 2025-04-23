package domain

import (
	"context"

	"github.com/booscaaa/hamburguer-go/internal/core/dto"
)

type Item struct {
	ID         int     `db:"id"`
	Name       string  `db:"name"`
	Price      float64 `db:"price"`
	InsertedAt string  `db:"inserted_at"`
}

type ItemScraperGateway interface {
	ScrapeItems(context.Context) ([]Item, error)
}

type ItemDatabaseRepository interface {
	Save(context.Context, Item) (*Item, error)
	FetchAllFromLastSync(context.Context) ([]Item, error)
}

type ItemUseCase interface {
	Sync(context.Context) error
	FetchAllFromLastSync(context.Context) ([]Item, error)
	GetRecomendation(context.Context) (*string, error)
}

type ItemLLMGateway interface {
	GenerateRecomendation(context.Context, []dto.Tool, []Item) ([]dto.Function, error)
}
