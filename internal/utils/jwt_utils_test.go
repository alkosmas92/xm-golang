package utils_test

import (
	"github.com/alkosmas92/xm-golang/internal/models"
	"github.com/alkosmas92/xm-golang/internal/utils"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var jwtKey = []byte(viper.GetString("jwt.secret_key"))

func TestGenerateJWT(t *testing.T) {
	userID := "testUserID"
	username := "testUsername"
	token, err := utils.GenerateJWT(userID, username)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if token == "" {
		t.Fatalf("Expected a token, got an empty string")
	}
}

func TestValidateJWT(t *testing.T) {
	userID := "testUserID"
	username := "testUsername"
	token, err := utils.GenerateJWT(userID, username)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	claims, err := utils.ValidateJWT(token)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if claims.UserID != userID {
		t.Fatalf("Expected userID %v, got %v", userID, claims.UserID)
	}

	if claims.Username != username {
		t.Fatalf("Expected username %v, got %v", username, claims.Username)
	}
}

func TestValidateJWT_ExpiredToken(t *testing.T) {
	userID := "testUserID"
	username := "testUsername"
	claims := &models.Claims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(-time.Hour)),
		},
	}
	var cl *models.Claims

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	cl, err = utils.ValidateJWT(tokenString)
	if err == nil {
		t.Fatalf("Expected error for expired token, got nil")
	}
	assert.Equal(t, cl.UserID, claims.UserID)
}
