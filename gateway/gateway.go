package main

import (
	ShoppingBasket "ShoppingBasket/proto"
	"context"
	"flag"
	"fmt"
	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"go-micro.dev/v4/client"
	"google.golang.org/grpc"
	"net"
	"net/http"
)

var (
	gateway = flag.Int("gateway", 8083, "gateway")
	port    = flag.Int("port", 60008, "shoppingbasket address")
	gRPC    = flag.Int("gRPC", 60008, "gRPC port")
)

var shoppingBasketClient ShoppingBasket.ShoppingBasketService

func init() {
	shoppingBasketClient = ShoppingBasket.NewShoppingBasketService("shoppingbasket", client.DefaultClient)
}

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	s := grpc.NewServer()

	ShoppingBasket.RegisterShoppingBasketServer(s, &ShoppingBasketProxy{})

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))

	if err != nil {
		return err
	}

	go func() {
		err := s.Serve(listen)
		if err != nil {
			return
		}
	}()

	mux := runtime.NewServeMux()
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", *gRPC), grpc.WithInsecure())
	if err != nil {
		return err
	}

	err = ShoppingBasket.RegisterShoppingBasketSB(ctx, mux, conn)
	if err != nil {
		return err
	}

	gwServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", *gateway),
		Handler: mux,
	}

	return gwServer.ListenAndServe()
}

func main() {
	flag.Parse()

	defer glog.Flush()
	err := run()
	if err != nil {
		return
	}
}

type ShoppingBasketProxy struct {
}

func (s *ShoppingBasketProxy) Checkout(ctx context.Context, request *ShoppingBasket.BasketCheckoutRequest) (*ShoppingBasket.BasketCheckoutResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *ShoppingBasketProxy) CreateEvent(ctx context.Context, request *ShoppingBasket.CreateEventRequest) (*ShoppingBasket.CreateEventResponse, error) {
	response, err := shoppingBasketClient.CreateEvent(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (s *ShoppingBasketProxy) GetEventById(ctx context.Context, request *ShoppingBasket.GetEventByIdRequest) (*ShoppingBasket.GetEventByIdResponse, error) {
	response, err := shoppingBasketClient.GetEventById(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (s *ShoppingBasketProxy) CreateBasket(ctx context.Context, request *ShoppingBasket.CreateBasketRequest) (*ShoppingBasket.CreateBasketResponse, error) {
	response, err := shoppingBasketClient.CreateBasket(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (s *ShoppingBasketProxy) GetBasketById(ctx context.Context, request *ShoppingBasket.GetBasketByIdRequest) (*ShoppingBasket.GetBasketByIdResponse, error) {
	response, err := shoppingBasketClient.GetBasketById(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (s *ShoppingBasketProxy) CreateBasketLine(ctx context.Context, request *ShoppingBasket.CreateBasketLineRequest) (*ShoppingBasket.CreateBasketLineResponse, error) {
	response, err := shoppingBasketClient.CreateBasketLine(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (s *ShoppingBasketProxy) GetBasketLineById(ctx context.Context, request *ShoppingBasket.GetBasketLineByIdRequest) (*ShoppingBasket.GetBasketLineByIdResponse, error) {
	response, err := shoppingBasketClient.GetBasketLineById(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (s *ShoppingBasketProxy) UpdateBasketLine(ctx context.Context, request *ShoppingBasket.UpdateBasketLineRequest) (*ShoppingBasket.UpdateBasketLineResponse, error) {
	response, err := shoppingBasketClient.UpdateBasketLine(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (s *ShoppingBasketProxy) DeleteBasketLineById(ctx context.Context, request *ShoppingBasket.DeleteBasketLineByIdRequest) (*ShoppingBasket.DeleteBasketLineByIdResponse, error) {
	response, err := shoppingBasketClient.DeleteBasketLineById(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (s *ShoppingBasketProxy) GetBasketLinesByBasketId(ctx context.Context, request *ShoppingBasket.GetBasketLinesByBasketIdRequest) (*ShoppingBasket.GetBasketLinesByBasketIdResponse, error) {
	response, err := shoppingBasketClient.GetBasketLinesByBasketId(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}
