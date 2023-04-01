package services

import (
	"errors"
	"musiclab-be/features/chats"
	"musiclab-be/features/mentors"
	"musiclab-be/features/students"
	"musiclab-be/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	mock_mentor = mentors.Core{
		Name:       "MusicLab Sejahtera",
		Email:      "musiclabsejahtera@gmail.com",
		Password:   "$2a$14$C5UwSnlWEPs4Ga4GNfuPjubIMZgt.6CJKqOX3so4ZiqXHvhsLRLZO",
		Role:       "Mentor",
		TokenOauth: "ya29.a0Ael9sCOjHI7GhoSU5i7to6mJPqNKkwhZY_mncqEZufV6mwGhuCgLUprwGXaa1A3Y36F0mgnvzkqu8X84IeZghJ1GqG62SUKxPaq2gVKmdCNE9_FJILyyvv2k5ZBrlnfbsssqBeDJgzHMOx28zB4V1-2Pvp8aaCgYKAesSARISFQF4udJhsf7dNPhk_fbruZypjDwjig0163",
	}
	mock_student = students.Core{
		Name:       "Student Musiclab",
		Email:      "studentmusiclab@gmail.com",
		Password:   "$2a$14$jLh92bx2atws9f.IHTFI..LZ7U6VDKPmd.tKjFtj6hXjnIZXvxEtu",
		Role:       "Student",
		TokenOauth: "ya29.a0Ael9sCMVbsszF3oVBGNiZdQJ0baKVZ5vahuavcznUlps9DFdIfeJ7Z7OK3yPyzHlae_Ly5MtWdCG4NgLgA6GkuUG_Pqktf7reTm1VTGuHMruEj9M1i1X9EKb_RF0jygvluLOJ5_VXQ19u9ecJDj52FTtMu_WaCgYKAU0SARASFQF4udJhRo-F0XIqzdmRCeddr7oPUA0163",
	}
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

func TestGetByStudent(t *testing.T) {
	repoChat := new(mocks.ChatData)
	repoStudent := new(mocks.StudentData)
	repoMentor := new(mocks.MentorData)

	returnData := []chats.Core{
		{
			MentorID:  1,
			StudentID: 1,
			Student: chats.StudentCore{
				Avatar: "url-avatar",
				Name:   "Bruno Mars",
			},
		},
	}

	t.Run("Success Get By Student", func(t *testing.T) {
		repoChat.On("GetByStudent", mock.Anything, mock.Anything, mock.Anything).Return(returnData, nil).Once()

		srv := New(repoChat, repoMentor, repoStudent)
		res, err := srv.GetByStudent(1, 5, 1)
		assert.Nil(t, err)
		assert.Equal(t, returnData, res)
		repoChat.AssertExpectations(t)
	})

	t.Run("Failed Get By Student", func(t *testing.T) {
		repoChat.On("GetByStudent", mock.Anything, mock.Anything, mock.Anything).Return([]chats.Core{}, errors.New("error")).Once()

		srv := New(repoChat, repoMentor, repoStudent)
		res, err := srv.GetByStudent(1, 5, 1)
		assert.NotNil(t, err)
		assert.Equal(t, []chats.Core{}, res)
		repoChat.AssertExpectations(t)
	})
}

