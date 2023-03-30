package services

import (
	"errors"
	"musiclab-be/features/students"
	"musiclab-be/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	mock_data_student = students.Core{
		ID:                   1,
		Name:                 "Muhammad Ali",
		Email:                "ali@mail.com",
		Password:             "$2a$14$J4AF0twBNp2Pxx/LY5McIu/0v0FEvP.T7TOU/ozo.afMSDA03aBZ6",
		NewPassword:          "",
		ConfirmationPassword: "",
		Sex:                  "",
		Address:              "",
	}
)

func TestDelete(t *testing.T) {
	repo := mocks.NewStudentData(t)

	t.Run("success deactivate student", func(t *testing.T) {
		repo.On("Delete", uint(1)).Return(nil).Once()

		srv := New(repo)
		err := srv.Delete(uint(1))
		assert.Nil(t, err)
		repo.AssertExpectations(t)

	})

	t.Run("data not found", func(t *testing.T) {
		repo.On("Delete", uint(1)).Return(errors.New("data not found")).Once()
		srv := New(repo)
		err := srv.Delete(uint(1))
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not found")
		repo.AssertExpectations(t)
	})

}

func TestSelectProfile(t *testing.T) {
	repo := mocks.NewStudentData(t)

	resData := students.Core{
		ID:    1,
		Name:  "Alif Muhamad Hafidz",
		Email: "alif@gmail.com",
		Phone: "0808",
	}

	t.Run("succes get student profile", func(t *testing.T) {
		repo.On("SelectProfile", uint(1)).Return(resData, nil).Once()
		srv := New(repo)
		res, err := srv.SelectProfile(uint(1))
		assert.Nil(t, err)
		assert.Equal(t, resData.ID, res.ID)
		repo.AssertExpectations(t)
	})

	t.Run("error select data", func(t *testing.T) {
		repo.On("SelectProfile", uint(1)).Return(students.Core{}, errors.New("error select data")).Once()
		srv := New(repo)
		res, err := srv.SelectProfile(uint(1))
		assert.NotNil(t, err)
		assert.NotEqual(t, resData.ID, res.ID)
		assert.ErrorContains(t, err, "error select data")
		repo.AssertExpectations(t)
	})

}

func TestUpdateData(t *testing.T) {
	repo := mocks.NewStudentData(t)

	inputData := students.Core{
		ID:    1,
		Name:  "Alif Muhamad Hafidz",
		Email: "alif@gmail.com",
		Phone: "0808",
	}

	t.Run("success updating student account", func(t *testing.T) {
		repo.On("UpdateData", uint(1), inputData).Return(nil).Once()
		srv := New(repo)
		err := srv.UpdateData(uint(1), inputData)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("error update data", func(t *testing.T) {
		repo.On("UpdateData", uint(1), inputData).Return(errors.New("error update data")).Once()
		srv := New(repo)
		err := srv.UpdateData(uint(1), inputData)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "error update data")
		repo.AssertExpectations(t)
	})

}

func TestUpdatePassword(t *testing.T) {
	core := mock_data_student
	id := mock_data_student.ID
	repo := mocks.NewStudentData(t)
	input := students.Core{
		Password:             "thegreatest",
		NewPassword:          "goat",
		ConfirmationPassword: "goat",
	}

	t.Run("success update student password", func(t *testing.T) {
		repo.On("UpdateData", mock.Anything, mock.Anything).Return(nil).Once()
		repo.On("SelectProfile", mock.Anything).Return(core, nil).Once()

		srv := New(repo)
		err := srv.UpdatePassword(id, input)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("failed when select data", func(t *testing.T) {
		repo.On("SelectProfile", mock.Anything).Return(students.Core{}, errors.New("error select")).Once()

		srv := New(repo)
		err := srv.UpdatePassword(id, input)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("password not matched", func(t *testing.T) {
		input := students.Core{
			Password:             "asdasd",
			NewPassword:          "goat",
			ConfirmationPassword: "asdsad",
		}
		repo.On("SelectProfile", mock.Anything).Return(core, nil).Once()

		srv := New(repo)
		err := srv.UpdatePassword(id, input)
		assert.NotNil(t, err)
		assert.Equal(t, "password not matched", err.Error())
		repo.AssertExpectations(t)
	})

	t.Run("failed when update data", func(t *testing.T) {
		repo.On("SelectProfile", mock.Anything).Return(core, nil).Once()
		repo.On("UpdateData", mock.Anything, mock.Anything).Return(errors.New("error update")).Once()

		srv := New(repo)
		err := srv.UpdatePassword(id, input)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("failed validate", func(t *testing.T) {
		inputWrong := students.Core{
			Password:             "",
			NewPassword:          "",
			ConfirmationPassword: "",
		}
		srv := New(repo)
		err := srv.UpdatePassword(id, inputWrong)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}
