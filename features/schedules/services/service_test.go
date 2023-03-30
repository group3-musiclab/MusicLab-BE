package services

import (
	"errors"
	"log"
	"mime/multipart"
	"musiclab-be/features/classes"
	"musiclab-be/features/schedules"
	"musiclab-be/mocks"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCheckSchedule(t *testing.T) {
	repo := mocks.NewScheduleData(t)
	repoClass := mocks.NewClassData(t)
	filePath := filepath.Join("..", "..", "..", "ERD_MusicLab@2x.png")
	imageTrue, err := os.Open(filePath)
	if err != nil {
		log.Println(err.Error())
	}
	imageTrueCnv := &multipart.FileHeader{
		Filename: imageTrue.Name(),
	}

	resData := classes.Core{
		ID:      1,
		Name:    "Guitar Class",
		Level:   "Basic",
		Image:   imageTrueCnv.Filename,
		ForWhom: "Newbie",
	}

	startDate := "2023-03-31"
	StartDate, _ := time.Parse("2006-01-02", startDate)

	inputData := schedules.Core{ID: 1, ClassID: 1, Transaction: schedules.Transaction{ScheduleID: 1, StartDate: StartDate}}
	var row int64 = 1

	t.Run("success check schedule", func(t *testing.T) {
		repo.On("CheckSchedule", inputData).Return(row, nil).Once()
		repo.On("GetMentorClassDetail", uint(1)).Return(resData, nil).Once()
		srv := New(repo, repoClass)
		err := srv.CheckSchedule(inputData)
		assert.Nil(t, err)
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
