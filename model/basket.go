package model

import "github.com/google/uuid"

type Basket struct {
	BasketId      uuid.UUID
	UserId        uuid.UUID
	NumberOfItems uint
}

type BasketForCreation struct {
	UserId uuid.UUID
}

type BasketLine struct {
	BasketLineId uuid.UUID
	BasketId     uuid.UUID
	EventId      uuid.UUID
	Price        uint
	TicketAmount uint
	Event
}

type Event struct {
	EventId uuid.UUID
	Name    string
	Date    int64
}
