package services

import (
	"errors"
	"musiclab-be/features/chats"
	"musiclab-be/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAll(t *testing.T) {
	repoChat := new(mocks.ChatData)
	repoStudent := new(mocks.StudentData)
	repoMentor := new(mocks.MentorData)

	returnData := []chats.Core{
		{
			ID:         1,
			SenderName: "Bruno Mars",
			Chat:       "Halo",
		},
	}

	t.Run("Success Get All", func(t *testing.T) {
		repoChat.On("GetAll", mock.Anything, mock.Anything, mock.Anything).Return(returnData, nil).Once()

		srv := New(repoChat, repoMentor, repoStudent)
		res, err := srv.GetAll(1, 1)
		assert.Nil(t, err)
		assert.Equal(t, returnData, res)
		repoChat.AssertExpectations(t)
	})

	t.Run("Failed Get All", func(t *testing.T) {
		repoChat.On("GetAll", mock.Anything, mock.Anything, mock.Anything).Return([]chats.Core{}, errors.New("error")).Once()

		srv := New(repoChat, repoMentor, repoStudent)
		res, err := srv.GetAll(1, 1)
		assert.NotNil(t, err)
		assert.Equal(t, []chats.Core{}, res)
		repoChat.AssertExpectations(t)
	})
}
