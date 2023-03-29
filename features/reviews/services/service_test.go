package services

import (
	"errors"
	"musiclab-be/features/reviews"
	"musiclab-be/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostMentorReview(t *testing.T) {
	repo := mocks.NewReviewData(t)
	inputData := reviews.Core{ID: 1, Rating: 3, Comment: "Good"}

	t.Run("success post mentor review", func(t *testing.T) {
		repo.On("PostMentorReview", uint(1), inputData).Return(nil).Once()
		srv := New(repo)
		err := srv.PostMentorReview(uint(1), inputData)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("data not found", func(t *testing.T) {
		repo.On("PostMentorReview", uint(1), inputData).Return(errors.New("data not found")).Once()
		srv := New(repo)
		err := srv.PostMentorReview(uint(1), inputData)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not found")
		repo.AssertExpectations(t)
	})

	t.Run("server problem", func(t *testing.T) {
		repo.On("PostMentorReview", uint(1), inputData).Return(errors.New("server problem")).Once()
		srv := New(repo)
		err := srv.PostMentorReview(uint(1), inputData)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		repo.AssertExpectations(t)
	})

}

func TestGetMentorReview(t *testing.T) {
	repo := mocks.NewReviewData(t)
	resData := []reviews.Core{{ID: 1, Rating: 3, Comment: "Good"}}
	var page int = 1
	var limit int = 6
	offset := (page - 1) * limit

	t.Run("success get mentor review", func(t *testing.T) {
		repo.On("GetMentorReview", limit, offset, uint(1)).Return(resData, nil).Once()
		srv := New(repo)
		res, err := srv.GetMentorReview(page, limit, uint(1))
		assert.Nil(t, err)
		assert.Equal(t, len(resData), len(res))
		repo.AssertExpectations(t)
	})

	t.Run("review not found", func(t *testing.T) {
		repo.On("GetMentorReview", limit, offset, uint(1)).Return([]reviews.Core{}, errors.New("review not found")).Once()
		srv := New(repo)
		res, err := srv.GetMentorReview(page, limit, uint(1))
		assert.NotNil(t, err)
		assert.Equal(t, 0, len(res))
		assert.ErrorContains(t, err, "not found")
		repo.AssertExpectations(t)
	})

	t.Run("there is a problem with the server", func(t *testing.T) {
		repo.On("GetMentorReview", limit, offset, uint(1)).Return([]reviews.Core{}, errors.New("there is a problem with the server")).Once()
		srv := New(repo)
		res, err := srv.GetMentorReview(page, limit, uint(1))
		assert.NotNil(t, err)
		assert.Equal(t, 0, len(res))
		assert.ErrorContains(t, err, "server")
		repo.AssertExpectations(t)
	})

}
