package repository

import (
	"context"

	"github.com/booscaaa/hamburguer-go/internal/core/domain"
	"github.com/booscaaa/hamburguer-go/internal/core/dto"
	"github.com/jmoiron/sqlx"
)

type reviewDatabaseRepository struct {
	database *sqlx.DB
}

// Count implements domain.ReviewDatabaseRepository.
func (repository *reviewDatabaseRepository) Count(ctx context.Context) (int, error) {
	var count int
	err := repository.database.GetContext(ctx, &count, "SELECT COUNT(id) FROM review;")
	return count, err
}

// Fetch implements domain.ReviewDatabaseRepository.
func (repository *reviewDatabaseRepository) Fetch(ctx context.Context) ([]domain.Review, error) {
	var reviews []domain.Review
	err := repository.database.SelectContext(ctx, &reviews, "SELECT * FROM review;")
	return reviews, err
}

// Save implements domain.ReviewDatabaseRepository.
func (repository *reviewDatabaseRepository) Save(ctx context.Context, input dto.Review) error {
	stmt, err := repository.database.PrepareContext(
		ctx,
		"INSERT INTO review (name, description) VALUES ($1, $2);",
	)
	if err != nil {
		return err
	}
	_, err = stmt.ExecContext(
		ctx,
		input.Name,
		input.Description,
	)
	return err
}

func NewReviewDatabaseRepository(database *sqlx.DB) domain.ReviewDatabaseRepository {
	return &reviewDatabaseRepository{
		database: database,
	}
}
