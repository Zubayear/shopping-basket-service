//go:build wireinject
// +build wireinject

package di

import (
	"ShoppingBasket/external"
	"ShoppingBasket/handler"
	"ShoppingBasket/repository"
	"github.com/google/wire"
)

func DependencyProvider() (*handler.ShoppingBasket, error) {
	wire.Build(
		handler.NewShoppingBasketHandler,
		external.HostProvider,
		repository.ClientProvider,
		repository.EventRepositoryProvider,
		repository.BasketRepositoryProvider,
		repository.BasketLineRepositoryProvider,
		wire.Bind(new(repository.IEventRepository), new(*repository.EventRepository)),
		wire.Bind(new(repository.IBasketRepository), new(*repository.BasketRepository)),
		wire.Bind(new(repository.IBasketLineRepository), new(*repository.BasketLineRepository)),
	)
	return &handler.ShoppingBasket{}, nil
}
