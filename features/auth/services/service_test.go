package services

import (
	"errors"
	"musiclab-be/features/auth"
	"musiclab-be/mocks"
	"musiclab-be/utils/helper"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/oauth2"
)

var (
	mock_mentor_core = auth.Core{
		Name:     "Michael Jackson",
		Email:    "mj@mail.com",
		Password: "$2a$14$ZUTFIHo0iJEKuy7OcE0gge.RIgkU.iJjBum0E0QupmnZMRHMjjgQ.",
		Role:     "Mentor",
	}
	mock_student_core = auth.Core{
		Name:     "Bruno Mars",
		Email:    "bruno@mail.com",
		Password: "$2a$14$4blAj2MQJnMEdyB.Dd6LD.2e9KD75TOefD1VcLWSa4TqcxNi4kejy",
		Role:     "Student",
	}
	mock_register_mentor = auth.Core{
		Name:     "Michael Jackson",
		Email:    "mj@mail.com",
		Password: "pop",
		Role:     "Mentor",
	}
	mock_register_student = auth.Core{
		Name:     "Michael Jackson",
		Email:    "mj@mail.com",
		Password: "pop",
		Role:     "Student",
	}
	mock_login_mentor = auth.Core{
		Email:    "mj@mail.com",
		Password: "pop",
		Role:     "Mentor",
	}
	mock_login_student = auth.Core{
		Email:    "bruno@mail.com",
		Password: "mars",
		Role:     "Student",
	}
	mock_google_core = helper.GoogleCore{
		Email: "musiclabsejahtera@gmail.com",
	}
	mock_mentor_musiclab = auth.Core{
		Name:       "MusicLab Sejahtera",
		Email:      "musiclabsejahtera@gmail.com",
		Password:   "$2a$14$C5UwSnlWEPs4Ga4GNfuPjubIMZgt.6CJKqOX3so4ZiqXHvhsLRLZO",
		Role:       "Mentor",
		TokenOauth: "ya29.a0Ael9sCOjHI7GhoSU5i7to6mJPqNKkwhZY_mncqEZufV6mwGhuCgLUprwGXaa1A3Y36F0mgnvzkqu8X84IeZghJ1GqG62SUKxPaq2gVKmdCNE9_FJILyyvv2k5ZBrlnfbsssqBeDJgzHMOx28zB4V1-2Pvp8aaCgYKAesSARISFQF4udJhsf7dNPhk_fbruZypjDwjig0163",
	}
	mock_student_musiclab = auth.Core{
		Name:       "Student Musiclab",
		Email:      "studentmusiclab@gmail.com",
		Password:   "$2a$14$jLh92bx2atws9f.IHTFI..LZ7U6VDKPmd.tKjFtj6hXjnIZXvxEtu",
		Role:       "Student",
		TokenOauth: "ya29.a0Ael9sCMVbsszF3oVBGNiZdQJ0baKVZ5vahuavcznUlps9DFdIfeJ7Z7OK3yPyzHlae_Ly5MtWdCG4NgLgA6GkuUG_Pqktf7reTm1VTGuHMruEj9M1i1X9EKb_RF0jygvluLOJ5_VXQ19u9ecJDj52FTtMu_WaCgYKAU0SARASFQF4udJhRo-F0XIqzdmRCeddr7oPUA0163",
	}
)

