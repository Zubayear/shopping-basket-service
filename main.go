package main

import (
	"ShoppingBasket/di"
	pb "ShoppingBasket/proto"
	"go-micro.dev/v4"
	"go-micro.dev/v4/broker"
	"go-micro.dev/v4/cmd"
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

	cmd.Init()

	if err := broker.Init(); err != nil {
		log.Fatalf("Broker Init error: %v", err)
	}
	if err := broker.Connect(); err != nil {
		log.Fatalf("Broker Connect error: %v", err)
	}

	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
