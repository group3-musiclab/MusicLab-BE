package services

import (
	"errors"
	"musiclab-be/features/instruments"
	"musiclab-be/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	mock_insert_instrument = instruments.MentorInstrumentCore{
		MentorID:     1,
		InstrumentID: 1,
	}
)

func TestDelete(t *testing.T) {
	repo := new(mocks.InstrumentData)

	t.Run("Success Delete", func(t *testing.T) {
		repo.On("Delete", mock.Anything, mock.Anything).Return(nil).Once()

		srv := New(repo)
		err := srv.Delete(1, 1)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Failed when func Delete return error", func(t *testing.T) {
		repo.On("Delete", mock.Anything, mock.Anything).Return(errors.New("error delete data")).Once()

		srv := New(repo)
		err := srv.Delete(1, 1)
		assert.NotNil(t, err)
		assert.Equal(t, "error delete data", err.Error())
		repo.AssertExpectations(t)
	})
}

func TestInsert(t *testing.T) {
	repo := new(mocks.InstrumentData)

	t.Run("Success Insert", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(nil).Once()

		srv := New(repo)
		err := srv.Insert(mock_insert_instrument)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Failed validate", func(t *testing.T) {
		inputData := instruments.MentorInstrumentCore{
			MentorID: 1,
		}
		srv := New(repo)
		err := srv.Insert(inputData)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Failed when func Insert return error", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(errors.New("error insert data")).Once()

		srv := New(repo)
		err := srv.Insert(mock_insert_instrument)
		assert.NotNil(t, err)
		assert.Equal(t, "error insert data", err.Error())
		repo.AssertExpectations(t)
	})
}
