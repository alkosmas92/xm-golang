package services_test

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"testing"

	"github.com/alkosmas92/xm-golang/internal/mocks"
	"github.com/alkosmas92/xm-golang/internal/models"
	"github.com/alkosmas92/xm-golang/internal/services"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestRegisterUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepository(ctrl)
	userService := services.NewUserService(mockRepo)
	ctx := context.Background()

	user := models.NewUser("testuser", "password", "Test", "User")

	mockRepo.EXPECT().CreateUser(ctx, user).Return(nil)

	err := userService.RegisterUser(ctx, user)
	assert.NoError(t, err)
}
func TestAuthenticateUser_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepository(ctrl)
	userService := services.NewUserService(mockRepo)
	ctx := context.Background()

	// Hash the password using bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	assert.NoError(t, err)

	// Create the user with the hashed password
	user := &models.User{
		UserID:    "testuser-id",
		Username:  "testuser",
		Password:  string(hashedPassword),
		FirstName: "Test",
		LastName:  "User",
	}

	mockRepo.EXPECT().GetUserByUsername(ctx, "testuser").Return(user, nil)

	// Authenticate the user
	result, err := userService.AuthenticateUser(ctx, "testuser", "password")
	assert.NoError(t, err)
	assert.Equal(t, user, result)
}
