// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"
	mock "github.com/stretchr/testify/mock"
)

// ReviewHandler is an autogenerated mock type for the ReviewHandler type
type ReviewHandler struct {
	mock.Mock
}

// GetMentorReview provides a mock function with given fields:
func (_m *ReviewHandler) GetMentorReview() echo.HandlerFunc {
	ret := _m.Called()

	var r0 echo.HandlerFunc
	if rf, ok := ret.Get(0).(func() echo.HandlerFunc); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(echo.HandlerFunc)
		}
	}

	return r0
}

// PostMentorReview provides a mock function with given fields:
func (_m *ReviewHandler) PostMentorReview() echo.HandlerFunc {
	ret := _m.Called()

	var r0 echo.HandlerFunc
	if rf, ok := ret.Get(0).(func() echo.HandlerFunc); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(echo.HandlerFunc)
		}
	}

	return r0
}

type mockConstructorTestingTNewReviewHandler interface {
	mock.TestingT
	Cleanup(func())
}

// NewReviewHandler creates a new instance of ReviewHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewReviewHandler(t mockConstructorTestingTNewReviewHandler) *ReviewHandler {
	mock := &ReviewHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}