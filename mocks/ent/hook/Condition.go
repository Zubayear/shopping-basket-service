// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	context "context"

	ent "entgo.io/ent"

	mock "github.com/stretchr/testify/mock"
)

// Condition is an autogenerated mock type for the Condition type
type Condition struct {
	mock.Mock
}

// Execute provides a mock function with given fields: _a0, _a1
func (_m *Condition) Execute(_a0 context.Context, _a1 ent.Mutation) bool {
	ret := _m.Called(_a0, _a1)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, ent.Mutation) bool); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}
