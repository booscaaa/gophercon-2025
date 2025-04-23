package dto

type Review struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required,max=200"`
}
