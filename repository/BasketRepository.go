package repository

import (
	"ShoppingBasket/ent"
	"context"
	"fmt"
	"github.com/google/uuid"
)

type IBasketRepository interface {
	AddBasket(ctx context.Context, basket *ent.Basket) (*ent.Basket, error)
	GetBasketById(ctx context.Context, basketId uuid.UUID) (*ent.Basket, error)
}

type BasketRepository struct {
	// client to save in db
	client *ent.Client
}

func (b *BasketRepository) AddBasket(ctx context.Context, basket *ent.Basket) (*ent.Basket, error) {
	savedBasket, err := b.client.Basket.Create().SetUserId(basket.UserId).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating basket: %w", err)
	}
	return savedBasket, nil
}

func (b *BasketRepository) GetBasketById(ctx context.Context, basketId uuid.UUID) (*ent.Basket, error) {
	basketFromRepo, err := b.client.Basket.Get(ctx, basketId)
	if err != nil {
		return nil, fmt.Errorf("failed getting basket: %w", err)
	}
	return basketFromRepo, nil
}

// BasketRepositoryProvider takes initialized client and provider BasketRepository
func BasketRepositoryProvider(client *ent.Client) (*BasketRepository, error) {
	return &BasketRepository{client: client}, nil
}

func NewBasketRepository() IBasketRepository {
	return &BasketRepository{}
}
