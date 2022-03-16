package mocks

import (
	ShoppingBasket "ShoppingBasket/proto"
	"context"
	"github.com/stretchr/testify/mock"
	"go-micro.dev/v4/client"
	"reflect"
	"testing"
)

func TestShoppingBasketService_CreateBasket(t *testing.T) {
	type fields struct {
		Mock mock.Mock
	}
	type args struct {
		ctx  context.Context
		in   *ShoppingBasket.CreateBasketRequest
		opts []client.CallOption
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *ShoppingBasket.CreateBasketResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_m := &ShoppingBasketService{
				Mock: tt.fields.Mock,
			}
			got, err := _m.CreateBasket(tt.args.ctx, tt.args.in, tt.args.opts...)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateBasket() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateBasket() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShoppingBasketService_CreateBasketLine(t *testing.T) {
	type fields struct {
		Mock mock.Mock
	}
	type args struct {
		ctx  context.Context
		in   *ShoppingBasket.CreateBasketLineRequest
		opts []client.CallOption
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *ShoppingBasket.CreateBasketLineResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_m := &ShoppingBasketService{
				Mock: tt.fields.Mock,
			}
			got, err := _m.CreateBasketLine(tt.args.ctx, tt.args.in, tt.args.opts...)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateBasketLine() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateBasketLine() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShoppingBasketService_CreateEvent(t *testing.T) {
	type fields struct {
		Mock mock.Mock
	}
	type args struct {
		ctx  context.Context
		in   *ShoppingBasket.CreateEventRequest
		opts []client.CallOption
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *ShoppingBasket.CreateEventResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_m := &ShoppingBasketService{
				Mock: tt.fields.Mock,
			}
			got, err := _m.CreateEvent(tt.args.ctx, tt.args.in, tt.args.opts...)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateEvent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateEvent() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShoppingBasketService_DeleteBasketLineById(t *testing.T) {
	type fields struct {
		Mock mock.Mock
	}
	type args struct {
		ctx  context.Context
		in   *ShoppingBasket.DeleteBasketLineByIdRequest
		opts []client.CallOption
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *ShoppingBasket.DeleteBasketLineByIdResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_m := &ShoppingBasketService{
				Mock: tt.fields.Mock,
			}
			got, err := _m.DeleteBasketLineById(tt.args.ctx, tt.args.in, tt.args.opts...)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteBasketLineById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteBasketLineById() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShoppingBasketService_GetBasketById(t *testing.T) {
	type fields struct {
		Mock mock.Mock
	}
	type args struct {
		ctx  context.Context
		in   *ShoppingBasket.GetBasketByIdRequest
		opts []client.CallOption
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *ShoppingBasket.GetBasketByIdResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_m := &ShoppingBasketService{
				Mock: tt.fields.Mock,
			}
			got, err := _m.GetBasketById(tt.args.ctx, tt.args.in, tt.args.opts...)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBasketById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBasketById() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShoppingBasketService_GetBasketLineById(t *testing.T) {
	type fields struct {
		Mock mock.Mock
	}
	type args struct {
		ctx  context.Context
		in   *ShoppingBasket.GetBasketLineByIdRequest
		opts []client.CallOption
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *ShoppingBasket.GetBasketLineByIdResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_m := &ShoppingBasketService{
				Mock: tt.fields.Mock,
			}
			got, err := _m.GetBasketLineById(tt.args.ctx, tt.args.in, tt.args.opts...)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBasketLineById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBasketLineById() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShoppingBasketService_GetBasketLinesByBasketId(t *testing.T) {
	type fields struct {
		Mock mock.Mock
	}
	type args struct {
		ctx  context.Context
		in   *ShoppingBasket.GetBasketLinesByBasketIdRequest
		opts []client.CallOption
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *ShoppingBasket.GetBasketLinesByBasketIdResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_m := &ShoppingBasketService{
				Mock: tt.fields.Mock,
			}
			got, err := _m.GetBasketLinesByBasketId(tt.args.ctx, tt.args.in, tt.args.opts...)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBasketLinesByBasketId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBasketLinesByBasketId() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShoppingBasketService_GetEventById(t *testing.T) {
	type fields struct {
		Mock mock.Mock
	}
	type args struct {
		ctx  context.Context
		in   *ShoppingBasket.GetEventByIdRequest
		opts []client.CallOption
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *ShoppingBasket.GetEventByIdResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_m := &ShoppingBasketService{
				Mock: tt.fields.Mock,
			}
			got, err := _m.GetEventById(tt.args.ctx, tt.args.in, tt.args.opts...)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetEventById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetEventById() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShoppingBasketService_UpdateBasketLine(t *testing.T) {
	type fields struct {
		Mock mock.Mock
	}
	type args struct {
		ctx  context.Context
		in   *ShoppingBasket.UpdateBasketLineRequest
		opts []client.CallOption
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *ShoppingBasket.UpdateBasketLineResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_m := &ShoppingBasketService{
				Mock: tt.fields.Mock,
			}
			got, err := _m.UpdateBasketLine(tt.args.ctx, tt.args.in, tt.args.opts...)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateBasketLine() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateBasketLine() got = %v, want %v", got, tt.want)
			}
		})
	}
}
