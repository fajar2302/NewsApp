package service_test

import (
	users "NEWSAPP/features/Users"
	"NEWSAPP/features/Users/service"
	"NEWSAPP/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegister(t *testing.T) {
	mockUserData := new(mocks.DataUserInterface)
	mockHashService := new(mocks.HashInterface)
	userService := service.New(mockUserData, mockHashService)

	t.Run("success", func(t *testing.T) {
		user := users.User{
			FullName:    "John Doe",
			Email:       "johndoe@example.com",
			Password:    "password123",
			PhoneNumber: "08123456789",
			Address:     "Some Address",
		}

		mockHashService.On("HashPassword", mock.Anything).Return("hashedpassword", nil).Once()
		mockUserData.On("CreateAccount", mock.Anything).Return(nil).Once()

		err := userService.RegistrasiAccount(user)

		assert.NoError(t, err)
		mockHashService.AssertExpectations(t)
		mockUserData.AssertExpectations(t)
	})

	t.Run("validation error", func(t *testing.T) {
		user := users.User{}

		err := userService.RegistrasiAccount(user)

		assert.Error(t, err)
		assert.Equal(t, "[validation] nama/email/password/phone/address tidak boleh kosong", err.Error())
	})

	t.Run("hash error", func(t *testing.T) {
		user := users.User{
			FullName:    "John Doe",
			Email:       "johndoe@example.com",
			Password:    "password123",
			PhoneNumber: "08123456789",
			Address:     "Some Address",
		}

		mockHashService.On("HashPassword", mock.Anything).Return("", errors.New("hash error")).Once()

		err := userService.RegistrasiAccount(user)

		assert.Error(t, err)
		assert.Equal(t, "hash error", err.Error())
		mockHashService.AssertExpectations(t)
	})
}

func TestLogin(t *testing.T) {
	mockUserData := new(mocks.DataUserInterface)
	mockHashService := new(mocks.HashInterface)
	userService := service.New(mockUserData, mockHashService)

	t.Run("success", func(t *testing.T) {
		email := "johndoe@example.com"
		password := "password123"
		hashedPassword := "hashedpassword"
		user := &users.User{
			FullName:    "John Doe",
			Email:       email,
			Password:    hashedPassword,
			PhoneNumber: "08123456789",
			Address:     "Some Address",
		}
		token := "sometoken"

		mockUserData.On("AccountByEmail", email).Return(user, nil).Once()
		mockHashService.On("CheckPasswordHash", hashedPassword, password).Return(true).Once()

		// Assuming that token generation is done within LoginAccount
		mockUserData.On("GenerateToken", user.UserID).Return(token, nil).Once()

		returnedUser, returnedToken, err := userService.LoginAccount(email, password)

		assert.NoError(t, err)
		assert.Equal(t, user, returnedUser)
		assert.Equal(t, token, returnedToken)
		mockUserData.AssertExpectations(t)
		mockHashService.AssertExpectations(t)
	})

	t.Run("account not found", func(t *testing.T) {
		email := "johndoe@example.com"
		password := "password123"

		mockUserData.On("AccountByEmail", email).Return(nil, errors.New("account not found")).Once()

		returnedUser, returnedToken, err := userService.LoginAccount(email, password)

		assert.Error(t, err)
		assert.Equal(t, "account not found", err.Error())
		assert.Nil(t, returnedUser)
		assert.Empty(t, returnedToken)
		mockUserData.AssertExpectations(t)
	})

	t.Run("password verification error", func(t *testing.T) {
		email := "johndoe@example.com"
		password := "password123"
		hashedPassword := "hashedpassword"
		user := &users.User{
			FullName:    "John Doe",
			Email:       email,
			Password:    hashedPassword,
			PhoneNumber: "08123456789",
			Address:     "Some Address",
		}

		mockUserData.On("AccountByEmail", email).Return(user, nil).Once()
		mockHashService.On("CheckPasswordHash", hashedPassword, password).Return(false).Once()

		returnedUser, returnedToken, err := userService.LoginAccount(email, password)

		assert.Error(t, err)
		assert.Equal(t, "password verification error", err.Error())
		assert.Nil(t, returnedUser)
		assert.Empty(t, returnedToken)
		mockUserData.AssertExpectations(t)
		mockHashService.AssertExpectations(t)
	})
}

func TestUpdateProfile(t *testing.T) {
	mockUserData := new(mocks.DataUserInterface)
	mockHashService := new(mocks.HashInterface)
	userService := service.New(mockUserData, mockHashService)

	t.Run("success", func(t *testing.T) {
		userID := uint(1)
		user := users.User{
			ProfilePicture: "piture.png",
			FullName:       "John Doe",
			Email:          "johndoe@example.com",
			Password:       "password",
			PhoneNumber:    "08123456789",
			Address:        "Some Address",
		}

		mockHashService.On("HashPassword", mock.Anything).Return("password", nil).Once()
		mockUserData.On("UpdateAccount", userID, user).Return(nil).Once()

		err := userService.UpdateProfile(userID, user)

		assert.NoError(t, err)
		mockUserData.AssertExpectations(t)
		mockHashService.AssertExpectations(t)
	})

	t.Run("update error", func(t *testing.T) {
		userID := uint(1)
		user := users.User{
			FullName:    "John Doe",
			Email:       "johndoe@example.com",
			PhoneNumber: "08123456789",
			Address:     "Some Address",
		}

		mockUserData.On("UpdateAccount", userID, user).Return(errors.New("update error")).Once()

		err := userService.UpdateProfile(userID, user)

		assert.Error(t, err)
		assert.Equal(t, "update error", err.Error())
		mockUserData.AssertExpectations(t)
	})
}

func TestDeleteProfile(t *testing.T) {
	mockUserData := new(mocks.DataUserInterface)
	mockHashService := new(mocks.HashInterface)
	userService := service.New(mockUserData, mockHashService)

	t.Run("success", func(t *testing.T) {
		userID := uint(1)

		mockUserData.On("DeleteAccount", userID).Return(nil).Once()

		err := userService.DeleteAccount(userID)

		assert.NoError(t, err)
		mockUserData.AssertExpectations(t)
	})

	t.Run("delete error", func(t *testing.T) {
		userID := uint(1)

		mockUserData.On("DeleteAccount", userID).Return(errors.New("delete error")).Once()

		err := userService.DeleteAccount(userID)

		assert.Error(t, err)
		assert.Equal(t, "delete error", err.Error())
		mockUserData.AssertExpectations(t)
	})
}

func TestGetProfile(t *testing.T) {
	mockUserData := new(mocks.DataUserInterface)
	mockHashService := new(mocks.HashInterface)
	userService := service.New(mockUserData, mockHashService)

	t.Run("success", func(t *testing.T) {
		userID := uint(1)
		user := &users.User{
			FullName:    "John Doe",
			Email:       "johndoe@example.com",
			PhoneNumber: "08123456789",
			Address:     "Some Address",
		}

		mockUserData.On("AccountById", userID).Return(user, nil).Once()

		returnedUser, err := userService.GetProfile(userID)

		assert.NoError(t, err)
		assert.Equal(t, user, returnedUser)
		mockUserData.AssertExpectations(t)
	})

	t.Run("profile not found", func(t *testing.T) {
		userID := uint(1)

		mockUserData.On("AccountById", userID).Return(nil, errors.New("profile not found")).Once()

		returnedUser, err := userService.GetProfile(userID)

		assert.Error(t, err)
		assert.Equal(t, "profile not found", err.Error())
		assert.Nil(t, returnedUser)
		mockUserData.AssertExpectations(t)
	})
}
