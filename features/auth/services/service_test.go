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
	mock_mentor_core = auth.Core{
		Name:     "Michael Jackson",
		Email:    "mj@mail.com",
		Password: "$2a$14$ZUTFIHo0iJEKuy7OcE0gge.RIgkU.iJjBum0E0QupmnZMRHMjjgQ.",
		Role:     "Mentor",
	}
	mock_student_core = auth.Core{
		Name:     "Bruno Mars",
		Email:    "bruno@mail.com",
		Password: "$2a$14$4blAj2MQJnMEdyB.Dd6LD.2e9KD75TOefD1VcLWSa4TqcxNi4kejy",
		Role:     "Student",
	}
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
	mock_login_mentor = auth.Core{
		Email:    "mj@mail.com",
		Password: "pop",
		Role:     "Mentor",
	}
	mock_login_student = auth.Core{
		Email:    "bruno@mail.com",
		Password: "mars",
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

func TestLogin(t *testing.T) {
	repoAuth := new(mocks.AuthData)
	repoGoogle := new(mocks.GoogleAPI)
	repoTransaction := new(mocks.TransactionData)
	repoClass := new(mocks.ClassData)
	repoStudent := new(mocks.StudentData)
	repoSchedule := new(mocks.ScheduleData)
	repoMentor := new(mocks.MentorData)

	t.Run("Failed validate", func(t *testing.T) {
		inputData := auth.Core{
			Email: "mj@mail.com",
		}
		srv := New(repoAuth, repoGoogle, repoTransaction, repoClass, repoStudent, repoSchedule, repoMentor)
		token, core, err := srv.Login(inputData)
		assert.NotNil(t, err)
		assert.Equal(t, "", token)
		assert.Equal(t, auth.Core{}, core)
		repoAuth.AssertExpectations(t)
	})

	t.Run("Login Mentor", func(t *testing.T) {
		repoAuth.On("LoginMentor", mock.Anything).Return(mock_mentor_core, nil).Once()

		srv := New(repoAuth, repoGoogle, repoTransaction, repoClass, repoStudent, repoSchedule, repoMentor)
		token, core, err := srv.Login(mock_login_mentor)
		assert.Nil(t, err)
		assert.Equal(t, token, token)
		assert.Equal(t, mock_mentor_core, core)
		repoAuth.AssertExpectations(t)
	})

	t.Run("Failed Login Mentor", func(t *testing.T) {
		repoAuth.On("LoginMentor", mock.Anything).Return(auth.Core{}, errors.New("error login")).Once()

		srv := New(repoAuth, repoGoogle, repoTransaction, repoClass, repoStudent, repoSchedule, repoMentor)
		token, core, err := srv.Login(mock_login_mentor)
		assert.NotNil(t, err)
		assert.Equal(t, "", token)
		assert.Equal(t, auth.Core{}, core)
		repoAuth.AssertExpectations(t)
	})

	t.Run("Login Student", func(t *testing.T) {
		repoAuth.On("LoginStudent", mock.Anything).Return(mock_student_core, nil).Once()

		srv := New(repoAuth, repoGoogle, repoTransaction, repoClass, repoStudent, repoSchedule, repoMentor)
		token, core, err := srv.Login(mock_login_student)
		assert.Nil(t, err)
		assert.Equal(t, token, token)
		assert.Equal(t, mock_student_core, core)
		repoAuth.AssertExpectations(t)
	})

	t.Run("Failed Login Student", func(t *testing.T) {
		repoAuth.On("LoginStudent", mock.Anything).Return(auth.Core{}, errors.New("error login")).Once()

		srv := New(repoAuth, repoGoogle, repoTransaction, repoClass, repoStudent, repoSchedule, repoMentor)
		token, core, err := srv.Login(mock_login_student)
		assert.NotNil(t, err)
		assert.Equal(t, "", token)
		assert.Equal(t, auth.Core{}, core)
		repoAuth.AssertExpectations(t)
	})

	t.Run("Failed Role", func(t *testing.T) {
		inputData := auth.Core{
			Name:     "Michael Jackson",
			Email:    "mj@mail.com",
			Password: "pop",
			Role:     "Admin",
		}

		srv := New(repoAuth, repoGoogle, repoTransaction, repoClass, repoStudent, repoSchedule, repoMentor)
		token, core, err := srv.Login(inputData)
		assert.NotNil(t, err)
		assert.Equal(t, "", token)
		assert.Equal(t, auth.Core{}, core)
		repoAuth.AssertExpectations(t)
	})

	t.Run("Failed Password", func(t *testing.T) {
		inputData := auth.Core{
			Email:    "bruno@mail.com",
			Password: "bruno",
			Role:     "Student",
		}

		repoAuth.On("LoginStudent", mock.Anything).Return(mock_student_core, nil).Once()

		srv := New(repoAuth, repoGoogle, repoTransaction, repoClass, repoStudent, repoSchedule, repoMentor)
		token, core, err := srv.Login(inputData)
		assert.NotNil(t, err)
		assert.Equal(t, "", token)
		assert.Equal(t, auth.Core{}, core)
		assert.Equal(t, "password not matched", err.Error())
		repoAuth.AssertExpectations(t)
	})
}
