package services

import (
	"errors"
	"musiclab-be/features/auth"
	"musiclab-be/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	mock_register_mentor = auth.Core{
		Name:     "Michael Jackson",
		Email:    "mj@mail.com",
		Password: "pop",
		Role:     "Mentor",
	}
	mock_register_student = auth.Core{
		Name:     "Michael Jackson",
		Email:    "mj@mail.com",
		Password: "pop",
		Role:     "Student",
	}
)

func TestRegister(t *testing.T) {
	repoAuth := new(mocks.AuthData)
	repoGoogle := new(mocks.GoogleAPI)
	repoTransaction := new(mocks.TransactionData)
	repoClass := new(mocks.ClassData)
	repoStudent := new(mocks.StudentData)
	repoSchedule := new(mocks.ScheduleData)
	repoMentor := new(mocks.MentorData)

	t.Run("Failed validate", func(t *testing.T) {
		inputData := auth.Core{
			Name:  "Michael Jackson",
			Email: "mj@mail.com",
		}
		srv := New(repoAuth, repoGoogle, repoTransaction, repoClass, repoStudent, repoSchedule, repoMentor)
		err := srv.Register(inputData)
		assert.NotNil(t, err)
		repoAuth.AssertExpectations(t)
	})

	t.Run("Only Letters Validation", func(t *testing.T) {
		inputData := auth.Core{
			Name:     "Michael Jackson 123",
			Email:    "mj@mail.com",
			Password: "pop",
			Role:     "Mentor",
		}

		srv := New(repoAuth, repoGoogle, repoTransaction, repoClass, repoStudent, repoSchedule, repoMentor)
		err := srv.Register(inputData)
		assert.NotNil(t, err)
		repoAuth.AssertExpectations(t)
	})

	t.Run("Duplicated Email", func(t *testing.T) {
		dataUser := auth.Core{
			Email: "mj@mail.com",
		}

		repoAuth.On("FindAccount", mock.Anything).Return(dataUser, nil).Once()

		srv := New(repoAuth, repoGoogle, repoTransaction, repoClass, repoStudent, repoSchedule, repoMentor)
		err := srv.Register(mock_register_mentor)
		assert.NotNil(t, err)
		repoAuth.AssertExpectations(t)
	})

	t.Run("Success Register Mentor", func(t *testing.T) {
		repoAuth.On("FindAccount", mock.Anything).Return(auth.Core{}, errors.New("")).Once()
		repoAuth.On("RegisterMentor", mock.Anything).Return(nil).Once()

		srv := New(repoAuth, repoGoogle, repoTransaction, repoClass, repoStudent, repoSchedule, repoMentor)
		err := srv.Register(mock_register_mentor)
		assert.Nil(t, err)
		repoAuth.AssertExpectations(t)
	})

	t.Run("Failed Register Mentor", func(t *testing.T) {
		repoAuth.On("FindAccount", mock.Anything).Return(auth.Core{}, errors.New("")).Once()
		repoAuth.On("RegisterMentor", mock.Anything).Return(errors.New("error register")).Once()

		srv := New(repoAuth, repoGoogle, repoTransaction, repoClass, repoStudent, repoSchedule, repoMentor)
		err := srv.Register(mock_register_mentor)
		assert.NotNil(t, err)
		repoAuth.AssertExpectations(t)
	})

	t.Run("Success Register Student", func(t *testing.T) {
		repoAuth.On("FindAccount", mock.Anything).Return(auth.Core{}, errors.New("")).Once()
		repoAuth.On("RegisterStudent", mock.Anything).Return(nil).Once()

		srv := New(repoAuth, repoGoogle, repoTransaction, repoClass, repoStudent, repoSchedule, repoMentor)
		err := srv.Register(mock_register_student)
		assert.Nil(t, err)
		repoAuth.AssertExpectations(t)
	})

	t.Run("Failed Register Student", func(t *testing.T) {
		repoAuth.On("FindAccount", mock.Anything).Return(auth.Core{}, errors.New("")).Once()
		repoAuth.On("RegisterStudent", mock.Anything).Return(errors.New("error register")).Once()

		srv := New(repoAuth, repoGoogle, repoTransaction, repoClass, repoStudent, repoSchedule, repoMentor)
		err := srv.Register(mock_register_student)
		assert.NotNil(t, err)
		repoAuth.AssertExpectations(t)
	})

	t.Run("Failed Role", func(t *testing.T) {
		inputData := auth.Core{
			Name:     "Michael Jackson",
			Email:    "mj@mail.com",
			Password: "pop",
			Role:     "Admin",
		}

		repoAuth.On("FindAccount", mock.Anything).Return(auth.Core{}, errors.New("")).Once()

		srv := New(repoAuth, repoGoogle, repoTransaction, repoClass, repoStudent, repoSchedule, repoMentor)
		err := srv.Register(inputData)
		assert.NotNil(t, err)
		repoAuth.AssertExpectations(t)
	})

}
