// Code generated by mockery v2.43.2. DO NOT EDIT.

package mock

import (
	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// Interface is an autogenerated mock type for the Interface type
type Interface struct {
	mock.Mock
}

// New provides a mock function with given fields:
func (_m *Interface) New() uuid.UUID {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for New")
	}

	var r0 uuid.UUID
	if rf, ok := ret.Get(0).(func() uuid.UUID); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(uuid.UUID)
		}
	}

	return r0
}

// NewString provides a mock function with given fields:
func (_m *Interface) NewString() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for NewString")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// NewInterface creates a new instance of Interface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *Interface {
	mock := &Interface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}