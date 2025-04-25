package usecase

import (
	"context"

	"github.com/booscaaa/hamburguer-go/internal/core/domain"
	"github.com/booscaaa/hamburguer-go/internal/core/dto"
)

type reviewUseCase struct {
	reviewLLMGateway domain.ReviewLLMGateway
	reviewRepository domain.ReviewDatabaseRepository
}

// Count implements domain.ReviewUseCase.
func (usecase *reviewUseCase) Count(ctx context.Context) (int, error) {
	return usecase.reviewRepository.Count(ctx)
}

// Fetch implements domain.ReviewUseCase.
func (usecase *reviewUseCase) GetTop3Reviews(ctx context.Context) (*string, error) {
	reviews, err := usecase.reviewRepository.Fetch(ctx)
	if err != nil {
		return nil, err
	}

	response, err := usecase.reviewLLMGateway.GetTop3Reviews(ctx, reviews)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Save implements domain.ReviewUseCase.
func (usecase *reviewUseCase) Save(ctx context.Context, input dto.Review) error {
	return usecase.reviewRepository.Save(ctx, input)
}

func NewReviewUseCase(
	reviewLLMGateway domain.ReviewLLMGateway,
	reviewRepository domain.ReviewDatabaseRepository,
) domain.ReviewUseCase {
	return &reviewUseCase{
		reviewLLMGateway: reviewLLMGateway,
		reviewRepository: reviewRepository,
	}
}
