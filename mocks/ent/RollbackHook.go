// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	ent "ShoppingBasket/ent"

	mock "github.com/stretchr/testify/mock"
)

// RollbackHook is an autogenerated mock type for the RollbackHook type
type RollbackHook struct {
	mock.Mock
}

// Execute provides a mock function with given fields: _a0
func (_m *RollbackHook) Execute(_a0 ent.Rollbacker) ent.Rollbacker {
	ret := _m.Called(_a0)

	var r0 ent.Rollbacker
	if rf, ok := ret.Get(0).(func(ent.Rollbacker) ent.Rollbacker); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ent.Rollbacker)
		}
	}

	return r0
}
