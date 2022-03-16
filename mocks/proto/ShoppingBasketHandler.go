// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	ShoppingBasket "ShoppingBasket/proto"
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// ShoppingBasketHandler is an autogenerated mock type for the ShoppingBasketHandler type
type ShoppingBasketHandler struct {
	mock.Mock
}

// CreateBasket provides a mock function with given fields: _a0, _a1, _a2
func (_m *ShoppingBasketHandler) CreateBasket(_a0 context.Context, _a1 *ShoppingBasket.CreateBasketRequest, _a2 *ShoppingBasket.CreateBasketResponse) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *ShoppingBasket.CreateBasketRequest, *ShoppingBasket.CreateBasketResponse) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateBasketLine provides a mock function with given fields: _a0, _a1, _a2
func (_m *ShoppingBasketHandler) CreateBasketLine(_a0 context.Context, _a1 *ShoppingBasket.CreateBasketLineRequest, _a2 *ShoppingBasket.CreateBasketLineResponse) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *ShoppingBasket.CreateBasketLineRequest, *ShoppingBasket.CreateBasketLineResponse) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateEvent provides a mock function with given fields: _a0, _a1, _a2
func (_m *ShoppingBasketHandler) CreateEvent(_a0 context.Context, _a1 *ShoppingBasket.CreateEventRequest, _a2 *ShoppingBasket.CreateEventResponse) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *ShoppingBasket.CreateEventRequest, *ShoppingBasket.CreateEventResponse) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteBasketLineById provides a mock function with given fields: _a0, _a1, _a2
func (_m *ShoppingBasketHandler) DeleteBasketLineById(_a0 context.Context, _a1 *ShoppingBasket.DeleteBasketLineByIdRequest, _a2 *ShoppingBasket.DeleteBasketLineByIdResponse) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *ShoppingBasket.DeleteBasketLineByIdRequest, *ShoppingBasket.DeleteBasketLineByIdResponse) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetBasketById provides a mock function with given fields: _a0, _a1, _a2
func (_m *ShoppingBasketHandler) GetBasketById(_a0 context.Context, _a1 *ShoppingBasket.GetBasketByIdRequest, _a2 *ShoppingBasket.GetBasketByIdResponse) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *ShoppingBasket.GetBasketByIdRequest, *ShoppingBasket.GetBasketByIdResponse) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetBasketLineById provides a mock function with given fields: _a0, _a1, _a2
func (_m *ShoppingBasketHandler) GetBasketLineById(_a0 context.Context, _a1 *ShoppingBasket.GetBasketLineByIdRequest, _a2 *ShoppingBasket.GetBasketLineByIdResponse) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *ShoppingBasket.GetBasketLineByIdRequest, *ShoppingBasket.GetBasketLineByIdResponse) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetBasketLinesByBasketId provides a mock function with given fields: _a0, _a1, _a2
func (_m *ShoppingBasketHandler) GetBasketLinesByBasketId(_a0 context.Context, _a1 *ShoppingBasket.GetBasketLinesByBasketIdRequest, _a2 *ShoppingBasket.GetBasketLinesByBasketIdResponse) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *ShoppingBasket.GetBasketLinesByBasketIdRequest, *ShoppingBasket.GetBasketLinesByBasketIdResponse) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetEventById provides a mock function with given fields: _a0, _a1, _a2
func (_m *ShoppingBasketHandler) GetEventById(_a0 context.Context, _a1 *ShoppingBasket.GetEventByIdRequest, _a2 *ShoppingBasket.GetEventByIdResponse) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *ShoppingBasket.GetEventByIdRequest, *ShoppingBasket.GetEventByIdResponse) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateBasketLine provides a mock function with given fields: _a0, _a1, _a2
func (_m *ShoppingBasketHandler) UpdateBasketLine(_a0 context.Context, _a1 *ShoppingBasket.UpdateBasketLineRequest, _a2 *ShoppingBasket.UpdateBasketLineResponse) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *ShoppingBasket.UpdateBasketLineRequest, *ShoppingBasket.UpdateBasketLineResponse) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
