// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	mentors "musiclab-be/features/mentors"

	mock "github.com/stretchr/testify/mock"
)

// MentorService is an autogenerated mock type for the MentorService type
type MentorService struct {
	mock.Mock
}

// Delete provides a mock function with given fields: mentorID
func (_m *MentorService) Delete(mentorID uint) error {
	ret := _m.Called(mentorID)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(mentorID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InsertCredential provides a mock function with given fields: input
func (_m *MentorService) InsertCredential(input mentors.CredentialCore) error {
	ret := _m.Called(input)

	var r0 error
	if rf, ok := ret.Get(0).(func(mentors.CredentialCore) error); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SelectAll provides a mock function with given fields: page, limit, filter
func (_m *MentorService) SelectAll(page int, limit int, filter mentors.MentorFilter) ([]mentors.Core, error) {
	ret := _m.Called(page, limit, filter)

	var r0 []mentors.Core
	if rf, ok := ret.Get(0).(func(int, int, mentors.MentorFilter) []mentors.Core); ok {
		r0 = rf(page, limit, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]mentors.Core)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int, mentors.MentorFilter) error); ok {
		r1 = rf(page, limit, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectAllByRating provides a mock function with given fields:
func (_m *MentorService) SelectAllByRating() ([]mentors.Core, error) {
	ret := _m.Called()

	var r0 []mentors.Core
	if rf, ok := ret.Get(0).(func() []mentors.Core); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]mentors.Core)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectProfile provides a mock function with given fields: mentorID
func (_m *MentorService) SelectProfile(mentorID uint) (mentors.Core, error) {
	ret := _m.Called(mentorID)

	var r0 mentors.Core
	if rf, ok := ret.Get(0).(func(uint) mentors.Core); ok {
		r0 = rf(mentorID)
	} else {
		r0 = ret.Get(0).(mentors.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(mentorID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateData provides a mock function with given fields: mentorID, input
func (_m *MentorService) UpdateData(mentorID uint, input mentors.Core) error {
	ret := _m.Called(mentorID, input)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, mentors.Core) error); ok {
		r0 = rf(mentorID, input)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdatePassword provides a mock function with given fields: mentorID, input
func (_m *MentorService) UpdatePassword(mentorID uint, input mentors.Core) error {
	ret := _m.Called(mentorID, input)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, mentors.Core) error); ok {
		r0 = rf(mentorID, input)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewMentorService interface {
	mock.TestingT
	Cleanup(func())
}

// NewMentorService creates a new instance of MentorService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMentorService(t mockConstructorTestingTNewMentorService) *MentorService {
	mock := &MentorService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
