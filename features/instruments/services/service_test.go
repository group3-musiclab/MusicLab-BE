package services

import (
	"errors"
	"musiclab-be/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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
