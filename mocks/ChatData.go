// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	chats "musiclab-be/features/chats"

	mock "github.com/stretchr/testify/mock"
)

// ChatData is an autogenerated mock type for the ChatData type
type ChatData struct {
	mock.Mock
}

// GetAll provides a mock function with given fields: mentorID, studentID
func (_m *ChatData) GetAll(mentorID uint, studentID uint) ([]chats.Core, error) {
	ret := _m.Called(mentorID, studentID)

	var r0 []chats.Core
	if rf, ok := ret.Get(0).(func(uint, uint) []chats.Core); ok {
		r0 = rf(mentorID, studentID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]chats.Core)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, uint) error); ok {
		r1 = rf(mentorID, studentID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByStudent provides a mock function with given fields: mentorID
func (_m *ChatData) GetByStudent(mentorID uint) ([]chats.Core, error) {
	ret := _m.Called(mentorID)

	var r0 []chats.Core
	if rf, ok := ret.Get(0).(func(uint) []chats.Core); ok {
		r0 = rf(mentorID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]chats.Core)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(mentorID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertChat provides a mock function with given fields: input
func (_m *ChatData) InsertChat(input chats.Core) error {
	ret := _m.Called(input)

	var r0 error
	if rf, ok := ret.Get(0).(func(chats.Core) error); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewChatData interface {
	mock.TestingT
	Cleanup(func())
}

// NewChatData creates a new instance of ChatData. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewChatData(t mockConstructorTestingTNewChatData) *ChatData {
	mock := &ChatData{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}