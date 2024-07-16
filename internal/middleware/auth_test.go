package middleware_test

import (
	appContext "github.com/alkosmas92/xm-golang/internal/context"
	"github.com/alkosmas92/xm-golang/internal/models"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/alkosmas92/xm-golang/internal/middleware"
	"github.com/stretchr/testify/assert"
)

var jwtKey = []byte(viper.GetString("jwt.secret_key"))

func GenerateTestJWT(userID, username string, expired bool) (string, error) {
	expirationTime := time.Now().Add(72 * time.Hour)
	if expired {
		expirationTime = time.Now().Add(-time.Hour)
	}
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

func TestAuthMiddleware(t *testing.T) {
	tests := []struct {
		name           string
		token          string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Missing Authorization Header",
			token:          "",
			expectedStatus: http.StatusUnauthorized,
			expectedBody:   "missing authorization header\n",
		},
		{
			name:           "Invalid Authorization Format",
			token:          "InvalidFormat",
			expectedStatus: http.StatusUnauthorized,
			expectedBody:   "token contains an invalid number of segments\n",
		},
		{
			name: "Expired Token",
			token: func() string {
				token, _ := GenerateTestJWT("6b5d04aa-8d5b-4a24-998d-a02dd324830f", "testuser", true)
				return token
			}(),
			expectedStatus: http.StatusUnauthorized,
			expectedBody:   "token is expired\n",
		},
		{
			name: "Valid Token",
			token: func() string {
				token, _ := GenerateTestJWT("6b5d04aa-8d5b-4a24-998d-a02dd324830f", "testuser", false)
				return token
			}(),
			expectedStatus: http.StatusOK,
			expectedBody:   "Success",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a test handler to use with the middleware
			testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				userID := r.Context().Value(appContext.UserIDKey).(string)
				username := r.Context().Value(appContext.UsernameKey).(string)
				assert.Equal(t, "6b5d04aa-8d5b-4a24-998d-a02dd324830f", userID)
				assert.Equal(t, "testuser", username)
				w.Write([]byte("Success"))
			})
			handlerToTest := middleware.AuthMiddleware(testHandler)

			ts := httptest.NewServer(handlerToTest)
			defer ts.Close()

			req, err := http.NewRequest("GET", ts.URL, nil)
			assert.NoError(t, err)

			if tt.token != "" {
				req.Header.Set("Authorization", "Bearer "+tt.token)
			}

			resp, err := http.DefaultClient.Do(req)
			assert.NoError(t, err)
			defer resp.Body.Close()

			body := new(strings.Builder)
			_, err = io.Copy(body, resp.Body)
			assert.NoError(t, err)

			assert.Equal(t, tt.expectedStatus, resp.StatusCode)
			assert.Equal(t, tt.expectedBody, body.String())
		})
	}
}
