package handler

import (
	"ShoppingBasket/ent"
	"ShoppingBasket/repository"
	"context"
	"github.com/google/uuid"
	log "go-micro.dev/v4/logger"

	pb "ShoppingBasket/proto"
)

type ShoppingBasket struct {
	// need this repo impl for saving to db from service
	basketRepo     repository.IBasketRepository
	basketLineRepo repository.IBasketLineRepository
	eventRepo      repository.IEventRepository
}

// NewShoppingBasket IDE support to generate interface stub
func NewShoppingBasket() pb.ShoppingBasketHandler {
	return &ShoppingBasket{}
}

// NewShoppingBasketHandler takes in interfaces and provider ShoppingBasket
func NewShoppingBasketHandler(
	basketRepo repository.IBasketRepository,
	basketLineRepo repository.IBasketLineRepository,
	eventRepo repository.IEventRepository) (*ShoppingBasket, error) {
	return &ShoppingBasket{
		basketRepo:     basketRepo,
		basketLineRepo: basketLineRepo,
		eventRepo:      eventRepo,
	}, nil
}

func (e *ShoppingBasket) CreateEvent(ctx context.Context, request *pb.CreateEventRequest, response *pb.CreateEventResponse) error {
	log.Infof("Received ShoppingBasket.CreateEvent request: %v", request)
	eventToSave := &ent.Event{Name: request.Name}
	eventFromRepo, err := e.eventRepo.AddEvent(ctx, eventToSave)
	if err != nil {
		return err
	}
	response.Event = &pb.Event{
		Id:   eventFromRepo.ID.String(),
		Name: eventFromRepo.Name,
		Date: eventFromRepo.Date.Unix(),
	}
	return nil
}

func (e *ShoppingBasket) GetEvent(ctx context.Context, request *pb.GetEventByIdRequest, response *pb.GetEventByIdResponse) error {
	log.Infof("Received ShoppingBasket.GetEvent request: %v", request)
	id := uuid.MustParse(request.Id)
	eventFromRepo, err := e.eventRepo.GetEventById(ctx, id)
	if err != nil {
		return err
	}
	response.Event = &pb.Event{
		Id:   eventFromRepo.ID.String(),
		Name: eventFromRepo.Name,
		Date: eventFromRepo.Date.Unix(),
	}
	return nil
}
