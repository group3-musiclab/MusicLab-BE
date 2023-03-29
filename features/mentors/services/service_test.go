package services

import (
	"errors"
	"log"
	"mime/multipart"
	"musiclab-be/features/mentors"
	"musiclab-be/mocks"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	mock_data_mentor = mentors.Core{
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

func TestSelectAllByRating(t *testing.T) {
	repo := mocks.NewMentorData(t)
	filePath := filepath.Join("..", "..", "..", "ERD_MusicLab.jpg")
	imageTrue, err := os.Open(filePath)
	if err != nil {
		log.Println(err.Error())
	}
	imageTrueCnv := &multipart.FileHeader{
		Filename: imageTrue.Name(),
	}

	resData := []mentors.Core{{
		ID:        1,
		Avatar:    imageTrueCnv.Filename,
		Name:      "Alif",
		AvgRating: 3.4,
	}}

	t.Run("succes get mentor by rating", func(t *testing.T) {
		repo.On("SelectAllByRating").Return(resData, nil).Once()
		srv := New(repo)
		res, err := srv.SelectAllByRating()
		assert.Nil(t, err)
		assert.Equal(t, len(resData), len(res))
		repo.AssertExpectations(t)
	})

	t.Run("error read data", func(t *testing.T) {
		repo.On("SelectAllByRating").Return([]mentors.Core{}, errors.New("error read data")).Once()
		srv := New(repo)
		res, err := srv.SelectAllByRating()
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "error read data")
		assert.Equal(t, 0, len(res))
		repo.AssertExpectations(t)
	})

}

func TestSelectAll(t *testing.T) {
	repo := mocks.NewMentorData(t)
	filePath := filepath.Join("..", "..", "..", "ERD_MusicLab.jpg")
	imageTrue, err := os.Open(filePath)
	if err != nil {
		log.Println(err.Error())
	}
	imageTrueCnv := &multipart.FileHeader{
		Filename: imageTrue.Name(),
	}

	resData := []mentors.Core{{
		ID:        1,
		Avatar:    imageTrueCnv.Filename,
		Name:      "Alif",
		AvgRating: 3.4,
	}}

	var page int = 1
	var limit int = 6
	offset := (page - 1) * limit

	t.Run("succes get all mentor", func(t *testing.T) {
		repo.On("SelectAll", limit, offset).Return(resData, nil).Once()
		srv := New(repo)
		res, err := srv.SelectAll(page, limit)
		assert.Nil(t, err)
		assert.Equal(t, len(resData), len(res))
		repo.AssertExpectations(t)
	})

	t.Run("error read data", func(t *testing.T) {
		repo.On("SelectAll", limit, offset).Return([]mentors.Core{}, errors.New("error read data")).Once()
		srv := New(repo)
		res, err := srv.SelectAll(page, limit)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "error read data")
		assert.Equal(t, 0, len(res))
		repo.AssertExpectations(t)
	})

}

func TestDelete(t *testing.T) {
	repo := mocks.NewMentorData(t)

	t.Run("success deactivate mentor", func(t *testing.T) {
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

func TestInsertCredential(t *testing.T) {
	repo := mocks.NewMentorData(t)

	inputData := mentors.CredentialCore{
		ID:   0,
		Name: "Bass Certificate",
		Type: "International Certificate",
	}

	t.Run("success add credential", func(t *testing.T) {
		repo.On("InsertCredential", inputData).Return(nil).Once()
		srv := New(repo)

		err := srv.InsertCredential(inputData)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("error insert data", func(t *testing.T) {
		repo.On("InsertCredential", inputData).Return(errors.New("error insert data")).Once()
		srv := New(repo)
		err := srv.InsertCredential(inputData)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "error insert data")
		repo.AssertExpectations(t)
	})

}

func TestUpdateData(t *testing.T) {
	repo := mocks.NewMentorData(t)

	inputData := mentors.Core{
		ID:    1,
		Name:  "Alif Muhamad Hafidz",
		Email: "alif@gmail.com",
		Phone: "0808",
	}

	t.Run("success updating mentor account", func(t *testing.T) {
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

func TestSelectProfile(t *testing.T) {
	repo := mocks.NewMentorData(t)

	resData := mentors.Core{
		ID:    1,
		Name:  "Alif Muhamad Hafidz",
		Email: "alif@gmail.com",
		Phone: "0808",
	}

	t.Run("succes get mentor profile", func(t *testing.T) {
		repo.On("SelectProfile", uint(1)).Return(resData, nil).Once()
		srv := New(repo)
		res, err := srv.SelectProfile(uint(1))
		assert.Nil(t, err)
		assert.Equal(t, resData.ID, res.ID)
		repo.AssertExpectations(t)
	})

	t.Run("error select data", func(t *testing.T) {
		repo.On("SelectProfile", uint(1)).Return(mentors.Core{}, errors.New("error select data")).Once()
		srv := New(repo)
		res, err := srv.SelectProfile(uint(1))
		assert.NotNil(t, err)
		assert.NotEqual(t, resData.ID, res.ID)
		assert.ErrorContains(t, err, "error select data")
		repo.AssertExpectations(t)
	})
}

// func TestUpdatePassword(t *testing.T) {
// 	repo := mocks.NewMentorData(t)
// 	core := mock_data_mentor
// 	id := mock_data_mentor.ID
// 	t.Run("success update mentor password", func(t *testing.T) {
// 		input := mentors.Core{
// 			Password:             "thegreatest",
// 			NewPassword:          "goat",
// 			ConfirmationPassword: "goat",
// 		}
// 		repo.On("UpdateData", id, input).Return(nil).Once()
// 		repo.On("SelectProfile", id).Return(core, nil).Once()
// 		srv := New(repo)
// 		err := srv.UpdatePassword(id, input)
// 		assert.Nil(t, err)
// 		repo.AssertExpectations(t)
// 	})

// }
