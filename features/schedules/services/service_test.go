package services

import (
	"errors"
	"musiclab-be/features/classes"
	"musiclab-be/features/schedules"
	"musiclab-be/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	mock_class = classes.Core{
		ID:          1,
		MentorID:    1,
		Name:        "How to Play Polyrithm Like Daney Carey",
		Image:       "image-url",
		Level:       "Intermediate",
		Description: "Play Like Daney Carey",
		Syllabus:    "Polyrithmic",
		Requirement: "Have a drum",
		ForWhom:     "For Drummer",
		Price:       100000,
		Duration:    1,
	}
)

func TestCheckSchedule(t *testing.T) {
	repo := mocks.NewScheduleData(t)
	repoClass := mocks.NewClassData(t)

	input := schedules.Core{
		ClassID: 1,
		Transaction: schedules.Transaction{
			ScheduleID: 1,
			StartDate:  time.Date(2023, 4, 2, 0, 0, 0, 0, time.UTC),
		},
	}

	t.Run("Failed Validate", func(t *testing.T) {
		input := schedules.Core{
			ClassID: 1,
		}

		srv := New(repo, repoClass)
		err := srv.CheckSchedule(input)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Failed Date Validation", func(t *testing.T) {
		input := schedules.Core{
			ClassID: 1,
			Transaction: schedules.Transaction{
				ScheduleID: 1,
				StartDate:  time.Date(2023, 3, 2, 0, 0, 0, 0, time.UTC),
			},
		}

		srv := New(repo, repoClass)
		err := srv.CheckSchedule(input)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Schedule Available", func(t *testing.T) {
		repoClass.On("GetMentorClassDetail", mock.Anything).Return(mock_class, nil).Once()
		repo.On("CheckSchedule", mock.Anything).Return(int64(0), nil).Once()

		srv := New(repo, repoClass)
		err := srv.CheckSchedule(input)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Schedule Not Available", func(t *testing.T) {
		repoClass.On("GetMentorClassDetail", mock.Anything).Return(mock_class, nil).Once()
		repo.On("CheckSchedule", mock.Anything).Return(int64(1), nil).Once()

		srv := New(repo, repoClass)
		err := srv.CheckSchedule(input)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Error Get Class", func(t *testing.T) {
		repoClass.On("GetMentorClassDetail", mock.Anything).Return(classes.Core{}, errors.New("error")).Once()

		srv := New(repo, repoClass)
		err := srv.CheckSchedule(input)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Error Check Available", func(t *testing.T) {
		repoClass.On("GetMentorClassDetail", mock.Anything).Return(mock_class, nil).Once()
		repo.On("CheckSchedule", mock.Anything).Return(int64(0), errors.New("error")).Once()

		srv := New(repo, repoClass)
		err := srv.CheckSchedule(input)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})

}

// s
func TestPostSchedule(t *testing.T) {
	repo := mocks.NewScheduleData(t)
	repoClass := mocks.NewClassData(t)

	inputData := schedules.Core{ID: 1, Day: "Monday", StartTime: "13:00", EndTime: "14:00"}

	t.Run("success post Schedule", func(t *testing.T) {
		repo.On("PostSchedule", inputData).Return(nil).Once()
		srv := New(repo, repoClass)
		err := srv.PostSchedule(inputData)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("data not found", func(t *testing.T) {
		repo.On("PostSchedule", inputData).Return(errors.New("data not found")).Once()
		srv := New(repo, repoClass)
		err := srv.PostSchedule(inputData)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "data not found")
		repo.AssertExpectations(t)
	})

	t.Run("server problem", func(t *testing.T) {
		repo.On("PostSchedule", inputData).Return(errors.New("server problem")).Once()
		srv := New(repo, repoClass)
		err := srv.PostSchedule(inputData)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		repo.AssertExpectations(t)
	})

}

func TestGetMentorSchedule(t *testing.T) {
	repo := mocks.NewScheduleData(t)
	repoClass := mocks.NewClassData(t)

	resData := []schedules.Core{{ID: 1, Day: "Monday", StartTime: "13:00", EndTime: "14:00"}}

	t.Run("succes get all mentor schedule", func(t *testing.T) {
		repo.On("GetMentorSchedule", uint(1)).Return(resData, nil).Once()
		srv := New(repo, repoClass)
		res, err := srv.GetMentorSchedule(uint(1))
		assert.Nil(t, err)
		assert.Equal(t, len(resData), len(res))
		repo.AssertExpectations(t)
	})

	t.Run("schedule not found", func(t *testing.T) {
		repo.On("GetMentorSchedule", uint(1)).Return([]schedules.Core{}, errors.New("schedule not found")).Once()
		srv := New(repo, repoClass)
		res, err := srv.GetMentorSchedule(uint(1))
		assert.NotNil(t, err)
		assert.NotEqual(t, len(resData), len(res))
		assert.ErrorContains(t, err, "not found")
		repo.AssertExpectations(t)
	})

	t.Run("server problem", func(t *testing.T) {
		repo.On("GetMentorSchedule", uint(1)).Return([]schedules.Core{}, errors.New("server problem")).Once()
		srv := New(repo, repoClass)
		res, err := srv.GetMentorSchedule(uint(1))
		assert.NotNil(t, err)
		assert.NotEqual(t, len(resData), len(res))
		assert.ErrorContains(t, err, "server")
		repo.AssertExpectations(t)
	})

}

func TestDelete(t *testing.T) {

	repo := mocks.NewScheduleData(t)
	repoClass := mocks.NewClassData(t)

	t.Run("success delete mentor schedule", func(t *testing.T) {
		repo.On("Delete", uint(1), uint(1)).Return(nil).Once()

		srv := New(repo, repoClass)
		err := srv.Delete(uint(1), uint(1))
		assert.Nil(t, err)
		repo.AssertExpectations(t)

	})

	t.Run("data not found", func(t *testing.T) {
		repo.On("Delete", uint(1), uint(1)).Return(errors.New("data not found")).Once()
		srv := New(repo, repoClass)
		err := srv.Delete(uint(1), uint(1))
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not found")
		repo.AssertExpectations(t)
	})

}
