// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	sql "entgo.io/ent/dialect/sql"
)

// Basket is an autogenerated mock type for the Basket type
type Basket struct {
	mock.Mock
}

// Execute provides a mock function with given fields: _a0
func (_m *Basket) Execute(_a0 *sql.Selector) {
	_m.Called(_a0)
}