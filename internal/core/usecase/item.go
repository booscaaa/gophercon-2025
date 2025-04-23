package usecase

import (
	"context"
	"fmt"

	"github.com/booscaaa/hamburguer-go/internal/core/domain"
	"github.com/booscaaa/hamburguer-go/pkg/adapter/llm"
)

type itemUseCase struct {
	itemLLMGateway         domain.ItemLLMGateway
	itemScraperGateway     domain.ItemScraperGateway
	itemDatabaseRepository domain.ItemDatabaseRepository
}

// GetRecomendation implements domain.ItemUseCase.
func (usecase *itemUseCase) GetRecomendation(ctx context.Context) (*string, error) {
	items, err := usecase.itemDatabaseRepository.FetchAllFromLastSync(ctx)
	if err != nil {
		return nil, err
	}

	functions, err := usecase.itemLLMGateway.GenerateRecomendation(ctx, llm.GetRecomendationTools(), items)
	if err != nil {
		return nil, err
	}

	for _, function := range functions {
		switch function.Name {
		case "get_hamburger_items":
			fmt.Println(len(function.Parameters["items"].([]any)))

			break

		case "get_alexa_response":
			message := function.Parameters["response"].(string)
			return &message, nil
		}
	}

	return nil, nil
}

// FetchAllFromLastSync implements domain.ItemUseCase.
func (usecase *itemUseCase) FetchAllFromLastSync(ctx context.Context) ([]domain.Item, error) {
	panic("unimplemented")
}

// Sync implements domain.ItemUseCase.
func (usecase *itemUseCase) Sync(ctx context.Context) error {
	items, err := usecase.itemScraperGateway.ScrapeItems(ctx)
	if err != nil {
		return err
	}

	for _, item := range items {
		if item.Price > 0 {
			_, err := usecase.itemDatabaseRepository.Save(ctx, item)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func NewItemUseCase(
	itemLLMGateway domain.ItemLLMGateway,
	itemScraperGateway domain.ItemScraperGateway,
	itemDatabaseRepository domain.ItemDatabaseRepository,
) domain.ItemUseCase {
	return &itemUseCase{
		itemLLMGateway:         itemLLMGateway,
		itemScraperGateway:     itemScraperGateway,
		itemDatabaseRepository: itemDatabaseRepository,
	}
}
