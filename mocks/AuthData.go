// Code generated by mockery v2.20.2. DO NOT EDIT.

package mocks

import (
	auth "musiclab-be/features/auth"

	mock "github.com/stretchr/testify/mock"
)

// AuthData is an autogenerated mock type for the AuthData type
type AuthData struct {
	mock.Mock
}

// FindAccount provides a mock function with given fields: email
func (_m *AuthData) FindAccount(email string) (auth.Core, error) {
	ret := _m.Called(email)

	var r0 auth.Core
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (auth.Core, error)); ok {
		return rf(email)
	}
	if rf, ok := ret.Get(0).(func(string) auth.Core); ok {
		r0 = rf(email)
	} else {
		r0 = ret.Get(0).(auth.Core)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LoginMentor provides a mock function with given fields: email
func (_m *AuthData) LoginMentor(email string) (auth.Core, error) {
	ret := _m.Called(email)

	var r0 auth.Core
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (auth.Core, error)); ok {
		return rf(email)
	}
	if rf, ok := ret.Get(0).(func(string) auth.Core); ok {
		r0 = rf(email)
	} else {
		r0 = ret.Get(0).(auth.Core)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LoginStudent provides a mock function with given fields: email
func (_m *AuthData) LoginStudent(email string) (auth.Core, error) {
	ret := _m.Called(email)

	var r0 auth.Core
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (auth.Core, error)); ok {
		return rf(email)
	}
	if rf, ok := ret.Get(0).(func(string) auth.Core); ok {
		r0 = rf(email)
	} else {
		r0 = ret.Get(0).(auth.Core)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterMentor provides a mock function with given fields: newUser
func (_m *AuthData) RegisterMentor(newUser auth.Core) error {
	ret := _m.Called(newUser)

	var r0 error
	if rf, ok := ret.Get(0).(func(auth.Core) error); ok {
		r0 = rf(newUser)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RegisterStudent provides a mock function with given fields: newUser
func (_m *AuthData) RegisterStudent(newUser auth.Core) error {
	ret := _m.Called(newUser)

	var r0 error
	if rf, ok := ret.Get(0).(func(auth.Core) error); ok {
		r0 = rf(newUser)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewAuthData interface {
	mock.TestingT
	Cleanup(func())
}

// NewAuthData creates a new instance of AuthData. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAuthData(t mockConstructorTestingTNewAuthData) *AuthData {
	mock := &AuthData{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}