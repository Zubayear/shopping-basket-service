package repository

import (
	"ShoppingBasket/ent"
	"github.com/google/uuid"
)

type IBasketLineRepository interface {
	GetBasketLinesByBasketId(basketId uuid.UUID) ([]*ent.BasketLine, error)
	GetBasketLineById(basketLineId uuid.UUID) (*ent.BasketLine, error)
	AddBasketLine(basketId uuid.UUID, basketLine *ent.BasketLine) (*ent.BasketLine, error)
	UpdateBasketLine(basketLine *ent.BasketLine) (*ent.BasketLine, error)
	DeleteBasketLine(basketLineId uuid.UUID) (string, error)
}

type BasketLineRepository struct {
	client *ent.Client
}

func BasketLineRepositoryProvider(client *ent.Client) (*BasketLineRepository, error) {
	return &BasketLineRepository{client: client}, nil
}

func (b *BasketLineRepository) GetBasketLinesByBasketId(basketId uuid.UUID) ([]*ent.BasketLine, error) {
	//TODO implement me
	panic("implement me")
}

func (b *BasketLineRepository) GetBasketLineById(basketLineId uuid.UUID) (*ent.BasketLine, error) {
	//TODO implement me
	panic("implement me")
}

func (b *BasketLineRepository) AddBasketLine(basketId uuid.UUID, basketLine *ent.BasketLine) (*ent.BasketLine, error) {
	//TODO implement me
	panic("implement me")
}

func (b *BasketLineRepository) UpdateBasketLine(basketLine *ent.BasketLine) (*ent.BasketLine, error) {
	//TODO implement me
	panic("implement me")
}

func (b *BasketLineRepository) DeleteBasketLine(basketLineId uuid.UUID) (string, error) {
	//TODO implement me
	panic("implement me")
}

func NewBasketLineRepository() IBasketLineRepository {
	return &BasketLineRepository{}
}
