package utils

import (
	"errors"
	"github.com/alkosmas92/xm-golang/internal/models"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
	"time"
)

var jwtKey = []byte(viper.GetString("jwt.secret_key"))
var oldJwtKey = []byte(viper.GetString("jwt.old_secret_key"))

func GenerateJWT(userID, username string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &models.Claims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
func ValidateJWT(tokenStr string) (*models.Claims, error) {
	claims := &models.Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		// Check if the error is due to the token being expired
		if ve, ok := err.(*jwt.ValidationError); ok && ve.Errors == jwt.ValidationErrorExpired {
			// Return claims even if the token is expired
			return claims, errors.New("token is expired")
		}
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