func TestRegister(t *testing.T) {
	repoAuth := new(mocks.AuthData)
	repoGoogle := new(mocks.GoogleAPI)
	repoTransaction := new(mocks.TransactionData)
	repoClass := new(mocks.ClassData)
	repoStudent := new(mocks.StudentData)
	repoSchedule := new(mocks.ScheduleData)
	repoMentor := new(mocks.MentorData)

	t.Run("Failed validate", func(t *testing.T) {
		inputData := auth.Core{
			Name:  "Michael Jackson",
			Email: "mj@mail.com",
		}
		srv := New(repoAuth, repoGoogle, repoTransaction, repoClass, repoStudent, repoSchedule, repoMentor)
		err := srv.Register(inputData)
		assert.NotNil(t, err)
		repoAuth.AssertExpectations(t)
	})

	t.Run("Only Letters Validation", func(t *testing.T) {
		inputData := auth.Core{
			Name:     "Michael Jackson 123",
			Email:    "mj@mail.com",
			Password: "pop",
			Role:     "Mentor",
		}

		srv := New(repoAuth, repoGoogle, repoTransaction, repoClass, repoStudent, repoSchedule, repoMentor)
		err := srv.Register(inputData)
		assert.NotNil(t, err)
		repoAuth.AssertExpectations(t)
	})

	t.Run("Duplicated Email", func(t *testing.T) {
		dataUser := auth.Core{
			Email: "mj@mail.com",
		}

		repoAuth.On("FindAccount", mock.Anything).Return(dataUser, nil).Once()

		srv := New(repoAuth, repoGoogle, repoTransaction, repoClass, repoStudent, repoSchedule, repoMentor)
		err := srv.Register(mock_register_mentor)
		assert.NotNil(t, err)
		repoAuth.AssertExpectations(t)
	})

	t.Run("Success Register Mentor", func(t *testing.T) {
		repoAuth.On("FindAccount", mock.Anything).Return(auth.Core{}, errors.New("")).Once()
		repoAuth.On("RegisterMentor", mock.Anything).Return(nil).Once()

		srv := New(repoAuth, repoGoogle, repoTransaction, repoClass, repoStudent, repoSchedule, repoMentor)
		err := srv.Register(mock_register_mentor)
		assert.Nil(t, err)
		repoAuth.AssertExpectations(t)
	})

	t.Run("Failed Register Mentor", func(t *testing.T) {
		repoAuth.On("FindAccount", mock.Anything).Return(auth.Core{}, errors.New("")).Once()
		repoAuth.On("RegisterMentor", mock.Anything).Return(errors.New("error register")).Once()

		srv := New(repoAuth, repoGoogle, repoTransaction, repoClass, repoStudent, repoSchedule, repoMentor)
		err := srv.Register(mock_register_mentor)
		assert.NotNil(t, err)
		repoAuth.AssertExpectations(t)
	})

	t.Run("Success Register Student", func(t *testing.T) {
		repoAuth.On("FindAccount", mock.Anything).Return(auth.Core{}, errors.New("")).Once()
		repoAuth.On("RegisterStudent", mock.Anything).Return(nil).Once()

		srv := New(repoAuth, repoGoogle, repoTransaction, repoClass, repoStudent, repoSchedule, repoMentor)
		err := srv.Register(mock_register_student)
		assert.Nil(t, err)
		repoAuth.AssertExpectations(t)
	})

	t.Run("Failed Register Student", func(t *testing.T) {
		repoAuth.On("FindAccount", mock.Anything).Return(auth.Core{}, errors.New("")).Once()
		repoAuth.On("RegisterStudent", mock.Anything).Return(errors.New("error register")).Once()

		srv := New(repoAuth, repoGoogle, repoTransaction, repoClass, repoStudent, repoSchedule, repoMentor)
		err := srv.Register(mock_register_student)
		assert.NotNil(t, err)
		repoAuth.AssertExpectations(t)
	})

	t.Run("Failed Role", func(t *testing.T) {
		inputData := auth.Core{
			Name:     "Michael Jackson",
			Email:    "mj@mail.com",
			Password: "pop",
			Role:     "Admin",
		}

		repoAuth.On("FindAccount", mock.Anything).Return(auth.Core{}, errors.New("")).Once()

		srv := New(repoAuth, repoGoogle, repoTransaction, repoClass, repoStudent, repoSchedule, repoMentor)
		err := srv.Register(inputData)
		assert.NotNil(t, err)
		repoAuth.AssertExpectations(t)
	})

}

