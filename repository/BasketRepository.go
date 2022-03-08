package repository

import "ShoppingBasket/ent"

type IBasketRepository interface {
	AddBasket(basket *ent.Basket) (*ent.Basket, error)
	GetBasketById(basketId *ent.Basket) (*ent.Basket, error)
}

type BasketRepository struct {
	// client to save in db
	client *ent.Client
}

// BasketRepositoryProvider takes initialized client and provider BasketRepository
func BasketRepositoryProvider(client *ent.Client) (*BasketRepository, error) {
	return &BasketRepository{client: client}, nil
}

func (b *BasketRepository) AddBasket(basket *ent.Basket) (*ent.Basket, error) {
	//TODO implement me
	panic("implement me")
}

func (b *BasketRepository) GetBasketById(basketId *ent.Basket) (*ent.Basket, error) {
	//TODO implement me
	panic("implement me")
}

func NewBasketRepository() IBasketRepository {
	return &BasketRepository{}
}
