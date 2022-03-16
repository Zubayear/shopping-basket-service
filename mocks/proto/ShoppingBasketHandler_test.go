package mocks

import (
	ShoppingBasket "ShoppingBasket/proto"
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestShoppingBasketHandler_CreateBasket(t *testing.T) {
	testObj := new(ShoppingBasketHandler)
	ctx := context.Background()
	request := &ShoppingBasket.CreateBasketRequest{
		UserId:     uuid.New().String(),
		CouponCode: nil,
	}
	basket := &ShoppingBasket.Basket{
		BasketId: uuid.New().String(),
		UserId:   uuid.New().String(),
	}
	response := &ShoppingBasket.CreateBasketResponse{
		Basket: basket,
	}
	testObj.On("CreateBasket", ctx, request, response).Return(nil)
	testObj.CreateBasket(ctx, request, response)
	testObj.AssertExpectations(t)
}

func TestShoppingBasketHandler_CreateBasketLine(t *testing.T) {
	testObj := new(ShoppingBasketHandler)
	ctx := context.Background()
	request := &ShoppingBasket.CreateBasketLineRequest{
		BasketId:     uuid.New().String(),
		Price:        78.09,
		TicketAmount: 4,
	}
	basketLine := &ShoppingBasket.BasketLine{
		BasketLineId: uuid.New().String(),
		TicketAmount: 4,
		Price:        67.09,
		BasketId:     uuid.New().String(),
	}
	response := &ShoppingBasket.CreateBasketLineResponse{
		BasketLine: basketLine,
	}
	testObj.On("CreateBasketLine", ctx, request, response).Return(nil)
	testObj.CreateBasketLine(ctx, request, response)
	testObj.AssertExpectations(t)
}

func TestShoppingBasketHandler_CreateEvent(t *testing.T) {
	type fields struct {
		Mock mock.Mock
	}
	type args struct {
		_a0 context.Context
		_a1 *ShoppingBasket.CreateEventRequest
		_a2 *ShoppingBasket.CreateEventResponse
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_m := &ShoppingBasketHandler{
				Mock: tt.fields.Mock,
			}
			if err := _m.CreateEvent(tt.args._a0, tt.args._a1, tt.args._a2); (err != nil) != tt.wantErr {
				t.Errorf("CreateEvent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestShoppingBasketHandler_DeleteBasketLineById(t *testing.T) {
	type fields struct {
		Mock mock.Mock
	}
	type args struct {
		_a0 context.Context
		_a1 *ShoppingBasket.DeleteBasketLineByIdRequest
		_a2 *ShoppingBasket.DeleteBasketLineByIdResponse
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_m := &ShoppingBasketHandler{
				Mock: tt.fields.Mock,
			}
			if err := _m.DeleteBasketLineById(tt.args._a0, tt.args._a1, tt.args._a2); (err != nil) != tt.wantErr {
				t.Errorf("DeleteBasketLineById() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestShoppingBasketHandler_GetBasketById(t *testing.T) {
	type fields struct {
		Mock mock.Mock
	}
	type args struct {
		_a0 context.Context
		_a1 *ShoppingBasket.GetBasketByIdRequest
		_a2 *ShoppingBasket.GetBasketByIdResponse
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_m := &ShoppingBasketHandler{
				Mock: tt.fields.Mock,
			}
			if err := _m.GetBasketById(tt.args._a0, tt.args._a1, tt.args._a2); (err != nil) != tt.wantErr {
				t.Errorf("GetBasketById() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestShoppingBasketHandler_GetBasketLineById(t *testing.T) {
	type fields struct {
		Mock mock.Mock
	}
	type args struct {
		_a0 context.Context
		_a1 *ShoppingBasket.GetBasketLineByIdRequest
		_a2 *ShoppingBasket.GetBasketLineByIdResponse
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_m := &ShoppingBasketHandler{
				Mock: tt.fields.Mock,
			}
			if err := _m.GetBasketLineById(tt.args._a0, tt.args._a1, tt.args._a2); (err != nil) != tt.wantErr {
				t.Errorf("GetBasketLineById() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestShoppingBasketHandler_GetBasketLinesByBasketId(t *testing.T) {
	type fields struct {
		Mock mock.Mock
	}
	type args struct {
		_a0 context.Context
		_a1 *ShoppingBasket.GetBasketLinesByBasketIdRequest
		_a2 *ShoppingBasket.GetBasketLinesByBasketIdResponse
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_m := &ShoppingBasketHandler{
				Mock: tt.fields.Mock,
			}
			if err := _m.GetBasketLinesByBasketId(tt.args._a0, tt.args._a1, tt.args._a2); (err != nil) != tt.wantErr {
				t.Errorf("GetBasketLinesByBasketId() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestShoppingBasketHandler_GetEventById(t *testing.T) {
	type fields struct {
		Mock mock.Mock
	}
	type args struct {
		_a0 context.Context
		_a1 *ShoppingBasket.GetEventByIdRequest
		_a2 *ShoppingBasket.GetEventByIdResponse
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_m := &ShoppingBasketHandler{
				Mock: tt.fields.Mock,
			}
			if err := _m.GetEventById(tt.args._a0, tt.args._a1, tt.args._a2); (err != nil) != tt.wantErr {
				t.Errorf("GetEventById() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestShoppingBasketHandler_UpdateBasketLine(t *testing.T) {
	type fields struct {
		Mock mock.Mock
	}
	type args struct {
		_a0 context.Context
		_a1 *ShoppingBasket.UpdateBasketLineRequest
		_a2 *ShoppingBasket.UpdateBasketLineResponse
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_m := &ShoppingBasketHandler{
				Mock: tt.fields.Mock,
			}
			if err := _m.UpdateBasketLine(tt.args._a0, tt.args._a1, tt.args._a2); (err != nil) != tt.wantErr {
				t.Errorf("UpdateBasketLine() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
