// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	genres "musiclab-be/features/genres"

	mock "github.com/stretchr/testify/mock"
)

// GenreService is an autogenerated mock type for the GenreService type
type GenreService struct {
	mock.Mock
}

// AddMentorGenre provides a mock function with given fields: newGenre
func (_m *GenreService) AddMentorGenre(newGenre genres.Core) error {
	ret := _m.Called(newGenre)

	var r0 error
	if rf, ok := ret.Get(0).(func(genres.Core) error); ok {
		r0 = rf(newGenre)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: mentorID, genreID
func (_m *GenreService) Delete(mentorID uint, genreID uint) error {
	ret := _m.Called(mentorID, genreID)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, uint) error); ok {
		r0 = rf(mentorID, genreID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetGenre provides a mock function with given fields:
func (_m *GenreService) GetGenre() ([]genres.Core, error) {
	ret := _m.Called()

	var r0 []genres.Core
	if rf, ok := ret.Get(0).(func() []genres.Core); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]genres.Core)
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

// GetMentorGenre provides a mock function with given fields: mentorID
func (_m *GenreService) GetMentorGenre(mentorID uint) ([]genres.MentorGenreCore, error) {
	ret := _m.Called(mentorID)

	var r0 []genres.MentorGenreCore
	if rf, ok := ret.Get(0).(func(uint) []genres.MentorGenreCore); ok {
		r0 = rf(mentorID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]genres.MentorGenreCore)
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

type mockConstructorTestingTNewGenreService interface {
	mock.TestingT
	Cleanup(func())
}

// NewGenreService creates a new instance of GenreService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewGenreService(t mockConstructorTestingTNewGenreService) *GenreService {
	mock := &GenreService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
