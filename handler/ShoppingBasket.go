package handler

import (
	Coupon "ShoppingBasket/coupon-client"
	"ShoppingBasket/ent"
	"ShoppingBasket/repository"
	"context"
	"fmt"
	"github.com/google/uuid"
	"go-micro.dev/v4/client"
	log "go-micro.dev/v4/logger"

	pb "ShoppingBasket/proto"
)

type ShoppingBasket struct {
	// need this repo impl for saving to db from service
	basketRepo     repository.IBasketRepository
	basketLineRepo repository.IBasketLineRepository
	eventRepo      repository.IEventRepository
}

func (e *ShoppingBasket) CreateBasketLine(ctx context.Context, request *pb.CreateBasketLineRequest, response *pb.CreateBasketLineResponse) error {
	log.Infof("Received ShoppingBasket.CreateBasketLine request: %v", request)
	basketLineToSave := &ent.BasketLine{
		TicketAmount: uint(request.TicketAmount),
		Price:        request.Price,
	}
	basketId, err := uuid.Parse(request.BasketId)
	if err != nil {
		return fmt.Errorf("failed parsing id: %w", err)
	}
	basketLine, err := e.basketLineRepo.AddBasketLine(ctx, basketId, basketLineToSave)
	if err != nil {
		return err
	}
	response.BasketLine = &pb.BasketLine{
		BasketLineId: basketLine.ID.String(),
		TicketAmount: uint32(basketLine.TicketAmount),
		Price:        basketLine.Price,
		BasketId:     basketId.String(),
	}
	return nil
}

func (e *ShoppingBasket) GetBasketLineById(ctx context.Context, request *pb.GetBasketLineByIdRequest, response *pb.GetBasketLineByIdResponse) error {
	log.Infof("Received ShoppingBasket.GetBasketLineById request: %v", request)
	basketLineId, err := uuid.Parse(request.BasketLineId)
	if err != nil {
		return fmt.Errorf("failed parsing id: %w", err)
	}
	basketLine, err := e.basketLineRepo.GetBasketLineById(ctx, basketLineId)
	if err != nil {
		return err
	}
	response.BasketLine = &pb.BasketLine{
		BasketLineId: basketLine.ID.String(),
		TicketAmount: uint32(basketLine.TicketAmount),
		Price:        basketLine.Price,
	}
	return nil
}

func (e *ShoppingBasket) UpdateBasketLine(ctx context.Context, request *pb.UpdateBasketLineRequest, response *pb.UpdateBasketLineResponse) error {
	log.Infof("Received ShoppingBasket.UpdateBasketLine request: %v", request)
	basketLineId, err := uuid.Parse(request.BasketLineId)
	if err != nil {
		return fmt.Errorf("failed parsing id: %w", err)
	}
	basketLineToUpdate := &ent.BasketLine{
		ID:           basketLineId,
		TicketAmount: uint(request.TicketAmount),
		Price:        request.Price,
	}
	basketLine, err := e.basketLineRepo.UpdateBasketLine(ctx, basketLineToUpdate)
	if err != nil {
		return err
	}
	response.BasketLine = &pb.BasketLine{
		BasketLineId: basketLine.ID.String(),
		TicketAmount: uint32(basketLine.TicketAmount),
		Price:        basketLine.Price,
	}
	return nil
}

func (e *ShoppingBasket) DeleteBasketLineById(ctx context.Context, request *pb.DeleteBasketLineByIdRequest, response *pb.DeleteBasketLineByIdResponse) error {
	log.Infof("Received ShoppingBasket.DeleteBasketLineById request: %v", request)
	basketLineId, err := uuid.Parse(request.BasketLineId)
	if err != nil {
		return fmt.Errorf("failed parsing id: %w", err)
	}
	msg, err := e.basketLineRepo.DeleteBasketLine(ctx, basketLineId)
	if err != nil {
		return err
	}
	response.Msg = msg
	response.Code = 60000 // success
	return nil
}

func (e *ShoppingBasket) GetBasketLinesByBasketId(ctx context.Context, request *pb.GetBasketLinesByBasketIdRequest, response *pb.GetBasketLinesByBasketIdResponse) error {
	log.Infof("Received ShoppingBasket.GetBasketLinesByBasketId request: %v", request)
	basketId, err := uuid.Parse(request.BasketId)
	if err != nil {
		return fmt.Errorf("failed parsing id: %w", basketId)
	}
	basketLines, err := e.basketLineRepo.GetBasketLinesByBasketId(ctx, basketId)
	if err != nil {
		return err
	}
	var basketLinesToReturn []*pb.BasketLine
	for _, basketLine := range basketLines {
		mappedBasketLine := &pb.BasketLine{}
		mapBasketLine(basketLine, mappedBasketLine)
		basketLinesToReturn = append(basketLinesToReturn, mappedBasketLine)
	}
	response.BasketLine = basketLinesToReturn
	return nil
}

func (e *ShoppingBasket) CreateBasket(ctx context.Context, request *pb.CreateBasketRequest, response *pb.CreateBasketResponse) error {
	log.Infof("Received ShoppingBasket.CreateBasket request: %v", request)
	userId := uuid.MustParse(request.UserId)
	basketToSave := &ent.Basket{UserId: userId, CouponCode: request.CouponCode}
	basketFromRepo, err := e.basketRepo.AddBasket(ctx, basketToSave)
	if err != nil {
		return err
	}
	couponClient := Coupon.NewCouponService("coupon", client.DefaultClient)
	couponStatus, err := couponClient.UseCoupon(ctx, &Coupon.UseCouponRequest{
		CouponCode: request.CouponCode,
	})
	if err != nil {
		return err
	}
	if couponStatus.Used != 1 {
		return fmt.Errorf("coupon is already used or expired")
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

func mapBasketLine(src *ent.BasketLine, dst *pb.BasketLine) {
	dst.BasketLineId = src.ID.String()
	dst.Price = src.Price
	dst.TicketAmount = uint32(src.TicketAmount)
}
