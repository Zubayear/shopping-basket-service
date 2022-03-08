package main

import (
	"ShoppingBasket/di"
	pb "ShoppingBasket/proto"
	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"
)

var (
	service = "shoppingbasket"
	version = "latest"
	port    = ":60008"
)

func main() {
	// Create service
	srv := micro.NewService(
		micro.Name(service),
		micro.Version(version),
		micro.Address(port),
	)
	srv.Init()

	// call the fully initialized dependency
	shoppingBasket, err := di.DependencyProvider()
	if err != nil {
		return
	}

	// Register handler
	err = pb.RegisterShoppingBasketHandler(srv.Server(), shoppingBasket)
	if err != nil {
		return
	}
	// Run service

	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