func TestInsertChat(t *testing.T) {
	repoChat := new(mocks.ChatData)
	repoStudent := new(mocks.StudentData)
	repoMentor := new(mocks.MentorData)

	t.Run("Failed Validation", func(t *testing.T) {
		input := chats.Core{
			MentorID: 1,
		}

		srv := New(repoChat, repoMentor, repoStudent)
		err := srv.InsertChat(input)
		assert.NotNil(t, err)
		repoChat.AssertExpectations(t)
	})

	t.Run("Success Insert as Mentor", func(t *testing.T) {
		input := chats.Core{
			MentorID:  1,
			StudentID: 1,
			Chat:      "Halo",
			Role:      "Mentor",
		}

		repoMentor.On("SelectProfile", mock.Anything).Return(mock_mentor, nil).Once()
		repoStudent.On("SelectProfile", mock.Anything).Return(mock_student, nil).Once()
		repoMentor.On("SelectProfile", mock.Anything).Return(mock_mentor, nil).Once()
		repoChat.On("InsertChat", mock.Anything).Return(nil).Once()

		srv := New(repoChat, repoMentor, repoStudent)
		err := srv.InsertChat(input)
		assert.Nil(t, err)
		repoChat.AssertExpectations(t)
	})

	t.Run("Success Insert as Student", func(t *testing.T) {
		input := chats.Core{
			MentorID:  1,
			StudentID: 1,
			Chat:      "Halo",
			Role:      "Student",
		}

		repoMentor.On("SelectProfile", mock.Anything).Return(mock_mentor, nil).Once()
		repoStudent.On("SelectProfile", mock.Anything).Return(mock_student, nil).Once()
		repoStudent.On("SelectProfile", mock.Anything).Return(mock_student, nil).Once()
		repoChat.On("InsertChat", mock.Anything).Return(nil).Once()

		srv := New(repoChat, repoMentor, repoStudent)
		err := srv.InsertChat(input)
		assert.Nil(t, err)
		repoChat.AssertExpectations(t)
	})

	t.Run("Invalid Id Mentor", func(t *testing.T) {
		input := chats.Core{
			MentorID:  1,
			StudentID: 1,
			Chat:      "Halo",
			Role:      "Student",
		}

		repoMentor.On("SelectProfile", mock.Anything).Return(mentors.Core{}, errors.New("error")).Once()

		srv := New(repoChat, repoMentor, repoStudent)
		err := srv.InsertChat(input)
		assert.NotNil(t, err)
		repoChat.AssertExpectations(t)
	})

	t.Run("Invalid Id Student", func(t *testing.T) {
		input := chats.Core{
			MentorID:  1,
			StudentID: 1,
			Chat:      "Halo",
			Role:      "Student",
		}

		repoMentor.On("SelectProfile", mock.Anything).Return(mock_mentor, nil).Once()
		repoStudent.On("SelectProfile", mock.Anything).Return(students.Core{}, errors.New("error")).Once()

		srv := New(repoChat, repoMentor, repoStudent)
		err := srv.InsertChat(input)
		assert.NotNil(t, err)
		repoChat.AssertExpectations(t)
	})

	t.Run("Failed Insert as Mentor", func(t *testing.T) {
		input := chats.Core{
			MentorID:  1,
			StudentID: 1,
			Chat:      "Halo",
			Role:      "Mentor",
		}

		repoMentor.On("SelectProfile", mock.Anything).Return(mock_mentor, nil).Once()
		repoStudent.On("SelectProfile", mock.Anything).Return(mock_student, nil).Once()
		repoMentor.On("SelectProfile", mock.Anything).Return(mentors.Core{}, errors.New("error")).Once()

		srv := New(repoChat, repoMentor, repoStudent)
		err := srv.InsertChat(input)
		assert.NotNil(t, err)
		repoChat.AssertExpectations(t)
	})

	t.Run("Failed Insert as Student", func(t *testing.T) {
		input := chats.Core{
			MentorID:  1,
			StudentID: 1,
			Chat:      "Halo",
			Role:      "Student",
		}

		repoMentor.On("SelectProfile", mock.Anything).Return(mock_mentor, nil).Once()
		repoStudent.On("SelectProfile", mock.Anything).Return(mock_student, nil).Once()
		repoStudent.On("SelectProfile", mock.Anything).Return(students.Core{}, errors.New("error")).Once()

		srv := New(repoChat, repoMentor, repoStudent)
		err := srv.InsertChat(input)
		assert.NotNil(t, err)
		repoChat.AssertExpectations(t)
	})

	t.Run("Failed Insert Chat", func(t *testing.T) {
		input := chats.Core{
			MentorID:  1,
			StudentID: 1,
			Chat:      "Halo",
			Role:      "Student",
		}

		repoMentor.On("SelectProfile", mock.Anything).Return(mock_mentor, nil).Once()
		repoStudent.On("SelectProfile", mock.Anything).Return(mock_student, nil).Once()
		repoStudent.On("SelectProfile", mock.Anything).Return(mock_student, nil).Once()
		repoChat.On("InsertChat", mock.Anything).Return(errors.New("error")).Once()

		srv := New(repoChat, repoMentor, repoStudent)
		err := srv.InsertChat(input)
		assert.NotNil(t, err)
		repoChat.AssertExpectations(t)
	})
}
