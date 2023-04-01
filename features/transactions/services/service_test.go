package services

import (
	"errors"
	"musiclab-be/features/classes"
	"musiclab-be/features/schedules"
	"musiclab-be/features/students"
	"musiclab-be/features/transactions"
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
	mock_schedule = schedules.Core{
		ID:        1,
		MentorID:  1,
		Day:       "Monday",
		StartTime: "13:00",
		EndTime:   "14:00",
	}
	mock_student = students.Core{
		Name:       "Student Musiclab",
		Email:      "studentmusiclab@gmail.com",
		Password:   "$2a$14$jLh92bx2atws9f.IHTFI..LZ7U6VDKPmd.tKjFtj6hXjnIZXvxEtu",
		Role:       "Student",
		TokenOauth: "ya29.a0Ael9sCMVbsszF3oVBGNiZdQJ0baKVZ5vahuavcznUlps9DFdIfeJ7Z7OK3yPyzHlae_Ly5MtWdCG4NgLgA6GkuUG_Pqktf7reTm1VTGuHMruEj9M1i1X9EKb_RF0jygvluLOJ5_VXQ19u9ecJDj52FTtMu_WaCgYKAU0SARASFQF4udJhRo-F0XIqzdmRCeddr7oPUA0163",
	}
)

func TestUpdateTransaction(t *testing.T) {
	repoTransaction := new(mocks.TransactionData)
	repoStudent := new(mocks.StudentData)
	repoClass := new(mocks.ClassData)

	t.Run("Success Update", func(t *testing.T) {
		input := transactions.Core{
			Status: "settlement",
		}
		repoTransaction.On("UpdateTransaction", mock.Anything).Return(nil).Once()

		srv := New(repoTransaction, repoStudent, repoClass)
		err := srv.UpdateTransaction(input)
		assert.Nil(t, err)
		repoTransaction.AssertExpectations(t)
	})

	t.Run("Failed Update", func(t *testing.T) {
		input := transactions.Core{
			Status: "settlement",
		}
		repoTransaction.On("UpdateTransaction", mock.Anything).Return(errors.New("error")).Once()

		srv := New(repoTransaction, repoStudent, repoClass)
		err := srv.UpdateTransaction(input)
		assert.NotNil(t, err)
		repoTransaction.AssertExpectations(t)
	})
}

func TestGetMentorTransaction(t *testing.T) {
	repoTransaction := new(mocks.TransactionData)
	repoStudent := new(mocks.StudentData)
	repoClass := new(mocks.ClassData)

	returnData := []transactions.Core{
		{
			ID:        1,
			Status:    "settlement",
			StartDate: time.Date(2023, 4, 1, 0, 0, 0, 0, time.UTC),
			EndDate:   time.Date(2023, 5, 1, 0, 0, 0, 0, time.UTC),
			Price:     100000,
			Student: transactions.Student{
				Name: "Bruno Mars",
			},
			Class: transactions.Class{
				Name: "How to Sing Like Celine Dion",
			},
		},
	}

	t.Run("Success Get Transaction", func(t *testing.T) {
		repoTransaction.On("GetMentorTransaction", mock.Anything, mock.Anything, mock.Anything).Return(returnData, nil).Once()

		srv := New(repoTransaction, repoStudent, repoClass)
		res, err := srv.GetMentorTransaction(1, 1, 15)
		assert.Nil(t, err)
		assert.Equal(t, returnData, res)
		repoTransaction.AssertExpectations(t)
	})

	t.Run("Failed Get Transaction", func(t *testing.T) {
		repoTransaction.On("GetMentorTransaction", mock.Anything, mock.Anything, mock.Anything).Return([]transactions.Core{}, errors.New("error")).Once()

		srv := New(repoTransaction, repoStudent, repoClass)
		res, err := srv.GetMentorTransaction(1, 1, 15)
		assert.NotNil(t, err)
		assert.Equal(t, []transactions.Core{}, res)
		repoTransaction.AssertExpectations(t)
	})
}

