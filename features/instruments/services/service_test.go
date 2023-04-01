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

func TestGetAll(t *testing.T) {
	repo := new(mocks.InstrumentData)
	returnData := []instruments.Core{
		{
			ID:   1,
			Name: "Suling",
		},
	}

	t.Run("Success Get All", func(t *testing.T) {
		repo.On("SelectAll").Return(returnData, nil).Once()

		srv := New(repo)
		response, err := srv.SelectAll()
		assert.Nil(t, err)
		assert.Equal(t, returnData[0].Name, response[0].Name)
		repo.AssertExpectations(t)
	})

	t.Run("Failed Get All", func(t *testing.T) {
		repo.On("SelectAll").Return([]instruments.Core{}, errors.New("error select")).Once()

		srv := New(repo)
		response, err := srv.SelectAll()
		assert.NotNil(t, err)
		assert.Equal(t, []instruments.Core{}, response)
		repo.AssertExpectations(t)
	})
}

func TestGetAllByMentorID(t *testing.T) {
	repo := new(mocks.InstrumentData)
	returnData := []instruments.MentorInstrumentCore{
		{

			InstrumentID: 1,
			Instrument: instruments.Core{
				Name: "Suling",
			},
		},
	}

	t.Run("Success Get All", func(t *testing.T) {
		repo.On("SelectAllByMentorID", mock.Anything).Return(returnData, nil).Once()

		srv := New(repo)
		response, err := srv.SelectAllByMentorID(1)
		assert.Nil(t, err)
		assert.Equal(t, returnData[0].Instrument.Name, response[0].Instrument.Name)
		repo.AssertExpectations(t)
	})

	t.Run("Failed Get All", func(t *testing.T) {
		repo.On("SelectAllByMentorID", mock.Anything).Return([]instruments.MentorInstrumentCore{}, errors.New("error select")).Once()

		srv := New(repo)
		response, err := srv.SelectAllByMentorID(1)
		assert.NotNil(t, err)
		assert.Equal(t, []instruments.MentorInstrumentCore{}, response)
		repo.AssertExpectations(t)
	})
}
