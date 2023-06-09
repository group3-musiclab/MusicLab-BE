// Code generated by mockery v2.20.2. DO NOT EDIT.

package mocks

import (
	helper "musiclab-be/utils/helper"

	mock "github.com/stretchr/testify/mock"

	oauth2 "golang.org/x/oauth2"
)

// GoogleAPI is an autogenerated mock type for the GoogleAPI type
type GoogleAPI struct {
	mock.Mock
}

// CreateCalendar provides a mock function with given fields: token, detail
func (_m *GoogleAPI) CreateCalendar(token *oauth2.Token, detail helper.CalendarDetail) error {
	ret := _m.Called(token, detail)

	var r0 error
	if rf, ok := ret.Get(0).(func(*oauth2.Token, helper.CalendarDetail) error); ok {
		r0 = rf(token, detail)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetToken provides a mock function with given fields: code
func (_m *GoogleAPI) GetToken(code string) (*oauth2.Token, error) {
	ret := _m.Called(code)

	var r0 *oauth2.Token
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*oauth2.Token, error)); ok {
		return rf(code)
	}
	if rf, ok := ret.Get(0).(func(string) *oauth2.Token); ok {
		r0 = rf(code)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*oauth2.Token)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(code)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUrlAuth provides a mock function with given fields: state
func (_m *GoogleAPI) GetUrlAuth(state string) string {
	ret := _m.Called(state)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(state)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetUserInfo provides a mock function with given fields: token
func (_m *GoogleAPI) GetUserInfo(token *oauth2.Token) (helper.GoogleCore, error) {
	ret := _m.Called(token)

	var r0 helper.GoogleCore
	var r1 error
	if rf, ok := ret.Get(0).(func(*oauth2.Token) (helper.GoogleCore, error)); ok {
		return rf(token)
	}
	if rf, ok := ret.Get(0).(func(*oauth2.Token) helper.GoogleCore); ok {
		r0 = rf(token)
	} else {
		r0 = ret.Get(0).(helper.GoogleCore)
	}

	if rf, ok := ret.Get(1).(func(*oauth2.Token) error); ok {
		r1 = rf(token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewGoogleAPI interface {
	mock.TestingT
	Cleanup(func())
}

// NewGoogleAPI creates a new instance of GoogleAPI. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewGoogleAPI(t mockConstructorTestingTNewGoogleAPI) *GoogleAPI {
	mock := &GoogleAPI{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