func TestGetStudentTransaction(t *testing.T) {
	repoTransaction := new(mocks.TransactionData)
	repoStudent := new(mocks.StudentData)
	repoClass := new(mocks.ClassData)

	returnData := []transactions.Core{
		{
			ID:        1,
			Status:    "settlement",
			StartDate: time.Date(2023, 4, 1, 0, 0, 0, 0, time.UTC),
			EndDate:   time.Date(2023, 5, 1, 0, 0, 0, 0, time.UTC),
			Price:     100000,
			Mentor: transactions.Mentor{
				Name: "Celine Dion",
			},
			Class: transactions.Class{
				Name: "How to Sing Like Celine Dion",
			},
		},
	}

	t.Run("Success Get Transaction", func(t *testing.T) {
		repoTransaction.On("GetStudentTransaction", mock.Anything, mock.Anything, mock.Anything).Return(returnData, nil).Once()

		srv := New(repoTransaction, repoStudent, repoClass)
		res, err := srv.GetStudentTransaction(1, 1, 15)
		assert.Nil(t, err)
		assert.Equal(t, returnData, res)
		repoTransaction.AssertExpectations(t)
	})

	t.Run("Failed Get Transaction", func(t *testing.T) {
		repoTransaction.On("GetStudentTransaction", mock.Anything, mock.Anything, mock.Anything).Return([]transactions.Core{}, errors.New("error")).Once()

		srv := New(repoTransaction, repoStudent, repoClass)
		res, err := srv.GetStudentTransaction(1, 1, 15)
		assert.NotNil(t, err)
		assert.Equal(t, []transactions.Core{}, res)
		repoTransaction.AssertExpectations(t)
	})
}

func TestMakeTransaction(t *testing.T) {
	repoTransaction := new(mocks.TransactionData)
	repoStudent := new(mocks.StudentData)
	repoClass := new(mocks.ClassData)

	input := transactions.Core{
		ClassID:    1,
		ScheduleID: 1,
		StartDate:  time.Date(2023, 4, 1, 0, 0, 0, 0, time.UTC),
	}

	dataReturn := transactions.Core{
		PaymentUrl: "",
	}

	t.Run("Success Make Transaction", func(t *testing.T) {
		repoClass.On("GetMentorClassDetail", mock.Anything).Return(mock_class, nil).Once()
		repoStudent.On("SelectProfile", mock.Anything).Return(mock_student, nil).Once()

		srv := New(repoTransaction, repoStudent, repoClass)
		res, err := srv.MakeTransaction(input)
		assert.NotNil(t, err)
		assert.Equal(t, dataReturn, res)
		repoTransaction.AssertExpectations(t)
	})

	t.Run("Failed Date Validation", func(t *testing.T) {
		input := transactions.Core{
			ClassID:    1,
			ScheduleID: 1,
			StartDate:  time.Date(2023, 3, 1, 0, 0, 0, 0, time.UTC),
		}

		srv := New(repoTransaction, repoStudent, repoClass)
		res, err := srv.MakeTransaction(input)
		assert.NotNil(t, err)
		assert.Equal(t, dataReturn, res)
		repoTransaction.AssertExpectations(t)
	})

	t.Run("Error Select Class", func(t *testing.T) {
		repoClass.On("GetMentorClassDetail", mock.Anything).Return(classes.Core{}, errors.New("error")).Once()

		srv := New(repoTransaction, repoStudent, repoClass)
		res, err := srv.MakeTransaction(input)
		assert.NotNil(t, err)
		assert.Equal(t, dataReturn, res)
		repoTransaction.AssertExpectations(t)
	})

	t.Run("Error Select Student", func(t *testing.T) {
		repoClass.On("GetMentorClassDetail", mock.Anything).Return(mock_class, nil).Once()
		repoStudent.On("SelectProfile", mock.Anything).Return(students.Core{}, errors.New("error")).Once()

		srv := New(repoTransaction, repoStudent, repoClass)
		res, err := srv.MakeTransaction(input)
		assert.NotNil(t, err)
		assert.Equal(t, dataReturn, res)
		repoTransaction.AssertExpectations(t)
	})
}
