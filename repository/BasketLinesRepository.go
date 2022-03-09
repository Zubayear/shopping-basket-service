package repository

import (
	"ShoppingBasket/ent"
	"context"
	"github.com/google/uuid"
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
	//TODO implement me
	panic("implement me")
}

func (b *BasketLineRepository) GetBasketLineById(ctx context.Context, basketLineId uuid.UUID) (*ent.BasketLine, error) {
	//TODO implement me
	panic("implement me")
}

func (b *BasketLineRepository) AddBasketLine(ctx context.Context, basketId uuid.UUID, basketLine *ent.BasketLine) (*ent.BasketLine, error) {
	//TODO implement me
	panic("implement me")
}

func (b *BasketLineRepository) UpdateBasketLine(ctx context.Context, basketLine *ent.BasketLine) (*ent.BasketLine, error) {
	//TODO implement me
	panic("implement me")
}

func (b *BasketLineRepository) DeleteBasketLine(ctx context.Context, basketLineId uuid.UUID) (string, error) {
	//TODO implement me
	panic("implement me")
}

func BasketLineRepositoryProvider(client *ent.Client) (*BasketLineRepository, error) {
	return &BasketLineRepository{client: client}, nil
}

func NewBasketLineRepository() IBasketLineRepository {
	return &BasketLineRepository{}
}