func TestLogin(t *testing.T) {
	repoAuth := new(mocks.AuthData)
	repoGoogle := new(mocks.GoogleAPI)
	repoTransaction := new(mocks.TransactionData)
	repoClass := new(mocks.ClassData)
	repoStudent := new(mocks.StudentData)
	repoSchedule := new(mocks.ScheduleData)
	repoMentor := new(mocks.MentorData)

	t.Run("Failed validate", func(t *testing.T) {
		inputData := auth.Core{
			Email: "mj@mail.com",
		}
		srv := New(repoAuth, repoGoogle, repoTransaction, repoClass, repoStudent, repoSchedule, repoMentor)
		token, core, err := srv.Login(inputData)
		assert.NotNil(t, err)
		assert.Equal(t, "", token)
		assert.Equal(t, auth.Core{}, core)
		repoAuth.AssertExpectations(t)
	})

	t.Run("Login Mentor", func(t *testing.T) {
		repoAuth.On("LoginMentor", mock.Anything).Return(mock_mentor_core, nil).Once()

		srv := New(repoAuth, repoGoogle, repoTransaction, repoClass, repoStudent, repoSchedule, repoMentor)
		token, core, err := srv.Login(mock_login_mentor)
		assert.Nil(t, err)
		assert.Equal(t, token, token)
		assert.Equal(t, mock_mentor_core, core)
		repoAuth.AssertExpectations(t)
	})

	t.Run("Failed Login Mentor", func(t *testing.T) {
		repoAuth.On("LoginMentor", mock.Anything).Return(auth.Core{}, errors.New("error login")).Once()

		srv := New(repoAuth, repoGoogle, repoTransaction, repoClass, repoStudent, repoSchedule, repoMentor)
		token, core, err := srv.Login(mock_login_mentor)
		assert.NotNil(t, err)
		assert.Equal(t, "", token)
		assert.Equal(t, auth.Core{}, core)
		repoAuth.AssertExpectations(t)
	})

	t.Run("Login Student", func(t *testing.T) {
		repoAuth.On("LoginStudent", mock.Anything).Return(mock_student_core, nil).Once()

		srv := New(repoAuth, repoGoogle, repoTransaction, repoClass, repoStudent, repoSchedule, repoMentor)
		token, core, err := srv.Login(mock_login_student)
		assert.Nil(t, err)
		assert.Equal(t, token, token)
		assert.Equal(t, mock_student_core, core)
		repoAuth.AssertExpectations(t)
	})

	t.Run("Failed Login Student", func(t *testing.T) {
		repoAuth.On("LoginStudent", mock.Anything).Return(auth.Core{}, errors.New("error login")).Once()

		srv := New(repoAuth, repoGoogle, repoTransaction, repoClass, repoStudent, repoSchedule, repoMentor)
		token, core, err := srv.Login(mock_login_student)
		assert.NotNil(t, err)
		assert.Equal(t, "", token)
		assert.Equal(t, auth.Core{}, core)
		repoAuth.AssertExpectations(t)
	})

	t.Run("Failed Role", func(t *testing.T) {
		inputData := auth.Core{
			Name:     "Michael Jackson",
			Email:    "mj@mail.com",
			Password: "pop",
			Role:     "Admin",
		}

		srv := New(repoAuth, repoGoogle, repoTransaction, repoClass, repoStudent, repoSchedule, repoMentor)
		token, core, err := srv.Login(inputData)
		assert.NotNil(t, err)
		assert.Equal(t, "", token)
		assert.Equal(t, auth.Core{}, core)
		repoAuth.AssertExpectations(t)
	})

	t.Run("Failed Password", func(t *testing.T) {
		inputData := auth.Core{
			Email:    "bruno@mail.com",
			Password: "bruno",
			Role:     "Student",
		}

		repoAuth.On("LoginStudent", mock.Anything).Return(mock_student_core, nil).Once()

		srv := New(repoAuth, repoGoogle, repoTransaction, repoClass, repoStudent, repoSchedule, repoMentor)
		token, core, err := srv.Login(inputData)
		assert.NotNil(t, err)
		assert.Equal(t, "", token)
		assert.Equal(t, auth.Core{}, core)
		assert.Equal(t, "password not matched", err.Error())
		repoAuth.AssertExpectations(t)
	})
}

