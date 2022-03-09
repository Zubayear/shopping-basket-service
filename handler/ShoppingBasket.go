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

func (e *ShoppingBasket) CreateBasket(ctx context.Context, request *pb.CreateBasketRequest, response *pb.CreateBasketResponse) error {
	//TODO implement me
	log.Infof("Received ShoppingBasket.CreateBasket request: %v", request)
	userId := uuid.MustParse(request.UserId)
	basketToSave := &ent.Basket{UserId: userId}
	basketFromRepo, err := e.basketRepo.AddBasket(ctx, basketToSave)
	if err != nil {
		return err
	}
	response.Basket = &pb.Basket{
		BasketId: basketFromRepo.ID.String(),
		UserId:   basketFromRepo.UserId.String(),
	}
	return nil
}

func (e *ShoppingBasket) GetBasketById(ctx context.Context, request *pb.GetBasketByIdRequest, response *pb.GetBasketByIdResponse) error {
	log.Infof("Received ShoppingBasket.GetBasketById request: %v", request)
	basketId := uuid.MustParse(request.BasketId)
	basketFromRepo, err := e.basketRepo.GetBasketById(ctx, basketId)
	if err != nil {
		return err
	}
	response.Basket = &pb.Basket{
		BasketId: basketFromRepo.ID.String(),
		UserId:   basketFromRepo.UserId.String(),
	}
	return nil
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

func (e *ShoppingBasket) GetEventById(ctx context.Context, request *pb.GetEventByIdRequest, response *pb.GetEventByIdResponse) error {
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

// NewShoppingBasket IDE support to generate interface stub
func NewShoppingBasket() pb.ShoppingBasketHandler {
	return &ShoppingBasket{}
}
