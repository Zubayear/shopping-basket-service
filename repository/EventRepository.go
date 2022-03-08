package repository

import (
	"ShoppingBasket/ent"
	"context"
	"fmt"
	"github.com/google/uuid"
)

type IEventRepository interface {
	AddEvent(ctx context.Context, event *ent.Event) (*ent.Event, error)
	GetEventById(ctx context.Context, eventId uuid.UUID) (*ent.Event, error)
}

type EventRepository struct {
	client *ent.Client
}

func EventRepositoryProvider(client *ent.Client) (*EventRepository, error) {
	return &EventRepository{client: client}, nil
}

func (e *EventRepository) AddEvent(ctx context.Context, event *ent.Event) (*ent.Event, error) {
	savedEvent, err := e.client.Event.Create().SetName(event.Name).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating event: %w", err)
	}
	return savedEvent, nil
}

func (e *EventRepository) GetEventById(ctx context.Context, eventId uuid.UUID) (*ent.Event, error) {
	eventFromRepo, err := e.client.Event.Get(ctx, eventId)
	if err != nil {
		return nil, err
	}
	return eventFromRepo, nil
}

func NewEventRepository() IEventRepository {
	return &EventRepository{}
}