func TestLoginOauth(t *testing.T) {
	repoAuth := new(mocks.AuthData)
	repoGoogle := new(mocks.GoogleAPI)
	repoTransaction := new(mocks.TransactionData)
	repoClass := new(mocks.ClassData)
	repoStudent := new(mocks.StudentData)
	repoSchedule := new(mocks.ScheduleData)
	repoMentor := new(mocks.MentorData)

	t.Run("Failed validate", func(t *testing.T) {
		inputData := auth.Core{
			TokenOauth: "",
		}
		srv := New(repoAuth, repoGoogle, repoTransaction, repoClass, repoStudent, repoSchedule, repoMentor)
		token, core, err := srv.LoginOauth(inputData)
		assert.NotNil(t, err)
		assert.Equal(t, "", token)
		assert.Equal(t, auth.Core{}, core)
		assert.Equal(t, "token oauth cannot empty", err.Error())
		repoAuth.AssertExpectations(t)
	})

	t.Run("Failed Get User Info", func(t *testing.T) {
		inputData := auth.Core{
			TokenOauth: "asdfjsadkfskdkdsfsjs",
		}

		repoGoogle.On("GetUserInfo", mock.Anything).Return(helper.GoogleCore{}, errors.New("error")).Once()

		srv := New(repoAuth, repoGoogle, repoTransaction, repoClass, repoStudent, repoSchedule, repoMentor)
		token, core, err := srv.LoginOauth(inputData)
		assert.NotNil(t, err)
		assert.Equal(t, "", token)
		assert.Equal(t, auth.Core{}, core)
		assert.Equal(t, err.Error(), err.Error())
		repoAuth.AssertExpectations(t)
	})

	t.Run("Login Oauth Mentor", func(t *testing.T) {
		inputData := auth.Core{
			TokenOauth: mock_mentor_musiclab.TokenOauth,
		}

		repoGoogle.On("GetUserInfo", mock.Anything).Return(mock_google_core, nil).Once()
		repoAuth.On("FindAccount", mock.Anything).Return(mock_mentor_musiclab, nil).Once()
		repoMentor.On("UpdateData", mock.Anything, mock.Anything).Return(nil).Once()

		srv := New(repoAuth, repoGoogle, repoTransaction, repoClass, repoStudent, repoSchedule, repoMentor)
		token, core, err := srv.LoginOauth(inputData)
		assert.Nil(t, err)
		assert.Equal(t, token, token)
		assert.Equal(t, mock_mentor_musiclab, core)
		repoAuth.AssertExpectations(t)
	})

	t.Run("Login Oauth Student", func(t *testing.T) {
		inputData := auth.Core{
			TokenOauth: mock_student_musiclab.TokenOauth,
		}

		repoGoogle.On("GetUserInfo", mock.Anything).Return(mock_google_core, nil).Once()
		repoAuth.On("FindAccount", mock.Anything).Return(mock_student_musiclab, nil).Once()
		repoStudent.On("UpdateData", mock.Anything, mock.Anything).Return(nil).Once()

		srv := New(repoAuth, repoGoogle, repoTransaction, repoClass, repoStudent, repoSchedule, repoMentor)
		token, core, err := srv.LoginOauth(inputData)
		assert.Nil(t, err)
		assert.Equal(t, token, token)
		assert.Equal(t, mock_student_musiclab, core)
		repoAuth.AssertExpectations(t)
	})

	t.Run("Failed Update Data Mentor", func(t *testing.T) {
		inputData := auth.Core{
			TokenOauth: mock_mentor_musiclab.TokenOauth,
		}

		repoGoogle.On("GetUserInfo", mock.Anything).Return(mock_google_core, nil).Once()
		repoAuth.On("FindAccount", mock.Anything).Return(mock_mentor_musiclab, nil).Once()
		repoMentor.On("UpdateData", mock.Anything, mock.Anything).Return(errors.New("error")).Once()

		srv := New(repoAuth, repoGoogle, repoTransaction, repoClass, repoStudent, repoSchedule, repoMentor)
		token, core, err := srv.LoginOauth(inputData)
		assert.NotNil(t, err)
		assert.Equal(t, "", token)
		assert.Equal(t, auth.Core{}, core)
		assert.Equal(t, err.Error(), err.Error())
		repoAuth.AssertExpectations(t)
	})

	t.Run("Failed Update Data Student", func(t *testing.T) {
		inputData := auth.Core{
			TokenOauth: mock_student_musiclab.TokenOauth,
		}

		repoGoogle.On("GetUserInfo", mock.Anything).Return(mock_google_core, nil).Once()
		repoAuth.On("FindAccount", mock.Anything).Return(mock_student_musiclab, nil).Once()
		repoStudent.On("UpdateData", mock.Anything, mock.Anything).Return(errors.New("error")).Once()

		srv := New(repoAuth, repoGoogle, repoTransaction, repoClass, repoStudent, repoSchedule, repoMentor)
		token, core, err := srv.LoginOauth(inputData)
		assert.NotNil(t, err)
		assert.Equal(t, "", token)
		assert.Equal(t, auth.Core{}, core)
		assert.Equal(t, err.Error(), err.Error())
		repoAuth.AssertExpectations(t)
	})

	t.Run("Failed Find Account", func(t *testing.T) {
		inputData := auth.Core{
			TokenOauth: mock_student_musiclab.TokenOauth,
		}

		repoGoogle.On("GetUserInfo", mock.Anything).Return(mock_google_core, nil).Once()
		repoAuth.On("FindAccount", mock.Anything).Return(auth.Core{}, errors.New("error")).Once()

		srv := New(repoAuth, repoGoogle, repoTransaction, repoClass, repoStudent, repoSchedule, repoMentor)
		token, core, err := srv.LoginOauth(inputData)
		assert.NotNil(t, err)
		assert.Equal(t, "", token)
		assert.Equal(t, auth.Core{}, core)
		assert.Equal(t, err.Error(), err.Error())
		repoAuth.AssertExpectations(t)
	})
}

