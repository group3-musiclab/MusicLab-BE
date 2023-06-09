// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	schedules "musiclab-be/features/schedules"

	mock "github.com/stretchr/testify/mock"
)

// ScheduleService is an autogenerated mock type for the ScheduleService type
type ScheduleService struct {
	mock.Mock
}

// CheckSchedule provides a mock function with given fields: input
func (_m *ScheduleService) CheckSchedule(input schedules.Core) error {
	ret := _m.Called(input)

	var r0 error
	if rf, ok := ret.Get(0).(func(schedules.Core) error); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: mentorID, scheduleID
func (_m *ScheduleService) Delete(mentorID uint, scheduleID uint) error {
	ret := _m.Called(mentorID, scheduleID)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, uint) error); ok {
		r0 = rf(mentorID, scheduleID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetMentorSchedule provides a mock function with given fields: mentorID
func (_m *ScheduleService) GetMentorSchedule(mentorID uint) ([]schedules.Core, error) {
	ret := _m.Called(mentorID)

	var r0 []schedules.Core
	if rf, ok := ret.Get(0).(func(uint) []schedules.Core); ok {
		r0 = rf(mentorID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]schedules.Core)
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

// PostSchedule provides a mock function with given fields: newSchedule
func (_m *ScheduleService) PostSchedule(newSchedule schedules.Core) error {
	ret := _m.Called(newSchedule)

	var r0 error
	if rf, ok := ret.Get(0).(func(schedules.Core) error); ok {
		r0 = rf(newSchedule)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewScheduleService interface {
	mock.TestingT
	Cleanup(func())
}

// NewScheduleService creates a new instance of ScheduleService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewScheduleService(t mockConstructorTestingTNewScheduleService) *ScheduleService {
	mock := &ScheduleService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
