// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	sql "entgo.io/ent/dialect/sql"
)

// Event is an autogenerated mock type for the Event type
type Event struct {
	mock.Mock
}

// Execute provides a mock function with given fields: _a0
func (_m *Event) Execute(_a0 *sql.Selector) {
	_m.Called(_a0)
}