func TestRedirectGoogleCallback(t *testing.T) {
	repoAuth := new(mocks.AuthData)
	repoGoogle := new(mocks.GoogleAPI)
	repoTransaction := new(mocks.TransactionData)
	repoClass := new(mocks.ClassData)
	repoStudent := new(mocks.StudentData)
	repoSchedule := new(mocks.ScheduleData)
	repoMentor := new(mocks.MentorData)

	t.Run("Redirect Oauth Mentor", func(t *testing.T) {
		tokenOauth := &oauth2.Token{
			AccessToken: mock_mentor_musiclab.TokenOauth,
		}

		repoGoogle.On("GetToken", mock.Anything).Return(tokenOauth, nil).Once()
		repoGoogle.On("GetUserInfo", mock.Anything).Return(mock_google_core, nil).Once()
		repoAuth.On("FindAccount", mock.Anything).Return(mock_mentor_musiclab, nil).Once()
		repoMentor.On("UpdateData", mock.Anything, mock.Anything).Return(nil).Once()

		srv := New(repoAuth, repoGoogle, repoTransaction, repoClass, repoStudent, repoSchedule, repoMentor)
		err := srv.RedirectGoogleCallback("code")
		assert.Nil(t, err)
		repoAuth.AssertExpectations(t)
	})

	t.Run("Redirect Oauth Student", func(t *testing.T) {
		tokenOauth := &oauth2.Token{
			AccessToken: mock_student_musiclab.TokenOauth,
		}

		repoGoogle.On("GetToken", mock.Anything).Return(tokenOauth, nil).Once()
		repoGoogle.On("GetUserInfo", mock.Anything).Return(mock_google_core, nil).Once()
		repoAuth.On("FindAccount", mock.Anything).Return(mock_student_musiclab, nil).Once()
		repoStudent.On("UpdateData", mock.Anything, mock.Anything).Return(nil).Once()

		srv := New(repoAuth, repoGoogle, repoTransaction, repoClass, repoStudent, repoSchedule, repoMentor)
		err := srv.RedirectGoogleCallback("code")
		assert.Nil(t, err)
		repoAuth.AssertExpectations(t)
	})

	t.Run("Failed Get User Info", func(t *testing.T) {
		tokenOauth := &oauth2.Token{
			AccessToken: mock_student_musiclab.TokenOauth,
		}

		repoGoogle.On("GetToken", mock.Anything).Return(tokenOauth, nil).Once()
		repoGoogle.On("GetUserInfo", mock.Anything).Return(helper.GoogleCore{}, errors.New("error")).Once()

		srv := New(repoAuth, repoGoogle, repoTransaction, repoClass, repoStudent, repoSchedule, repoMentor)
		err := srv.RedirectGoogleCallback("code")
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), err.Error())
		repoAuth.AssertExpectations(t)
	})

	t.Run("Failed Find Account", func(t *testing.T) {
		tokenOauth := &oauth2.Token{
			AccessToken: mock_student_musiclab.TokenOauth,
		}

		repoGoogle.On("GetToken", mock.Anything).Return(tokenOauth, nil).Once()
		repoGoogle.On("GetUserInfo", mock.Anything).Return(mock_google_core, nil).Once()
		repoAuth.On("FindAccount", mock.Anything).Return(auth.Core{}, errors.New("error")).Once()

		srv := New(repoAuth, repoGoogle, repoTransaction, repoClass, repoStudent, repoSchedule, repoMentor)
		err := srv.RedirectGoogleCallback("code")
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), err.Error())
		repoAuth.AssertExpectations(t)
	})

	t.Run("Failed Update Data Mentor", func(t *testing.T) {
		tokenOauth := &oauth2.Token{
			AccessToken: mock_mentor_musiclab.TokenOauth,
		}

		repoGoogle.On("GetToken", mock.Anything).Return(tokenOauth, nil).Once()
		repoGoogle.On("GetUserInfo", mock.Anything).Return(mock_google_core, nil).Once()
		repoAuth.On("FindAccount", mock.Anything).Return(mock_mentor_musiclab, nil).Once()
		repoMentor.On("UpdateData", mock.Anything, mock.Anything).Return(errors.New("error")).Once()

		srv := New(repoAuth, repoGoogle, repoTransaction, repoClass, repoStudent, repoSchedule, repoMentor)
		err := srv.RedirectGoogleCallback("code")
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), err.Error())
		repoAuth.AssertExpectations(t)
	})

	t.Run("Failed Update Data Student", func(t *testing.T) {
		tokenOauth := &oauth2.Token{
			AccessToken: mock_student_musiclab.TokenOauth,
		}

		repoGoogle.On("GetToken", mock.Anything).Return(tokenOauth, nil).Once()
		repoGoogle.On("GetUserInfo", mock.Anything).Return(mock_google_core, nil).Once()
		repoAuth.On("FindAccount", mock.Anything).Return(mock_student_musiclab, nil).Once()
		repoStudent.On("UpdateData", mock.Anything, mock.Anything).Return(errors.New("error")).Once()

		srv := New(repoAuth, repoGoogle, repoTransaction, repoClass, repoStudent, repoSchedule, repoMentor)
		err := srv.RedirectGoogleCallback("code")
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), err.Error())
		repoAuth.AssertExpectations(t)
	})
}
