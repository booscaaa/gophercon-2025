package domain

import (
	"context"
	"net/http"

	"github.com/booscaaa/hamburguer-go/internal/core/dto"
)

type Review struct {
	ID          int    `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
	InsertedAt  string `db:"inserted_at"`
}

type ReviewDatabaseRepository interface {
	Fetch(context.Context) ([]Review, error)
	Save(context.Context, dto.Review) error
}

type ReviewUseCase interface {
	GetTop3Reviews(context.Context) (*string, error)
	Save(context.Context, dto.Review) error
}

type ReviewController interface {
	GetTop3Reviews(http.ResponseWriter, *http.Request)
	Save(http.ResponseWriter, *http.Request)
}

type ReviewLLMGateway interface {
	GetTop3Reviews(context.Context, []Review) (*string, error)
}
