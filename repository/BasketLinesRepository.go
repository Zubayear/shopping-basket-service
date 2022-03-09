package repository

import (
	"ShoppingBasket/ent"
	"ShoppingBasket/ent/basket"
	"context"
	"fmt"
	"github.com/google/uuid"
	log "go-micro.dev/v4/logger"
)

type IBasketLineRepository interface {
	GetBasketLinesByBasketId(ctx context.Context, basketId uuid.UUID) ([]*ent.BasketLine, error)
	GetBasketLineById(ctx context.Context, basketLineId uuid.UUID) (*ent.BasketLine, error)
	AddBasketLine(ctx context.Context, basketId uuid.UUID, basketLine *ent.BasketLine) (*ent.BasketLine, error)
	UpdateBasketLine(ctx context.Context, basketLine *ent.BasketLine) (*ent.BasketLine, error)
	DeleteBasketLine(ctx context.Context, basketLineId uuid.UUID) (string, error)
}

type BasketLineRepository struct {
	client *ent.Client
}

func (b *BasketLineRepository) GetBasketLinesByBasketId(ctx context.Context, basketId uuid.UUID) ([]*ent.BasketLine, error) {
	basketLinesByBasketId, err := b.client.Basket.Query().Where(basket.ID(basketId)).QueryBasketLine().All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed getting basketlines with basketId: %w", err)
	}
	return basketLinesByBasketId, nil
}

func (b *BasketLineRepository) GetBasketLineById(ctx context.Context, basketLineId uuid.UUID) (*ent.BasketLine, error) {
	log.Infof("Received BasketLineRepository.GetBasketLineById request: %v", basketLineId)
	basketLine, err := b.client.BasketLine.Get(ctx, basketLineId)
	if err != nil {
		return nil, fmt.Errorf("failed getting basketline: %w", err)
	}
	log.Infof("basket line from repo %v", basketLine)
	return basketLine, nil
}

func (b *BasketLineRepository) AddBasketLine(ctx context.Context, basketId uuid.UUID, basketLine *ent.BasketLine) (*ent.BasketLine, error) {
	savedBasketLine, err := b.client.BasketLine.Create().
		SetBasketID(basketId).
		SetPrice(basketLine.Price).
		SetTicketAmount(basketLine.TicketAmount).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating basketline: %w", err)
	}
	return savedBasketLine, nil
}

func (b *BasketLineRepository) UpdateBasketLine(ctx context.Context, basketLine *ent.BasketLine) (*ent.BasketLine, error) {
	updatedBasketLine, err := b.client.BasketLine.UpdateOneID(basketLine.ID).SetPrice(basketLine.Price).SetTicketAmount(basketLine.TicketAmount).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed updating basketLine: %w", err)
	}
	return updatedBasketLine, nil
}

func (b *BasketLineRepository) DeleteBasketLine(ctx context.Context, basketLineId uuid.UUID) (string, error) {
	err := b.client.BasketLine.DeleteOneID(basketLineId).Exec(ctx)
	if err != nil {
		return "", fmt.Errorf("failed deleting basketLine with basketLineId: %w", err)
	}
	return fmt.Sprintf("basket line with id '%v' deleted", basketLineId), nil
}

func BasketLineRepositoryProvider(client *ent.Client) (*BasketLineRepository, error) {
	return &BasketLineRepository{client: client}, nil
}

func NewBasketLineRepository() IBasketLineRepository {
	return &BasketLineRepository{}
}
