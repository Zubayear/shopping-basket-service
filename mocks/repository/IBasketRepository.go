// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	ent "ShoppingBasket/ent"
	context "context"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// IBasketRepository is an autogenerated mock type for the IBasketRepository type
type IBasketRepository struct {
	mock.Mock
}

// AddBasket provides a mock function with given fields: ctx, basket
func (_m *IBasketRepository) AddBasket(ctx context.Context, basket *ent.Basket) (*ent.Basket, error) {
	ret := _m.Called(ctx, basket)

	var r0 *ent.Basket
	if rf, ok := ret.Get(0).(func(context.Context, *ent.Basket) *ent.Basket); ok {
		r0 = rf(ctx, basket)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ent.Basket)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *ent.Basket) error); ok {
		r1 = rf(ctx, basket)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBasketById provides a mock function with given fields: ctx, basketId
func (_m *IBasketRepository) GetBasketById(ctx context.Context, basketId uuid.UUID) (*ent.Basket, error) {
	ret := _m.Called(ctx, basketId)

	var r0 *ent.Basket
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *ent.Basket); ok {
		r0 = rf(ctx, basketId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ent.Basket)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, basketId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}