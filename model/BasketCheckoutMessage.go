package model

import "github.com/google/uuid"

type BasketCheckoutMessage struct {
	BasketId       uuid.UUID           `json:"basket_id"`
	FirstName      string              `json:"first_name"`
	LastName       string              `json:"last_name"`
	Email          string              `json:"email"`
	Address        string              `json:"address"`
	ZipCode        string              `json:"zip_code"`
	City           string              `json:"city"`
	Country        string              `json:"country"`
	UserId         uuid.UUID           `json:"user_id"`
	CardNumber     string              `json:"card_number"`
	CardName       string              `json:"card_name"`
	CardExpiration string              `json:"card_expiration"`
	CvvCode        string              `json:"cvv_code"`
	BasketLines    []BasketLineMessage `json:"basket_lines"`
	BasketTotal    float64             `json:"basket_total"`
}
