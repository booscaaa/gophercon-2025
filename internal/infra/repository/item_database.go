package repository

import (
	"context"

	"github.com/booscaaa/hamburguer-go/internal/core/domain"
	"github.com/jmoiron/sqlx"
)

type itemDatabase struct {
	database *sqlx.DB
}

// FetchAllFromLastSync implements domain.ItemDatabaseRepository.
func (repository *itemDatabase) FetchAllFromLastSync(ctx context.Context) ([]domain.Item, error) {
	var items []domain.Item
	query := `
		SELECT id, name, price, inserted_at
		FROM item
		WHERE inserted_at > NOW() - INTERVAL '1 day';
	`
	err := repository.database.SelectContext(ctx, &items, query)
	if err != nil {
		return nil, err
	}

	return items, nil
}

// Save implements domain.ItemDatabaseRepository.
func (repository *itemDatabase) Save(ctx context.Context, item domain.Item) (*domain.Item, error) {
	var itemCreated domain.Item
	query := `
		INSERT INTO item (name, price)
		VALUES ($1, $2)
		RETURNING id, name, price, inserted_at;
	`
	err := repository.database.QueryRowxContext(
		ctx,
		query,
		item.Name, item.Price,
	).StructScan(&itemCreated)
	if err != nil {
		return nil, err
	}

	return &itemCreated, nil
}

func NewItemDatabase(database *sqlx.DB) domain.ItemDatabaseRepository {
	return &itemDatabase{
		database: database,
	}
}
