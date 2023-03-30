package services

import (
	"errors"
	"log"
	"mime/multipart"
	"musiclab-be/features/classes"
	"musiclab-be/mocks"
	"musiclab-be/utils/helper"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPostClass(t *testing.T) {
	repo := mocks.NewClassData(t)

	filePath := filepath.Join("..", "..", "..", "ERD_MusicLab@2x.png")
	imageTrue, err := os.Open(filePath)
	if err != nil {
		log.Println(err.Error())
	}
	imageTrueCnv := &multipart.FileHeader{
		Filename: imageTrue.Name(),
	}

	inputData := classes.Core{
		ID:      0,
		Name:    "Guitar Class",
		Level:   "Basic",
		ForWhom: "Newbie",
	}

	t.Run("success post class", func(t *testing.T) {
		repo.On("PostClass", inputData).Return(nil).Once()
		srv := New(repo)
		err := srv.PostClass(*imageTrueCnv, inputData)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("cannot post class", func(t *testing.T) {
		repo.On("PostClass", mock.Anything).Return(errors.New("cannot post class"))
		srv := New(repo)
		err := srv.PostClass(*imageTrueCnv, inputData)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		repo.AssertExpectations(t)
	})

	t.Run("invalid file validation", func(t *testing.T) {
		filePathFake := filepath.Join("..", "..", "..", "test.csv")
		headerFake, err := helper.UnitTestingUploadFileMock(filePathFake)
		if err != nil {
			log.Panic("from file header", err.Error())
		}
		srv := New(repo)
		err = srv.PostClass(*headerFake, inputData)
		assert.ErrorContains(t, err, "type")
		repo.AssertExpectations(t)

	})

}

func TestGetMentorClass(t *testing.T) {
	repo := mocks.NewClassData(t)
	filePath := filepath.Join("..", "..", "..", "ERD_MusicLab@2x.png")
	imageTrue, err := os.Open(filePath)
	if err != nil {
		log.Println(err.Error())
	}
	imageTrueCnv := &multipart.FileHeader{
		Filename: imageTrue.Name(),
	}

	resData := []classes.Core{{
		ID:      1,
		Name:    "Guitar Class",
		Level:   "Basic",
		Image:   imageTrueCnv.Filename,
		ForWhom: "Newbie",
	}}

	var page int = 1
	var limit int = 6
	offset := (page - 1) * limit
	t.Run("success get all mentor class", func(t *testing.T) {
		repo.On("GetMentorClass", uint(1), limit, offset).Return(resData, nil).Once()
		srv := New(repo)
		res, err := srv.GetMentorClass(uint(1), page, limit)
		assert.Nil(t, err)
		assert.Equal(t, len(resData), len(res))
		repo.AssertExpectations(t)
	})

	t.Run("class not found", func(t *testing.T) {
		repo.On("GetMentorClass", uint(1), limit, offset).Return([]classes.Core{}, errors.New("class not found")).Once()

		srv := New(repo)
		res, err := srv.GetMentorClass(uint(1), page, limit)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not found")
		assert.Equal(t, 0, len(res))
		repo.AssertExpectations(t)
	})

	t.Run("there is a problem with the server", func(t *testing.T) {
		repo.On("GetMentorClass", uint(1), limit, offset).Return([]classes.Core{}, errors.New("there is a problem with the server")).Once()

		srv := New(repo)
		res, err := srv.GetMentorClass(uint(1), page, limit)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		assert.Equal(t, 0, len(res))
		repo.AssertExpectations(t)
	})

}

func TestGetMentorClassDetail(t *testing.T) {
	repo := mocks.NewClassData(t)
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

	t.Run("success get mentor class detail", func(t *testing.T) {
		repo.On("GetMentorClassDetail", uint(1)).Return(resData, nil).Once()
		srv := New(repo)
		res, err := srv.GetMentorClassDetail(uint(1))
		assert.Nil(t, err)
		assert.Equal(t, resData.ID, res.ID)
		repo.AssertExpectations(t)
	})

	t.Run("data not found", func(t *testing.T) {
		repo.On("GetMentorClassDetail", uint(1)).Return(classes.Core{}, errors.New("data not found")).Once()
		srv := New(repo)
		res, err := srv.GetMentorClassDetail(uint(1))
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not found")
		assert.NotEqual(t, 0, res.ID)
		repo.AssertExpectations(t)
	})
	t.Run("server problem", func(t *testing.T) {
		repo.On("GetMentorClassDetail", uint(1)).Return(classes.Core{}, errors.New("server problem")).Once()
		srv := New(repo)
		res, err := srv.GetMentorClassDetail(uint(1))
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		assert.NotEqual(t, 0, res.ID)
		repo.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	repo := mocks.NewClassData(t)
	filePath := filepath.Join("..", "..", "..", "ERD_MusicLab@2x.png")
	imageTrue, err := os.Open(filePath)
	if err != nil {
		log.Println(err.Error())
	}
	imageTrueCnv := &multipart.FileHeader{
		Filename: imageTrue.Name(),
	}

	inputData := classes.Core{
		ID:      1,
		Name:    "Guitar Class",
		Level:   "Basic",
		ForWhom: "Newbie",
	}

	t.Run("success update class", func(t *testing.T) {
		repo.On("Update", uint(1), uint(1), inputData).Return(nil).Once()
		srv := New(repo)
		err := srv.Update(uint(1), uint(1), *imageTrueCnv, inputData)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("invalid file validation", func(t *testing.T) {
		filePathFake := filepath.Join("..", "..", "..", "test.csv")
		headerFake, err := helper.UnitTestingUploadFileMock(filePathFake)
		if err != nil {
			log.Panic("from file header", err.Error())
		}
		srv := New(repo)
		err = srv.Update(uint(1), uint(1), *headerFake, inputData)
		assert.ErrorContains(t, err, "validate")
		repo.AssertExpectations(t)

	})

	t.Run("data not found", func(t *testing.T) {
		repo.On("Update", uint(1), uint(1), inputData).Return(errors.New("data not found")).Once()
		srv := New(repo)
		err := srv.Update(uint(1), uint(1), *imageTrueCnv, inputData)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not found")
		repo.AssertExpectations(t)
	})

	t.Run("internal server error", func(t *testing.T) {
		repo.On("Update", uint(1), uint(1), inputData).Return(errors.New("internal server error")).Once()
		srv := New(repo)
		err := srv.Update(uint(1), uint(1), *imageTrueCnv, inputData)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		repo.AssertExpectations(t)
	})

}

func TestDelete(t *testing.T) {
	repo := mocks.NewClassData(t)

	t.Run("success delete book", func(t *testing.T) {
		repo.On("Delete", uint(1), uint(1)).Return(nil).Once()
		srv := New(repo)
		err := srv.Delete(uint(1), 1)
		assert.Nil(t, err)
		repo.AssertExpectations(t)

	})

	t.Run("data not found", func(t *testing.T) {
		repo.On("Delete", uint(2), uint(2)).Return(errors.New("data not found")).Once()

		srv := New(repo)
		err := srv.Delete(uint(2), 2)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "found")
		repo.AssertExpectations(t)
	})
}
