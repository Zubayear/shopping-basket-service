package model

import "github.com/google/uuid"

type BasketLineMessage struct {
	BasketLineId uuid.UUID
	Price        float32
	TicketAmount uint
}
