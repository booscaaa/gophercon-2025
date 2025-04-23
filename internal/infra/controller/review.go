package controller

import (
	"encoding/json"
	"net/http"

	"github.com/booscaaa/hamburguer-go/internal/core/domain"
	"github.com/booscaaa/hamburguer-go/internal/core/dto"
)

type reviewController struct {
	reviewUseCase domain.ReviewUseCase
}

// GetTop3Reviews implements domain.ReviewController.
func (controller *reviewController) GetTop3Reviews(response http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	message, err := controller.reviewUseCase.GetTop3Reviews(ctx)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	var output dto.AlexaResponse
	output.Version = "1.0"
	output.Response.OutputSpeech.Type = "PlainText"
	output.Response.ShouldEndSession = true
	output.Response.OutputSpeech.Text = *message

	response.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(response).Encode(output)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Save implements domain.ReviewController.
func (controller *reviewController) Save(response http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	var input dto.Review
	err := json.NewDecoder(request.Body).Decode(&input)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}
	err = controller.reviewUseCase.Save(ctx, input)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusCreated)
}

func NewReviewController(reviewUseCase domain.ReviewUseCase) domain.ReviewController {
	return &reviewController{
		reviewUseCase: reviewUseCase,
	}
}
