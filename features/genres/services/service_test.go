package services

import (
	"errors"
	"musiclab-be/features/genres"
	"musiclab-be/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAddMentorGenre(t *testing.T) {
	repo := mocks.NewGenreData(t)
	inputData := genres.Core{Name: "Jazz", GenreID: 1}

	t.Run("success add mentor genre", func(t *testing.T) {
		repo.On("AddMentorGenre", inputData).Return(nil).Once()
		srv := New(repo)
		err := srv.AddMentorGenre(inputData)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("genre not found", func(t *testing.T) {
		repo.On("AddMentorGenre", mock.Anything).Return(errors.New("not found"))
		srv := New(repo)
		err := srv.AddMentorGenre(inputData)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not found")
		repo.AssertExpectations(t)
	})

}

func TestGetGenre(t *testing.T) {
	repo := mocks.NewGenreData(t)
	resData := []genres.Core{{ID: uint(1), Name: "Jazz"}}

	t.Run("success get all genre", func(t *testing.T) {
		repo.On("GetGenre").Return(resData, nil).Once()
		srv := New(repo)
		res, err := srv.GetGenre()
		assert.Nil(t, err)
		assert.Equal(t, len(resData), len(res))
		repo.AssertExpectations(t)
	})

	t.Run("genre not found", func(t *testing.T) {
		repo.On("GetGenre").Return([]genres.Core{}, errors.New("genre not found")).Once()
		srv := New(repo)
		res, err := srv.GetGenre()
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not found")
		assert.Equal(t, 0, len(res))
		repo.AssertExpectations(t)
	})

	t.Run("problem with the server", func(t *testing.T) {
		repo.On("GetGenre").Return([]genres.Core{}, errors.New("problem with the server")).Once()
		srv := New(repo)
		res, err := srv.GetGenre()
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		assert.Equal(t, 0, len(res))
		repo.AssertExpectations(t)
	})

}

func TestGetMentorGenre(t *testing.T) {
	repo := mocks.NewGenreData(t)
	resData := []genres.MentorGenreCore{{ID: uint(1)}}

	t.Run("success get mentor genre", func(t *testing.T) {
		repo.On("GetMentorGenre", uint(1)).Return(resData, nil).Once()
		srv := New(repo)
		res, err := srv.GetMentorGenre(uint(1))
		assert.Nil(t, err)
		assert.Equal(t, len(resData), len(res))
		repo.AssertExpectations(t)
	})

	t.Run("problem with the server", func(t *testing.T) {
		repo.On("GetMentorGenre", uint(1)).Return([]genres.MentorGenreCore{}, errors.New("problem with the server")).Once()
		srv := New(repo)
		res, err := srv.GetMentorGenre(uint(1))
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		assert.Equal(t, 0, len(res))
		repo.AssertExpectations(t)
	})

}

func TestDelete(t *testing.T) {
	repo := mocks.NewGenreData(t)

	t.Run("success delete mentor genre", func(t *testing.T) {
		repo.On("Delete", uint(1), uint(1)).Return(nil).Once()

		srv := New(repo)
		err := srv.Delete(uint(1), uint(1))
		assert.Nil(t, err)
		repo.AssertExpectations(t)

	})

	t.Run("problem with the server", func(t *testing.T) {
		repo.On("Delete", uint(1), uint(1)).Return(errors.New("problem with the server")).Once()
		srv := New(repo)
		err := srv.Delete(uint(1), uint(1))
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		repo.AssertExpectations(t)
	})

}
