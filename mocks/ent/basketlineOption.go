// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	ent "ShoppingBasket/ent"

	mock "github.com/stretchr/testify/mock"
)

// basketlineOption is an autogenerated mock type for the basketlineOption type
type basketlineOption struct {
	mock.Mock
}

// Execute provides a mock function with given fields: _a0
func (_m *basketlineOption) Execute(_a0 *ent.BasketLineMutation) {
	_m.Called(_a0)
}
